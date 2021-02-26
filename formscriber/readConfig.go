package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
)

type Configuration struct {
	DB struct {
		ID       string `json:"id"`
		User     string `json:"User"`
		Password string `json:"Password"`
		Database string `json:"Database"`
	} `json:"DB"`
	Webhook struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Method   string `json:"method"`
	} `json:"Webhook"`
	Oauth struct {
		URL string `json:"url"`
	} `json:"Oauth"`
}

func apiUsername() string {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	body, err := ioutil.ReadAll(jsonFile)
	username := gjson.Get(string(body), "Webhook.Username")
	return username.String()
}

func apiPassword() string {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	body, err := ioutil.ReadAll(jsonFile)
	password := gjson.Get(string(body), "Webhook.Password")
	return password.String()
}

func Oauth() string {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	body, err := ioutil.ReadAll(jsonFile)
	oauth := gjson.Get(string(body), "Oauth.Url")
	return oauth.String()
}
