package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		type responsePayload struct {
			Success bool `json:"success"`
		}

		payload := responsePayload{Success: true}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
