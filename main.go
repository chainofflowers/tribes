package main

import (
	"./legione/"
    "./nntp/"
	"./upnp/"
	"log"
	"os"
)

// Channels to connect nntp and wendy

var IncomingMessages = make(chan string,100) 
var AnnounceMessages = make(chan string,100)
var IncomingGroups   = make(chan string,100)
var AnnounceGroups   = make(chan string,100)


// No root. End of story

func init() {

	if (os.Getuid() == 0) || (os.Getgid() == 0) {
		log.Printf("[OMG] %s", "AAAARGH! ROOT! ROOT! ROOOOOT! ")
		os.Exit(1)
	}

	go upnp.AllUpnpOpen()
	go legion.Initialize()

}


// main will only manage local data

func main() {
    
    nntp.NNTP_Frontend()

//	select {}

	os.Exit(0)

}
