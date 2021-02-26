package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func updateDocument(docid, intent, value string) {

	url := "https://docs.googleapis.com/v1/documents/" + docid + ":batchUpdate"
	method := "POST"

	payload := strings.NewReader(`{
	"requests": [
	  {
		"replaceAllText": {
		  "containsText": {
			"text": "{{` + intent + `}}",
			"matchCase": "false"
		  },
		  "replaceText": "` + value + `",
		}
	  }
	]
  }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+getToken())
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
