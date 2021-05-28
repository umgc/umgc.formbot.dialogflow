package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
)

const cfg = "config.json"

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
	jsonFile, err := os.Open(cfg)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	body, err := ioutil.ReadAll(jsonFile)
	username := gjson.Get(string(body), "Webhook.Username")
	return username.String()
}

func apiPassword() string {
	jsonFile, err := os.Open(cfg)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	body, err := ioutil.ReadAll(jsonFile)
	password := gjson.Get(string(body), "Webhook.Password")
	return password.String()
}

func Oauth() string {
	jsonFile, err := os.Open(cfg)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	body, err := ioutil.ReadAll(jsonFile)
	oauth := gjson.Get(string(body), "Oauth.Url")
	return oauth.String()
}

func ServiceWorker() (string, string, string, string, string, string, string) {
	jsonFile, err := os.Open(cfg)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	body, err := ioutil.ReadAll(jsonFile)
	email := gjson.Get(string(body), "ServiceWorker.email")
	client_id := gjson.Get(string(body), "ServiceWorker.client_id")
	auth_uri := gjson.Get(string(body), "ServiceWorker.auth_uri")
	token_uri := gjson.Get(string(body), "ServiceWorker.token_uri")
	auth_provider_x509_cert_url := gjson.Get(string(body), "ServiceWorker.auth_provider_x509_cert_url")
	client_secret := gjson.Get(string(body), "ServiceWorker.client_secret")
	public_key := gjson.Get(string(body), "ServiceWorker.public_key")
	return email.String(), client_id.String(), auth_uri.String(), token_uri.String(), auth_provider_x509_cert_url.String(), client_secret.String(), public_key.String()
}
