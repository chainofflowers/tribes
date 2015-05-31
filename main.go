package main

import (
	"log"
	"os"
	"tribes/3be"
	"tribes/cripta"
	"tribes/nntp"
	"tribes/tools"
)

// No root. End of story
// This is only to force all engines to start the init()
func init() {

	if (os.Getuid() == 0) || (os.Getgid() == 0) {
		log.Printf("[OMG] AAAARGH! ROOT! ROOT! ROOOOOT! ")
		os.Exit(1)
	}

	// now start all the engines asyncronously.

	tools.Log_Engine_Start()
	cripta.AES_Engine_Start()
	nntp.NNTP_Engine_Start()
	tribe.Tribe_Engine_Start()

}

// main will only wait , doing nothing
func main() {

	log.Println("[OMG] TRIBES starts now!")

	select {}

	os.Exit(0)

}
