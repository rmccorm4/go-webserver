package main

import (
  "fmt"
  "net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<h1>Sample Webpage</h1>")
}

func main() {
  http.HandleFunc("/", handlerFunc)
  // nil here means to use the default ServerMux
  http.ListenAndServe("localhost:3000", nil)
}