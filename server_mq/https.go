package main

import (
	"fmt"
	"log"
	"net/http"
)

func HttpsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        fmt.Fprintf(w, "1")
    }
}

func RunHttps(port string) {
    http.HandleFunc("/", HttpsHandler)
    log.Fatal(http.ListenAndServeTLS(port, "ssl/cert.pem", "ssl/key.pem", nil))
}