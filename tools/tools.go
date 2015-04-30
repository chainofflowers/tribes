package tools

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func ReadIpFromHost() string {

	conn, err := net.Dial("udp", "example.com:80")
	if err != nil {
		log.Printf("[TOOLS] SYSADMIIIIIN : cannot use UDP")
		return "0.0.0.0"
	}
	defer conn.Close()
	torn := strings.Split(conn.LocalAddr().String(), ":")
	return torn[0]
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

func GetHomeDir() string {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		log.Printf("[WTF] can't get homedir for user! SYSADMIIIN!")
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

func SetLogFolder() {

	var user_home = GetHomeDir()
	avernologfile := filepath.Join(user_home, "News", "averno.log")
	fmt.Println("Logfile is: " + avernologfile)

	f, err := os.OpenFile(avernologfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening " + avernologfile)
	}

	log.SetOutput(f)
}

// Checks if a string is in a slice

func GrepASlice(str string, list []string) bool {
	for _, v := range list {
		if strings.Contains(v, str) {
			return true
		}
	}
	return false
}

// find the position of a string in the slice

func StringPosInSlice(str string, list []string) int {
	for i, v := range list {
		if strings.Contains(v, str) {
			return i
		}
	}
	return -1
}
