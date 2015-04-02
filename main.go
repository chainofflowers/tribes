package main

import (
	"fmt"
	"github.com/prestonTao/upnp"
	"log"
	"os"
)

// just to avoid running it as root

func init() {

	if (os.Getuid() == 0) || (os.Getgid() == 0) {
		fmt.Println("AAAARGH! ROOT! ROOT! ROOOOOT! ")
		fmt.Println("This is not a tree! We need no roots!")
		os.Exit(1)
	}

}

func main() {

	NNTP_UPnP_open()

	os.Exit(0)

}

func NNTP_UPnP_open() {

	mapping := new(upnp.Upnp)
	if err := mapping.AddPortMapping(11119, 11119, "TCP"); err == nil {
		log.Printf("[INFO] %s", "UPnP redirect 11119 successful")
	} else {
		log.Printf("[WARNING] %s", "UPnP redirect failed for port 11119: router is UPnP-agnostic")
	}
}
