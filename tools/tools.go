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

	host, _ := os.Hostname()
	log.Printf("[INFO] Own Hostname is: %s", host)
	addrs, _ := net.LookupIP(host)
	log.Printf("[INFO] Own IP is: %s", addrs[0].String())

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
