package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	recaptcha "github.com/dpapathanasiou/go-recaptcha"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/sessions"
)

var err error
var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func CheckGoogleCaptcha(response string) bool {
	var googleCaptcha string = "6Lc3XT4UAAAAABac5-cbX23gBDnzUd9_TUYNbWQF"
	req, err := http.NewRequest("POST", "https://www.google.com/recaptcha/api/siteverify", nil)
	q := req.URL.Query()
	q.Add("secret", googleCaptcha)
	q.Add("response", response)
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	var googleResponse map[string]interface{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &googleResponse)
	return googleResponse["success"].(bool)
}

func errorHandler(res http.ResponseWriter, req *http.Request, status int) {
	res.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(res, "404 Page Not Found")

		//For now, dont server any html
		//http.ServeFile(res, req, "404.html")
	}
}

func robot(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "robots.txt")

}

// jsonResponse(w, http.StatusCreated, "File uploaded successfully!.")
func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}

func RequestLogger(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// log request by who(IP address)
	requesterIP := r.RemoteAddr

	log.Printf(
		"%s\t\t%s\t\t%s\t\t%v",
		r.Method,
		r.RequestURI,
		requesterIP,

		time.Since(start),
	)

}

func about(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/about" {
		errorHandler(res, req, http.StatusNotFound)

		//log ip
		//log.Println(req.RemoteAddr)

		return
	}
	log.Println("index was accesssed")
	// log request by who(IP address)
	start := time.Now()
	requesterIP := req.RemoteAddr
	log.Printf(
		"%s\t\t%s\t\t%s\t\t%v",
		req.Method,
		req.RequestURI,
		requesterIP,
		time.Since(start),
	)
	//end log
	log.Println("successfully served index!")

	http.ServeFile(res, req, "about.html")
}

func main() {

	//REad config file
	config()

	//Initialize captcha key
	recaptcha.Init("6Lc3XT4UAAAAABac5-cbX23gBDnzUd9_TUYNbWQF")


	if err != nil {
		panic(err.Error())
	}
	log.Println("Listening...")

	http.HandleFunc("/", use(myHandler, basicAuth))
	http.HandleFunc("/about", about)
	//http.HandleFunc("/google259e7adf5a143f76.html", googleSearchConsole)
	//http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("assets/css/"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	//http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	//http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("vendor"))))
	//http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))

	//log file system
	fileName := "webrequests.log"
	// https://www.socketloop.com/tutorials/golang-how-to-save-log-messages-to-file
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	// direct all log messages to webrequests.log
	log.SetOutput(logFile)

	// Start the HTTP server
	//write a small error catch

	/*if err := http.ListenAndServe(":80", nil); err != nil { //LIVE: 80
		log.Fatal("failed to start server", err)
	}*/

	// Start the HTTPS server in a goroutine
	if err := http.ListenAndServeTLS(":443", "edigenerator.com.pem", "edigenerator.key", nil); err != nil {
		log.Fatal("failed to start server", err)
	}

	// Cerbot Free SSL instruction: https://certbot.eff.org/lets-encrypt/windows-other
}

/* Author:  Caleb Crickette 2021

This basic template has a lot of ideas to quickly get working in Go which will serve easily HTTP/S website or RESTful API's.

Some good articles on Go:
https://stackoverflow.blog/2020/11/02/go-golang-learn-fast-programming-languages/
https://dev.to/techschoolguru/implement-restful-http-api-in-go-using-gin-4ap1
https://medium.com/helidon/can-java-microservices-be-as-fast-as-go-5ceb9a45d673

For all documentation visit www.golang.org

Inside there are examples of database inserts and reads as well as several examples of serving and serializing/unserializing JSON data.


HOW TO RUN AND COMPILE THE CODE TO AN .EXE (binary)
You need to have Go installed. (golang.org) then to test it open a terminal/cmd and type "go" you should get some information if not. Recheck installation.const
There are different ways to run Go code. While quickly developing it is better to just run the following command in your terminal or IDE:
"Go run whatever_the_name_of_your_program.go"
It will then run in your console.  If you want to create an executable and compile to machine code to deploy its very similar, but instead of RUN use "build":
"Go build whatever_the_name_of_your_program.go"
NOTE: if you want to deploy on other environments you need to set an environment variblae on your computer AND THEN run the  build command.
REMBMER to set it back if you are working on your local machine!

For  mac for example run the following code in your terminal console or cmd:
1. $ GOOS=darwin GOARCH=386
2. Go build whatever_the_name_of_your_program.go


Then if you want windows exe:
1. $ GOOS=windows GOARCH=amd64
2. Go build whatever_the_name_of_your_program.go

Sometimes Windows gives a problem so try the following:
1. $env:GOOS = "linux"
2. read variable: $env:GOOS
3. Go build whatever_the_name_of_your_program.go


For a complete list visit here:
https://github.com/ccrickette/Go/blob/master/ps_cross_compile.txt
https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04


TIP: if you do not want the console open just add the following to your run/build command:
-ldflags -H=windowsgui

*/
