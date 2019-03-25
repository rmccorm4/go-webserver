package main

import (
  "fmt"
  "net/http"

  "github.com/julienschmidt/httprouter"
)

// Handler for whenever our webpage receives an HTTP request
func rootHandler(w http.ResponseWriter, r *http.Request) {
  // Webpage Output
  w.Header().Set("Content-Type", "text/html")
  if r.URL.Path == "/" {
    fmt.Fprint(w, "<h1>Sample Webpage</h1>")
  } else if r.URL.Path == "/contact" {
    fmt.Fprint(w, "Contact information placeholder.")
  } else {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprint(w, "<h1>404 Not Found.</h1>")
  }
  // Command-Line Output
  fmt.Printf("%s accessed.\n", r.URL.Path)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
  router := httprouter.New()
  //router.GET("/", rootHandler)
  router.GET("/hello/:name", Hello)
  //http.HandleFunc("/", rootHandler)
  url := "localhost:3000"
  fmt.Printf("Web server running on %s\n", url)
  // nil instead of router here means to use the default ServeMux
  http.ListenAndServe(url, router)
}
