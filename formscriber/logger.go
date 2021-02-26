package main

import (
	"log"
	"time"
)

func logger(record string) {

	// log request by who(IP address)
	start := time.Now()
	//requesterIP := req.RemoteAddr
	log.Printf(
		"%s\t",
		//req.Method,
		//req.RequestURI,
		//requesterIP,
		time.Since(start),
	)
	//end log
	log.Println(record)

}
