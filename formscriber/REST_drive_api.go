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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func driveAuth(h http.HandlerFunc) http.HandlerFunc {
	//var driveid string
	fmt.Println("\nStarting Drive APi Authentication....\n")
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("\nDrive Auth.. Step 2.....\n")
		enableCors(&w)
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		fmt.Println("\nAuth step 1\n")
		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(w, "Not authorized", 401)
			fmt.Println("\nNot Authorized1\n")
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
			fmt.Println("\nNot Authorized2\n")
			return
		}
		fmt.Println("\nCorrect username and password!\n")
		switch r.Method {

		case "GET":
			fmt.Println("\nCheck GET method\n")
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println("\nCould not parse GET\n")
				fmt.Println(err)
				return
			}
			fmt.Println("\nGet was parsed\n")
			//

			fmt.Println("\nparsing DRIVE body!\n")
			fmt.Println(string(body))
			// parse out the intent/action
			if len(string(body)) < 1 {
				fmt.Println("Body empty! Default response")
				response := `{
					"kind": "drive#fileList",
					"incompleteSearch": false,
					"files": [
					 {
					  "kind": "drive#file",
					  "id": "1F_0unyG-HHrOqRj1o-QfEQDxmkl_xw5xSo4dVLyqvPA",
					  "name": "Test Account",
					  "mimeType": "application/vnd.google-apps.document"
					 }
					]
				   }`
				fmt.Fprintf(w, response)

				fmt.Println("\n RESPONSE BACK WEB APP: " + response)
				fmt.Fprintf(w, response)
			} else {

				driveUrl := gjson.Get(string(body), "driveUrl")
				fmt.Println("\nDrtive URL parsed\n")
				//drive.google.com/drive/u/0/folders/([a-zA-Z0-9-_]+)
				regex, _ := regexp.Compile("/folders/([a-zA-Z0-9-_]+)")
				fmt.Println("\nRegex step1\n")
				//fmt.Println(r.FindString(url.String()))
				driveid := regex.FindString(driveUrl.String())
				fmt.Println("\nRegex find url\n")

				// Extract Drive ID:

				fmt.Println("\nDriveID without trim: " + driveid)
				fmt.Println("\nDriveID with trim: " + trimLeftChars(driveid, 9))

				driveid = trimLeftChars(driveid, 9)
				fmt.Println("\nDriveID Final: " + driveid)

				url := "https://www.googleapis.com/drive/v3/files?q='" + driveid + "'%20in%20parents"

				fmt.Println("\nDRIVE URL: " + url)

				token := getAccessToken()
				fmt.Println("\nTOKEN: " + token)

				//Initiate Google Drive Call
				method := "GET"
				fmt.Println("\nCall Google step1\n")
				payload := strings.NewReader(``)

				client := &http.Client{}
				req, err := http.NewRequest(method, url, payload)
				fmt.Println("\nCall Google step2\n")
				if err != nil {
					fmt.Println("\nError calling google API1\n")
					fmt.Println(err)
					return
				}
				fmt.Println("\nGoogel APi called1\n")
				req.Header.Add("Authorization", "Bearer "+token)
				req.Header.Add("Content-Type", "application/json")
				fmt.Println("\nGoogel APi called2\n")
				res, err := client.Do(req)
				if err != nil {
					fmt.Println("\nError calling google API2\n")
					fmt.Println(err)
					return
				}
				defer res.Body.Close()
				fmt.Println("\nReturn Response back to client1\n")
				body, err = ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Println("\nERROR Return Response back to client1\n")
					fmt.Println(err)
					return
				}
				fmt.Println(string(body))
				fmt.Fprintf(w, string(body))
				fmt.Println("\nGet Drive Files complete!\n")
				return
			}
		case "POST":
			fmt.Println("\nPOST FOUND ERROR\n")
			fmt.Fprintf(w, "Sorry, only GET methods are supported.")
			return

		default:
			fmt.Println("\nUnhandled ERROR!\n")
			fmt.Fprintf(w, "Sorry, only GET methods are supported.")
			return
		}
		fmt.Println("\nProcess done1\n")
	}

}
