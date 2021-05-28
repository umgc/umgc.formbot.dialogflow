package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func createDynamicIntent(intent string) {

	token := getAccessToken()
	url := "https://dialogflow.googleapis.com/v2/projects/form-bot-1577a/agent/intents:batchUpdate"
	method := "POST"

	payload := strings.NewReader(`{
    
		"intentBatchInline": {
		"intents": [
			{
				"displayName": "` + intent + `",
				"priority": 500000,
				"webhookState": "WEBHOOK_STATE_ENABLED",
				"inputContextNames": [
					"projects/form-bot-1577a/agent/sessions/-/contexts/GoogleTemplate",
					"projects/form-bot-1577a/agent/sessions/-/contexts/DefaultWelcomeIntent-custom-followup"
				],
				"trainingPhrases": [
                    {
                        "type": "EXAMPLE",
                        "parts": [
                                 {
                                "text": "` + intent + `"
                                }
                            ]
                            }

                ],
				"action": "FillDocField",
				"parameters": [
					{
						"name": "b8bcf595-1a59-44d9-9fb3-c5f9e7bc959e",
						"displayName": "` + intent + `",
						"value": "$` + intent + `",
						"entityTypeDisplayName": "@sys.any",
						"mandatory": true
					}
				],
				"messages": [
					{
						"text": {
							"text": [
								"I didn't catch that. Could you please try again?"
							]
						}
					}
				],
				"parentFollowupIntentName": "projects/form-bot-1577a/agent/intents/fd80bbe1-072f-4d61-82c5-c491bd532166"
			}        
		]
	}
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		log.Println("Error : something terrible happen -> ", err)

	}
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Println("Error : something terrible happen -> ", err)

	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	fmt.Println("\n Dynamic Intent made: " + intent + "\n")
}
