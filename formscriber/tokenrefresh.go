package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func refreshtoken() {

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

	fmt.Println(token)

}
