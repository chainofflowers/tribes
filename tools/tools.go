package tools

import (
	"log"
	"net"
	"os"
)

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
