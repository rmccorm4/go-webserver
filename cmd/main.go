package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
  // Webpage Output
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<h1>Sample Webpage</h1>")
  // Command-Line Output
  fmt.Printf("%s accessed.\n", r.URL.Path)
}

func contact(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "Contact information placeholder.")
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", home)
  router.HandleFunc("/contact", contact)

  url := "localhost:3000"
  fmt.Printf("Web server running on %s\n", url)

  // nil instead of router here means to use the default ServeMux
  http.ListenAndServe(url, router)
}
