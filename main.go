package tribes

import (
	"./nntp/"
	"./tools/"
	"./upnp/"
	"log"
	"os"
)

// No root. End of story

func init() {

	log.Printf("\n\n[OMG] %s", "AVERNO starts now!")
	go tools.SetLogFolder()

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
