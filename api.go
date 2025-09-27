package main

import (
    "encoding/json"
    "net/http"
)

func StartServer() {
    http.HandleFunc("/blocks", handleGetBlocks)
    http.HandleFunc("/mine", handleMineBlock)
    http.ListenAndServe(":8080", nil)
}

func handleGetBlocks(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Blockchain)
}

func handleMineBlock(w http.ResponseWriter, r *http.Request) {
    data := r.URL.Query().Get("data")
    if data == "" {
        http.Error(w, "Missing data", http.StatusBadRequest)
        return
    }
    newBlock := AddBlock(data)
    json.NewEncoder(w).Encode(newBlock)
}
