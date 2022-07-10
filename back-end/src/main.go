package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Quote struct {
	Body   string `json:"body"`
	Author string `json:"author"`
}

//	stoic-server returns an extra [] in the body. this function removes them
func cleanUpRequestBody(requestBody []byte) []byte {

	cleanBody := requestBody[1 : len(requestBody)-1]

	return cleanBody
}

func getQuoteBytes() []byte {
	const quoteApiUrl = "https://stoic-server.herokuapp.com/random"

	response, err := http.Get(quoteApiUrl)
	if err != nil {
		log.Fatalln(err)
	}

	//	close the request
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	cleanBody := cleanUpRequestBody(body)

	return cleanBody
}

func main() {
	var quoteBytes []byte = getQuoteBytes()

	var responseString strings.Builder
	responseString.Write(quoteBytes)

	var quote Quote

	err := json.Unmarshal([]byte(quoteBytes), &quote)
	if err != nil {
		panic(err)
	}

	fmt.Println(quote)
}
