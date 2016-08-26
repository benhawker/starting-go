// A minimal "echo "server

package main

import "fmt"
import "log"
import "net/http"

func main() {
  http.HandleFunc("/", handler) // Each request calls handler
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Handler echoes the path component of the requested URL

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Fire up the server with go run server.go &
// Access in your browser at localhost 8000.