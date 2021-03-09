package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/* 2.3 GET HAQs
 * URL: http://localhost:8080/getFAQs
 */
type FAQs struct {
	Type string `json:"type"`
	D    []FAQ  `json:"d"`
}
type FAQ struct {
	Q string `json:"q"`
	A string `json:"a"`
}

func GetFAQsEndPoint(w http.ResponseWriter, request *http.Request) {
	//	REST endpoint to get articles
	jsonFile, err := os.Open("DATA/GetFAQs.JSON")

	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened JSON as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var faqs FAQs
	json.Unmarshal(byteValue, &faqs)

	json.NewEncoder(w).Encode(faqs)
}
