package tools

import (
	"log"
	"math/rand"
	"net"
	"os"
    "os/user"
	"time"
    "bufio"
    "fmt"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func ReadIpFromHost() string {

	host, err := os.Hostname()

	if err == nil {
		log.Printf("[INFO] Own Hostname is: %s", host)
	} else {
		log.Printf("[WTF] Can't get my own hostname? SYSADMIN!")
		os.Exit(1)
	}

	addrs, err := net.LookupIP(host)
	if err == nil {
		log.Printf("[INFO] Own IP is: %s", addrs[0].String())
	} else {
		log.Printf("[WTF] Can't get my own IP? SYSADMIN!")
		os.Exit(1)
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



func TheFileExists(filename string) (bool, error) {
  _, err := os.Stat(filename)
  if os.IsNotExist(err) {
    return false, nil
  }
  return err != nil, err
}


// just gets the home directory. to be moved in "tools"



func GetHomeDir() (string) {

    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
        log.Printf("[WTF] can't get homedir for user! SYSADMIIIN!"  )
        return "/tmp"
    } else {
    return usr.HomeDir
    }
}



func WriteMessages(lines []string, path string) error {
   file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  return w.Flush()
}
