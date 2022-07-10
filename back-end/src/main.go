package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Quote struct {
	Body   string
	Author string
}

// func handleRequest() {

// }

// func convertBodyFromBytesToJson() {

// 	return

// }

func main() {
	resp, err := http.Get("https://stoic-server.herokuapp.com/random")

	// fmt.Println(resp)
	if err != nil {
		log.Fatalln(err)
	}

	// data, err := ioutil.ReadAll(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// stringifiedBody := string(body)

	// fmt.Print(stringifiedBody)

}
