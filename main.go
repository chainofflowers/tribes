package main

import (
	"./3be/"
	"./nntp/"
	"log"
	"os"
)

// No root. End of story

func init() {

	if (os.Getuid() == 0) || (os.Getgid() == 0) {
		log.Printf("[OMG] %s", "AAAARGH! ROOT! ROOT! ROOOOOT! ")
		os.Exit(1)
	}

}

// main will only manage local data

func main() {

	log.Println("[OMG] AVERNO starts now!")
	tribe.AES_Engine_Start()

	nntp.NNTP_Frontend()

	//	select {}

	os.Exit(0)

}
