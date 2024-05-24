package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandlePing(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling request\t\t ping")
	response := map[string]interface{}{
		"resp_code": 0,
		"msg":       "success",
		"data":      "pong",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/ping", HandlePing)

	log.Println("Starting node registry on port 8007...")
	http.ListenAndServe(":8007", nil)
}
