package main

import (
	"log"
	"os"
	"strconv"
)

var EmailUser = os.Getenv("EMAIL_USER")
var EmailPass = os.Getenv("EMAIL_PASS")
var EmailHost = os.Getenv("EMAIL_HOST")
var EmailPort = os.Getenv("EMAIL_PORT")
var EmailPortInt int

var CfgAddr = ":3000"

var CfgMgoUrl = "mongodb://0.0.0.0:27017/puppy"

func CfgInit() {
	var port string
	var addr string

	// Email ENV variables
	if EmailUser == "" || EmailPass == "" {
		log.Println("WARNING: Email variables are not in scope")
	}

	if EmailHost == "" {
		log.Println("No email host")
		EmailHost = "smtp.gmail.com"
	}

	if EmailPort == "" {
		log.Println("No email port")
		EmailPortInt = 587
	} else {
		var err error
		EmailPortInt, err = strconv.Atoi(EmailPort)
		if err != nil {
			log.Println("Invalid email port")
			EmailPortInt = 587
		}
	}

	// Mongo ENV variables
	port = os.Getenv("MONGO_PORT_27017_TCP_PORT")
	addr = os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	if port != "" && addr != "" {
		CfgMgoUrl = "mongodb://" + addr + ":" + port + "/puppy"
	}
}
