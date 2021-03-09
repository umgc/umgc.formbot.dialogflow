package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/* 2.1 GET Hep Articles
 * URL: http://localhost:8080/getArticles
 */
type Articles struct {
	Type string    `json:"type"`
	D    []Article `json:"d"`
}
type Article struct {
	ID       string   `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Desc     string   `json:"description,omitempty"`
	URL      string   `json:"URL"`
	KeyWords *KeyWord `json:"keywords,omitempty"`
}
type KeyWord struct {
	Name string `json:"name,omitempty"`
}

func GetArticlesEndPoint(w http.ResponseWriter, request *http.Request) {
	//	REST endpoint to get articles
	jsonFile, err := os.Open("DATA/GetHelpArticleList.JSON")

	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened JSON as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var articles Articles
	json.Unmarshal(byteValue, &articles)

	json.NewEncoder(w).Encode(articles)
}
