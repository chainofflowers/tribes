package main

import (
	"log"
	"os"
	"tribes/3be"
	"tribes/cripta"
	"tribes/nntp"
	"tribes/tools"
	"tribes/upnp"
)

//
// No root. End of story
//
func init() {

	if (os.Getuid() == 0) || (os.Getgid() == 0) {
		log.Printf("[OMG] AAAARGH! ROOT! ROOT! ROOOOOT! ")
		os.Exit(1)
	}

	// now start all the engines.

	tools.Log_Engine_Start()
	cripta.GPG_Engine_Start()
	nntp.NNTP_Engine_Start()
	tribe.TribeEngineStart()
	upnp.UPNP_Engine_Start()

}

// main will only wait , doing nothing
func main() {

	log.Println("[OMG] TRIBES starts now!")

	select {}

	os.Exit(0)

}
