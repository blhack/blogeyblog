package routes

import("fmt"
       "net/http"
       "io/ioutil"
       "github.com/blhack/blogeyBlog/config"
       "os"
       "bufio"
       "bytes"
       "text/template"
       "path/filepath"
       "strings")

type ListOfLinks struct {
  Links string
}

func renderLinkTemplate(w http.ResponseWriter, htmlLinks string) {
  links := ListOfLinks{htmlLinks}
  tmpl, err := template.ParseFiles("./templates/index.template")
  if err != nil {
    fmt.Println(err)
  }
  var rendered bytes.Buffer
  tmpl.Execute(&rendered, links)
  fmt.Fprintf(w, rendered.String())
}

func Index(w http.ResponseWriter, r *http.Request) {

  if r.URL.Path != "/" {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "404 not found")
    return
  }

  files, _ := ioutil.ReadDir(config.BlogPostDirectory)

  linkList := ""

  for _,file := range files {

    if filepath.Ext(file.Name()) == ".blogeyBlog" {

      fileName := fmt.Sprintf("%v/%v", config.BlogPostDirectory,file.Name())
      f,_ := os.Open(fileName)
      defer f.Close()

      reader := bufio.NewReader(f)
      blogTitle,_ := reader.ReadString('\n')

      linkList = linkList + fmt.Sprintf("<a href='/show?p=%v'>%v</a>", strings.Split(file.Name(),".")[0],blogTitle)
    }
  }
  renderLinkTemplate(w, linkList)
}