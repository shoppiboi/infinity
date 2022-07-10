package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// type Quote struct {
// 	Body   string
// 	Author string
// }

func getQuote() string {
	const quoteApiUrl = "https://stoic-server.herokuapp.com/random"

	response, err := http.Get(quoteApiUrl)
	if err != nil {
		log.Fatalln(err)
	}

	//	close the request
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	return string(content)
}

func main() {

	quote := getQuote()

	fmt.Printf(string(quote))

	// body, err := ioutil.ReadAll(response.Body)

	// fmt.Println(string(body))

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// stringifiedBody := string(body)

	// fmt.Print(stringifiedBody)

}
