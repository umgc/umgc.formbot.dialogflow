package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func updateDocument(docid, intent, value string) {
	var t Tokenresponse
	var url = string(Oauth())

	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
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

	err = json.Unmarshal(body, &t)

	if err != nil {
		panic(err)
	}

	token := t.Access_token

	url = "https://docs.googleapis.com/v1/documents/" + docid + ":batchUpdate"
	method = "POST"

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

	client = &http.Client{}
	req, err = http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
