package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "os"
)

func main() {

    // Create a new logger that writes to stdout
    logger := log.New(os.Stdout, "", log.LstdFlags)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    logger.Println("Starting server on :8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        logger.Fatal("Failed to start server: ", err)
    }
}
