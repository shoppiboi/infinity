package main

import (
	"encoding/json"
	"log"
	"net/http"

	q "src/quote"
)

func main() {

	http.HandleFunc("/get_quote", func(w http.ResponseWriter, r *http.Request) {
		quote := q.GetQuote()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"body": quote.Body, "author": quote.Author})	
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
