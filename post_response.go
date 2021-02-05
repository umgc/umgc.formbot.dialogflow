package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func basicAuth(h http.HandlerFunc) http.HandlerFunc {
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
			http.Error(w, "Not authorized", 401)
			return
		}

		if pair[0] != "test" || pair[1] != "test" {
			http.Error(w, "Not authorized", 401)
			return
		}

		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "webhook.json")
		case "POST":

			/*if r.Header.Get("Content-Type") != "" {
				value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
				if value != "application/json" {
					msg := "Content-Type header is not application/json"
					http.Error(w, msg, http.StatusUnsupportedMediaType)
					return
				}
			}*/

			decoder := json.NewDecoder(r.Body)
			var t test_struct
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			//log.Println(t.Test)
			//jsonResponse(w, http.StatusCreated, "File uploaded successfully!.")
			/*
				// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
				if err := r.ParseForm(); err != nil {
					fmt.Fprintf(w, "ParseForm() err: %v", err)
					return
				}
				fmt.Fprintf(w, "API repsonse success:) r.PostFrom = %v\n", r.PostForm)
				name := r.FormValue("name")
				address := r.FormValue("address")
				fmt.Fprintf(w, "Name = %s\n", name)
				fmt.Fprintf(w, "Address = %s\n", address)
			*/
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}

		//basic response
		//	jsonResponse(w, http.StatusCreated, "File uploaded successfully123123!.")

		//server a json file:
		http.ServeFile(w, r, "webhook_response.json")

	}
}
