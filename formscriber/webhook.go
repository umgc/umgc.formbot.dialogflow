package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/tidwall/gjson"
)

func basicAuth(h http.HandlerFunc) http.HandlerFunc {
	var docid string

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {

			return
		}

		//var username = apiUsername()
		//var password = apiPassword()
		if pair[0] != string(apiUsername()) || pair[1] != string(apiPassword()) {

			http.Error(w, "Not authorized", 401)
			return
		}

		switch r.Method {
		case "GET":

			fmt.Fprintf(w, "Sorry, only POST methods are supported.")

		case "POST":

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			if err != nil {
				panic(err)
			}

			//
			fmt.Println(string(body))
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("parsing body!")
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println(" ")

			// parse out the intent/action
			action := gjson.Get(string(body), "queryResult.action")

			//Tip: Struct is only made for easy JSON node parsing using JsonPath
			//Tip: var request WebhookRequest
			//
			// 1) If ACTION = New report with URL then parse URL and get DocID to get all fields from template then build Dialogflow intents with batchUpdate
			//var request WebhookRequest
			//request.Session
			if action.String() == "NewReportURL" {
				// parse out the url.
				url := gjson.Get(string(body), "queryResult.parameters.template")
				fmt.Println(" ")
				fmt.Println(" ")
				fmt.Println("URL: " + url.String())
				fmt.Println(" ")
				fmt.Println(" ")
				//
				//
				//
				//
				//get_document routine to get list of fields from report:

				//extract id from url
				// regex: /document/d/([a-zA-Z0-9-_]+)
				regex, _ := regexp.Compile("/document/d/([a-zA-Z0-9-_]+)")
				//fmt.Println(r.FindString(url.String()))
				docid = regex.FindString(url.String())
				// Extract Doc ID:
				fmt.Println("docid: " + trimLeftChars(docid, 12))
				fmt.Println(getDocument(trimLeftChars(docid, 12)))

				//
				//
				session := gjson.Get(string(body), "session")
				fmt.Println(" ")
				fmt.Println(" ")
				fmt.Println("Session: " + session.String())
				fmt.Println(" ")
				fmt.Println(" ")
				//
				http.ServeFile(w, r, "url_action_response.json")
				// 2) If ACTION = Field intent, then take URL and make a POST request to the doc and update the field.
				// parse out the intent/field name
			} else if action.String() == "FillDocField" {
				fmt.Println(" ")
				fmt.Println(" ")
				fmt.Println("ACTION: " + action.String())
				fmt.Println(" ")
				fmt.Println(" ")

				intent := gjson.Get(string(body), "queryResult.intent.displayName")

				value := gjson.Get(string(body), "queryResult.parameters."+intent.String())

				fmt.Println(" ")
				fmt.Println("INTENT: " + intent.String())
				fmt.Println("VALUE: " + value.String())
				fmt.Println(" ")
				fmt.Println(" ")
				//
				// Update the Doc:
				updateDocument(trimLeftChars(docid, 12), intent.String(), value.String())
				//
				//
				//
				http.ServeFile(w, r, "formfield_action_response.json")
			} else {
				fmt.Println("action not found")

				http.ServeFile(w, r, "wildcard_response.json")
			}

			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			println(action.String())

		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}

	}

}

type WebhookRequest struct {
	ResponseID  string `json:"responseId"`
	QueryResult struct {
		QueryText  string `json:"queryText"`
		Action     string `json:"action"`
		Parameters struct {
			Template string `json:"template"`
		} `json:"parameters"`
		AllRequiredParamsPresent bool   `json:"allRequiredParamsPresent"`
		FulfillmentText          string `json:"fulfillmentText"`
		FulfillmentMessages      []struct {
			Text struct {
				Text []string `json:"text"`
			} `json:"text"`
		} `json:"fulfillmentMessages"`
		OutputContexts []struct {
			Name          string `json:"name"`
			LifespanCount int    `json:"lifespanCount,omitempty"`
			Parameters2   struct {
				Template         string `json:"template"`
				TemplateOriginal string `json:"template.original"`
			} `json:"parameters,omitempty"`
			Parameters3 struct {
				NoInput          int    `json:"no-input"`
				NoMatch          int    `json:"no-match"`
				Template         string `json:"template"`
				TemplateOriginal string `json:"template.original"`
			} `json:"parameters,omitempty"`
		} `json:"outputContexts"`
		Intent struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
		} `json:"intent"`
		IntentDetectionConfidence int    `json:"intentDetectionConfidence"`
		LanguageCode              string `json:"languageCode"`
	} `json:"queryResult"`
	OriginalDetectIntentRequest struct {
		Source  string `json:"source"`
		Payload struct {
		} `json:"payload"`
	} `json:"originalDetectIntentRequest"`
	Session string `json:"session"`
}
