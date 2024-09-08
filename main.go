package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from Go backend!"))
}

func main() {
    fs := http.FileServer(http.Dir("frontend"))
    http.Handle("/", fs)

    http.HandleFunc("/api/hello", helloHandler)
    log.Println("Listening on :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}