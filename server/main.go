package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mitchan/gymbro/db"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize DB connection: %s", err)
	}
	defer dbConn.Close()

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
