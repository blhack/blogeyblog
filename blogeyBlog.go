package main

import("net/http"
       "github.com/blhack/blogeyBlog/routes")

func main() {
   http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
   http.HandleFunc("/show", routes.Show)
   http.HandleFunc("/", routes.Index)
   http.ListenAndServe(":80", nil)
}
