package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/* 2.2 GET Team
 * URL: http://localhost:8080/getTeam
 */
type Team struct {
	Type string       `json:"type"`
	D    []TeamMember `json:"d"`
}
type TeamMember struct {
	FName string `json:"fname"`
	LName string `json:"lname"`
	Role  string `json:"role"`
}

func GetTeamEndPoint(w http.ResponseWriter, request *http.Request) {
	//	REST endpoint to get articles
	jsonFile, err := os.Open("DATA/GetTeam.JSON")

	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened JSON as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var team Team
	json.Unmarshal(byteValue, &team)

	json.NewEncoder(w).Encode(team)
}
