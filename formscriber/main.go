package main

import (
	"fmt"
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

	http.ServeFile(res, req, "assets/robots.txt")

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

	http.ServeFile(res, req, "html/about.html")
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

	http.ServeFile(res, req, "html/index.html")
}

/* 2.0 RESTful services
 * the following are the REST request handlers
 */

/* Main Code
 * THis is where everything is activated
 */

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

	http.HandleFunc("/getArticles", GetArticlesEndPoint)
	http.HandleFunc("/getTeam", GetTeamEndPoint)
	http.HandleFunc("/getFAQs", GetFAQsEndPoint)

	//log file system
	fileName := "webrequests.log"
	// https://www.socketloop.com/tutorials/golang-how-to-save-log-messages-to-file
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()


	// Start the HTTPS server in a goroutine
	log.Fatal(http.ListenAndServe(":80", nil))

	if err := http.ListenAndServeTLS(":443", "formscriber.com.pem", "formscriber.key", nil); err != nil {
		log.Fatal("failed to start server", err)
	} //*/

	log.Println("Server running on http 80 and https 443")
	// Cerbot Free SSL instruction: https://certbot.eff.org/lets-encrypt/windows-other
	
	// direct all log messages to webrequests.log
	log.SetOutput(logFile)

}
