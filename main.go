package main

import (
    "net/http"
    "fmt"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    var url = os.Getenv("NEW_HOST") + r.URL.String()

    w.Header().Add("Location", url)
    w.WriteHeader(http.StatusFound)

    fmt.Println("Redirecting " + r.URL.String() + " ...")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
