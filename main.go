package main

import (
	"./legione/"
	"./upnp/"
	"log"
	"os"
)

// No root. End of story

func init() {

	if (os.Getuid() == 0) || (os.Getgid() == 0) {
		log.Printf("[OMG] %s", "AAAARGH! ROOT! ROOT! ROOOOOT! ")
		os.Exit(1)
	}

	upnp.AllUpnpOpen()

}

func main() {

	legion.Initialize()

	os.Exit(0)

}
