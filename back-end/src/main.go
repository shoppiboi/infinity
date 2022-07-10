package main

import (
	"fmt"
	"log"
	"net/http"
	q "src/quote"
)

func main() {

	http.HandleFunc("/get_quote", func(w http.ResponseWriter, r *http.Request) {
		quote := q.GetQuote()

		fmt.Fprintf(w, quote.Body)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
