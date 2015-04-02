package main

import (
	"fmt"
	"log"
	"os"
	"./tools/"
	"./upnp/"
)

// No root. End of story

func init() {

	if (os.Getuid() == 0) || (os.Getgid() == 0) {
		log.Printf("[OMG] %s", "AAAARGH! ROOT! ROOT! ROOOOOT! ")
		os.Exit(1)
	}

}

func main() {

	tools.ReadIpFromInterface()
	fmt.Println("now by host")
	tools.ReadIpFromHost()
	upnp.AllUpnpOpen()
	os.Exit(0)

}
