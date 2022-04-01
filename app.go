package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Message struct {
	Detail string
}

func GetHello(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Message{"Hello Go! 😎"})
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", GetHello).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", router))
}