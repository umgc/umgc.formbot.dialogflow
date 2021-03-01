package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Tokenresponse struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
	Scope        string `json:"scope"`
	Token_type   string `json:"token_type"`
}

var err error

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

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}

func about(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/about" {
		errorHandler(res, req, http.StatusNotFound)

		//log ip
		//log.Println(req.RemoteAddr)

		return
	}
	log.Println("about was accesssed")
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
	log.Println("successfully served about!")

	http.ServeFile(res, req, "about.html")
}

func index(res http.ResponseWriter, req *http.Request) {

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

	http.ServeFile(res, req, "index.html")
}

// RESTful services
type Articles struct {
	Type string    `json:"type"`
	D    []Article `json:"d"`
}
type Article struct {
	ID       string   `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Desc     string   `json:"description,omitempty"`
	URL      string   `json:"URL"`
	KeyWords *KeyWord `json:"keywords,omitempty"`
}
type KeyWord struct {
	Name string `json:"name,omitempty"`
}

func GetArticleEndPoint(w http.ResponseWriter, request *http.Request) {
	//	REST endpoint to get articles
	//decoder := json.NewDecoder(request.Body)

	jsonFile, err := os.Open("DATA/GetHelpArticleList.JSON")

	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var articles Articles

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &articles)

	/*	fmt.Println("Data Type: " + articles.Type)
		for i := 0; i < len(articles.D); i++ {
			fmt.Println("Name: " + articles.D[i].Name)
		}//*/

	json.NewEncoder(w).Encode(articles)
	/*
		var numsData numbers
		var numsResData numsResponseData

		decoder.Decode(&numsData)

		numsResData = process(numsData)//*/
	/*	fmt.Println(jsonFile)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonFile); err != nil {
			panic(err)
		} //*/
}

func main() {

	if err != nil {
		panic(err.Error())
	}
	log.Println("Engine running...")

	http.HandleFunc("/api", use(myHandler, basicAuth))
	http.HandleFunc("/about", about)
	http.HandleFunc("/", index)
	//http.HandleFunc("/google259e7adf5a143f76.html", googleSearchConsole)
	//http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("assets/css/"))))
	http.Handle("/assets/js/", http.StripPrefix("/assets/js/", http.FileServer(http.Dir("assets/js/"))))
	http.Handle("/assets/images/", http.StripPrefix("/assets/images/", http.FileServer(http.Dir("images"))))
	//http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	//http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("vendor"))))
	//http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))

	http.Handle("/site/", http.StripPrefix("/site/", http.FileServer(http.Dir("site"))))
	http.Handle("/site/STYLE/", http.StripPrefix("/site/STYLE/", http.FileServer(http.Dir("site/STYLE/"))))
	http.Handle("/site/JS/", http.StripPrefix("/site/JS/", http.FileServer(http.Dir("site/JS/"))))
	http.Handle("/site/IMG/", http.StripPrefix("/site/IMG/", http.FileServer(http.Dir("site/IMG"))))

	http.HandleFunc("/getArticles", GetArticleEndPoint)

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

	// Start the HTTPS server in a goroutine
	/*
		if err := http.ListenAndServeTLS(":8080", "formscriber.com.pem", "formscriber.key", nil); err != nil {
			log.Fatal("failed to start server", err)
		}//*/

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("failed to start server", err)
	} //*/

	// Cerbot Free SSL instruction: https://certbot.eff.org/lets-encrypt/windows-other

}
