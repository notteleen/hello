package main

import (
	"fmt"
	"net/http"

)

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
          fmt.Fprintln(w, "Hello, World!")
        })
    
        http.ListenAndServe("0.0.0.0:8080", nil)
}
