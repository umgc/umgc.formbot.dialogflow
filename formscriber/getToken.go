package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getToken() string {

	var t Tokenresponse
	var url = string(Oauth())
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)

	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)

	}
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	token := t.Access_token

	return token
}
