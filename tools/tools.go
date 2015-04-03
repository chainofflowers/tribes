package tools

import (
	"log"
	"net"
	"os"
	"math/rand"
	"time"
)

var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")


func ReadIpFromHost() (string) {

	host, err := os.Hostname()

	if  err == nil {
		log.Printf("[INFO] Own Hostname is: %s", host)
	} else {
		log.Printf("[WTF] Can't get my own hostname? SYSADMIN!")
		panic(err.Error())
	}

	addrs, err := net.LookupIP(host)
	if  err == nil {
		log.Printf("[INFO] Own IP is: %s", addrs[0].String())
	} else {
		log.Printf("[WTF] Can't get my own IP? SYSADMIN!")
		panic(err.Error())
	}
	return addrs[0].String()
}


func RandSeq(n int) string {
	  rand.Seed(time.Now().UTC().UnixNano())
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
