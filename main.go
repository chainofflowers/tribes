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

func main() {

	ReadIpFromInterface()
	fmt.Println("now by host")
	ReadIpFromHost()
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

func ReadIpFromHost() {
	host, _ := os.Hostname()
	log.Printf("[INFO] %s", host)
	addrs, _ := net.LookupIP(host)
	log.Printf("[INFO] %s", addrs)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			log.Printf("[INFO] %s", ipv4)
		}
	}

}

func ReadIpFromInterface() {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			log.Printf("[INFO] %s", ipnet.IP)
		}
	}

}
