package main

import (
	"fmt"
	"net/http"
	"os"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
  host, _ := os.Hostname()
  fmt.Fprintf(w, "Hello, %s", host)
}

func main() {
  http.HandleFunc("/", HelloServer)
  http.ListenAndServe(":8080",nil)
}
