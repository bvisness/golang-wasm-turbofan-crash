package main

import (
    "log"
    "net/http"
)

func main() {
    log.Println("listening on :8989...")
    err := http.ListenAndServe(":8989", http.FileServer(http.Dir(".")))
    log.Fatalln(err)
}
