// Server2 is a minimal echo and counter server
package main

import (
  "fmt"
  "log"
  "net/http"
  "sync"
)


//Init a var called mu with type "sync.Mutex"
var mu sync.Mutex

// Init a var count with type int
var count int


// The server has 2 handlers - the request URL determines which one is called.
func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/count", counter)
  log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

// To avoid race conditions when 2 concurrent requests
// try to update count at the same time with use a mutex lock/unlock.

// Handler echoes the path component of the request URL
func handler(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  count++
  mu.Unlock()
  fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Counter echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  fmt.Fprintf(w, "Count %d\n", count)
  mu.Unlock()
}