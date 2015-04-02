package main

import (
	"fmt"
	"github.com/prestonTao/upnp"
	"log"
	"net"

	"os"
)

// No root. End of story

func init() {

	if (os.Getuid() == 0) || (os.Getgid() == 0) {
		fmt.Println("AAAARGH! ROOT! ROOT! ROOOOOT! ")
		fmt.Println("This is not a tree! We need no roots!")
		os.Exit(1)
	}

}

const (
	httpPortTCP = 8080
	numTarget   = 10
	exampleIH   = "deca7a89a1dbdc4b213de1c0d5351e92582f31fb" // ubuntu-12.04.4-desktop-amd64.iso
)

func main() {

	ReadIpFromInterface()
	NntpUpnpOpen()

	os.Exit(0)

}

func NntpUpnpOpen() {

	mapping := new(upnp.Upnp)
	if err := mapping.AddPortMapping(11119, 11119, "TCP"); err == nil {
		log.Printf("[INFO] %s", "UPnP redirect 11119 successful")
	} else {
		log.Printf("[WARNING] %s", "UPnP redirect failed for port 11119: router is UPnP-agnostic")
	}
}

func ReadIpFromInterface() {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {

			log.Printf("[INFO] %s", ipv4)
		}
	}

}
