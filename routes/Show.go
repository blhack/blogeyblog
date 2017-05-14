package routes

import("fmt"
       "net/http"
       "github.com/blhack/blogeyBlog/config"
       "os"
       "path/filepath"
       "bufio"
       "github.com/russross/blackfriday"
       "text/template"
       "bytes"
)

type BlogPost struct {
  Title string
  Body string 
}

func getBlog(path string) (string, string) {
  f,_ := os.Open(path)
  defer f.Close()

  scanner := bufio.NewScanner(f)

  body := ""
  title := ""
  index := 0
  for scanner.Scan() {
    if index == 0 {
            title = scanner.Text()
    } else {
            body = body + scanner.Text() + "\n"
    }
    index++
  }
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
  }

  return title, body

}

func renderPostTemplate(w http.ResponseWriter, title string, body string) {
  post := BlogPost{title,body}
  tmpl, err := template.ParseFiles("./templates/post.template")
  if err != nil {
    fmt.Println(err)
  }
  var rendered bytes.Buffer
  tmpl.Execute(&rendered, post)
  fmt.Fprintf(w, rendered.String())
}

func Show(w http.ResponseWriter, r *http.Request) {
  postName := r.FormValue("p")
  _, file := filepath.Split(postName)
  path := filepath.Join(config.BlogPostDirectory, file)
  path = fmt.Sprintf("%v.blogeyBlog", path)
  _, err := os.Stat(path)
 
  fmt.Fprintf(w, "<html>")

  if err == nil {
    title,body := getBlog(path)
    body = string(blackfriday.MarkdownCommon([]byte(body)))

    renderPostTemplate(w,title,body)
  } else {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "404 not found")
  }
  fmt.Fprintf(w, "</html>")
}