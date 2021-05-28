package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tidwall/gjson"
)

func getAccessToken() string {
	getToken := getToken()
	fmt.Println("getToken: " + getToken)
	url := "https://oauth2.googleapis.com/token?grant_type=urn:ietf:params:oauth:grant-type:jwt-bearer&assertion=" + getToken
	fmt.Println("ACCESS_TOKEN!!!: " + url)
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
	fmt.Println("")
	fmt.Println("so far so good" + string(body))
	fmt.Println("")
	accesstoken := gjson.Get(string(body), "access_token")
	fmt.Println("ROUND 2!!!!!!" + accesstoken.String())
	return accesstoken.String()
}

func getToken() string {
	tokenString, err := createSignedTokenString()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Signed token string:\n%v\n", tokenString)
	return tokenString
}

func createSignedTokenString() (string, error) {
	privateKey, err := ioutil.ReadFile("private_key.rsa")
	if err != nil {
		return "", fmt.Errorf("error reading private key file: %v\n", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", fmt.Errorf("error parsing RSA private key: %v\n", err)
	}

	type MyCustomClaims struct {
		Scope string `json:"scope"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"https://www.googleapis.com/auth/drive https://www.googleapis.com/auth/dialogflow https://www.googleapis.com/auth/drive.file",
		jwt.StandardClaims{
			Issuer:    "formbot@form-bot-1577a.iam.gserviceaccount.com",
			Audience:  "https://oauth2.googleapis.com/token",
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error signing token: %v\n", err)
	}

	return tokenString, nil
}
