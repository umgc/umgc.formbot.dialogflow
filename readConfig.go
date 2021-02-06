package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type Configuration struct {
	Db struct {
		Host     string
		User     string
		Password string
		Database string
	}
	Listen struct {
		Address string
		Port    string
	}
	OutboundCall struct {
		CallerID  string
		Retries   int
		SpoolPath string
	}
	VmRoot string
}

func config() {
	c := flag.String("c", "config.json", "Specify the configuration file.")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		log.Fatal("can't open config file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}
	log.Println(Config.Db.Host)
	log.Println("log parsed!")

}
