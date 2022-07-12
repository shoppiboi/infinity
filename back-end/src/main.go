package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	q "src/quote"
)

func main() {

	PORT := os.Getenv("PORT")

	fmt.Println("Listening for requests on port " + PORT)

	http.HandleFunc("/get_quote", func(w http.ResponseWriter, r *http.Request) {
		quote := q.GetQuote()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"body": quote.Body, "author": quote.Author})	
	})

	if err := http.ListenAndServe(":"+ PORT, nil); err != nil {
		log.Fatal(err)
	}
}
