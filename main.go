package main

import (

    "./nntp/"
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

	go upnp.AllUpnpOpen()


}


// main will only manage local data

func main() {
    
    nntp.NNTP_Frontend()

//	select {}

	os.Exit(0)



}
