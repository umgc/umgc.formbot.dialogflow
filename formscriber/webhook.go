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

			fmt.Println("\nparsing body!")

			// parse out the intent/action
			action := gjson.Get(string(body), "queryResult.action")

			// parse out the source:
			//DIALOGFLOW_CONSOLE
			//google
			source := gjson.Get(string(body), "originalDetectIntentRequest.source")

			//Tip: Struct is only made for easy JSON node parsing using JsonPath
			//Tip: var request WebhookRequest
			//
			// 1) If ACTION = New report with URL then parse URL and get DocID to get all fields from template then build Dialogflow intents with batchUpdate
			//var request WebhookRequest
			//request.Session
			if action.String() == "NewReportURL" {
				// parse out the url.
				url := gjson.Get(string(body), "queryResult.parameters.template")

				fmt.Println("\nURL: " + url.String())

				//Check that user really gives a valid Google Doc URL, otherwise return back response it is not a valid Google Doc!
				//regex, _ := regexp.Compile("(docs.google.com/document/d/([a-zA-Z0-9-_]+))")
				//checkurl := regex.FindString(url.String())
				matched, err := regexp.MatchString("(docs.google.com/document/d/([a-zA-Z0-9-_]+))", url.String())
				fmt.Println(matched, err)
				if matched == false {
					var wrongurl string
					if source.String() != "google" {
						wrongurl = `{
						"fulfillmentMessages": [
						  {
							"text": {
							  "text": [
								"That is an incorrect url. Please visit formscriber.com to learn how to use this app first. Good bye!"
							  ]
							}
						  }
						]
					  }`
					} else {
						wrongurl = `{
						"payload": {
						  "google": {
							"expectUserResponse": true,
							"richResponse": {
							  "items": [
								{
								  "simpleResponse": {
									"textToSpeech": "That is an incorrect url. Please visit formscriber.com to learn how to use this app first. Good bye!",
									"displayText": "That is an incorrect url. Please visit formscriber.com to learn how to use this app first. Good bye!"
								  }
								}
							  ]
							}
						  }
						}
					  }`
					}
					fmt.Fprintf(w, wrongurl)
				} else {
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
					//fmt.Println(getDocument(trimLeftChars(docid, 12)))
					docid = trimLeftChars(docid, 12)

					//CHeck if there are fields to update
					//doc_access := getDocument(docid)

					// store all fields from the doc in a list
					// store the docid to check for errors and no fields found
					//fieldList, fields := getDocument(docid)
					if getDocument(docid) == "error" {
						//fieldList = "error"
						//fields = fieldList
						//if fields == "error" {

						fmt.Println("action not found")
						var response string
						if source.String() != "google" {
							response = `{
					"fulfillmentMessages": [
					  {
						"text": {
						  "text": [
							"There are no fields to update. Please check your Document or visit www.formscriber.com for info! Goodbye!"
						  ]
						}
					  }
					]
				  }`
						} else {
							response = `{
							"payload": {
							  "google": {
								"expectUserResponse": true,
								"richResponse": {
								  "items": [
									{
									  "simpleResponse": {
										"textToSpeech": "There are no fields to update. Please check your Document and return! Goodbye!",
										"displayText": "There are no fields to update. Please check your Document and return! Goodbye!"
									  }
									}
								  ]
								}
							  }
							}
						  }`
						}
						fmt.Fprintf(w, response)
					} else {

						//
						//
						session := gjson.Get(string(body), "session")

						fmt.Println("\nSession: " + session.String())
						var response string
						if source.String() != "google" {
							response = `{
							"fulfillmentMessages": [
							  {
								"text": {
								  "text": [
									"Thanks, I was able to find your form! Which field would you like to fill out? Or you can print your report by saying print."
								  ]
								}
							  }
							]
						  }`
						} else {
							response = `{
						"payload": {
						  "google": {
							"expectUserResponse": true,
							"richResponse": {
							  "items": [
								{
								  "simpleResponse": {
									"textToSpeech": "Thanks, I was able to find your form! Which field would you like to fill out? Or you can print your report by saying print.",
									"displayText": "Thanks, I was able to find your form! Which field would you like to fill out? Or you can print your report by saying print."
								  }
								}
							  ]
							}
						  }
						}
					  }`
						}
						fmt.Fprintf(w, response)

						//
						//http.ServeFile(w, r, "url_action_response.json")
					}
				}
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
				//make all lowercase

				fmt.Println(" ")
				fmt.Println("INTENT: " + intent.String())
				fmt.Println("VALUE: " + value.String())
				fmt.Println("DOCid: " + docid)
				fmt.Println("DOCid with trim: " + trimLeftChars(docid, 12))
				fmt.Println(" ")
				fmt.Println(" ")
				//
				// Update the Doc:
				//updateDocument(trimLeftChars(docid, 12), intent.String(), value.String()) //Removed 9MAR because trim is not needed
				//
				//strings.ToLower(value.String()
				//!"!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!TODO"
				updateDocument(docid, intent.String(), value.String())
				//
				//
				//
				var response string
				if source.String() != "google" {
					response = `{
					"fulfillmentMessages": [
					  {
						"text": {
						  "text": [
							"` + value.String() + ` recorded! Which field would like me to fill out next? Or if you're finished with your report just say print."
						  ]
						}
					  }
					]
				  }`
				} else {
					response = `{
							"payload": {
							  "google": {
								"expectUserResponse": true,
								"richResponse": {
								  "items": [
									{
									  "simpleResponse": {
										"textToSpeech": "` + value.String() + ` recorded! Which field would like me to fill out next? Or if you're finished with your report just say print.",
										"displayText": "` + value.String() + ` recorded! Which field would like me to fill out next? Or if you're finished with your report just say print."
									  }
									}
								  ]
								}
							  }
							}
						  }`
				}
				fmt.Fprintf(w, response)

				//http.ServeFile(w, r, "formfield_action_response.json")
			} else if action.String() == "end" {
				fmt.Println("\nACTION: " + action.String())

				var response string

				if source.String() != "google" {

					print := gjson.Get(string(body), "queryResult.outputContexts.1.parameters.template")
					fmt.Println("\nACTION: " + print.String())
					//pdf := print.String()
					//amend file to download pdf
					pdf := strings.ReplaceAll(print.String(), "/edit", "")
					pdf = pdf + "/export?format=pdf"
					fmt.Println("\n PDF: " + pdf)

					response = `{
					"fulfillmentMessages": [
					  {
						"card": {
						  "title": "Form Scriber",
						  "subtitle": "Report",
						  "imageUri": "https://formscriber.com/assets/media/FormScriber.com-logo/default.png",
						  "buttons": [
							{
							  "text": "Download",
							  "postback": "` + pdf + `"
							}
						  ]
						}
					  },
					  {
						"text": {
						  "text": [
							"Thank you for using Form Scriber! Here is your document. Good bye!"
						  ]
						}
					  },
					  {
						"text": {
						  "text": [
							"` + pdf + `"
						  ]
						}
					  }
					]
				  }`
				} else {

					print := gjson.Get(string(body), "queryResult.outputContexts.5.parameters.template")
					fmt.Println("\nACTION: " + print.String())
					//pdf := print.String()
					//amend file to download pdf
					pdf := strings.ReplaceAll(print.String(), "/edit", "")
					pdf = pdf + "/export?format=pdf"
					fmt.Println("\n PDF: " + pdf)

					fmt.Println("\n pdf was printed!:--> " + pdf)
					response = `{
						"payload": {
						  "google": {
							"expectUserResponse": false,
							"richResponse": {
							  "items": [
								{
								  "simpleResponse": {
									"textToSpeech": "Report Print."
								  }
								},
								{
								  "basicCard": {
									"title": "Form Scriber",
									"subtitle": "Print Report",
									"formattedText": "` + pdf + `",
									"image": {
									  "url": "https://formscriber.com/assets/media/FormScriber.com-logo/default.png",
									  "accessibilityText": "Form Scriber"
									},
									"buttons": [
									  {
										"title": "Report PDF",
										"openUrlAction": {
										  "url": "` + pdf + `"
										}
									  }
									],
									"imageDisplayOptions": "CROPPED"
								  }
								},
								{
								  "simpleResponse": {
									"textToSpeech": "Here is your report. Thank you for using Form Scriber! Goodbye!"
								  }
								}
							  ]
							}
						  }
						}
					  }`
				}
				/*	response = `{
						"title": "Report",
						"openUrlAction": {
						"url": www.123123123.com
						}
					  }`
				}*/

				fmt.Println("\n RESPONSE BACK TO GOOGLE: " + response)
				fmt.Fprintf(w, response)

			} else {
				fmt.Println("action not found")
				var response string
				if source.String() == "DIALOGFLOW_CONSOLE" {
					response = `{
					"fulfillmentMessages": [
					  {
						"text": {
						  "text": [
							"Thank you for using Form Scriber! Good bye!"
						  ]
						}
					  }
					]
				  }`
				} else {
					response = `{
							"payload": {
							  "google": {
								"expectUserResponse": true,
								"richResponse": {
								  "items": [
									{
									  "simpleResponse": {
										"textToSpeech": "Thank you for using Form Scriber! Good bye!",
										"displayText": "Thank you for using Form Scriber! Good bye!"
									  }
									}
								  ]
								}
							  }
							}
						  }`
				}
				fmt.Fprintf(w, response)

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
