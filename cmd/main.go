package main

import (
  "fmt"
  "net/http"
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

func main() {
  http.HandleFunc("/", rootHandler)
  // nil here means to use the default ServerMux
  url := "localhost:3000"
  fmt.Printf("Web server running on %s\n", url)
  http.ListenAndServe(url, nil)
}
