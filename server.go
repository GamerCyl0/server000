package main

import (
    "log"
    "net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {

    http.Redirect(w, r, "http://github.com/GamerCyl0/server000/blob/main/server.go", 302)
}

func main() {
    http.HandleFunc("/", redirect)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
