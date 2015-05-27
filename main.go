package main

import (
	"./3be/"
	"./nntp/"
	"./tools/"
	"log"
	"os"
)

// No root. End of story

func init() {

	if (os.Getuid() == 0) || (os.Getgid() == 0) {
		log.Printf("[OMG] AAAARGH! ROOT! ROOT! ROOOOOT! ")
		os.Exit(1)
	}

	tools.Log_Engine_Start()
	tribe.AES_Engine_Start()
	nntp.NNTP_Engine_Start()

}

// main will only manage local data

func main() {

	log.Println("[OMG] AVERNO starts now!")

	select {}

	os.Exit(0)

}
