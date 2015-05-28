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
	"strconv"
	"strings"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func ReadIpFromHost() string {

	conn, err := net.Dial("udp", "255.255.255.255:80")
	if err != nil {
		log.Printf("[TOOLS] SYSADMIIIIIN : cannot use UDP")
		return "127.0.0.1"
	}
	conn.Close()
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

func TheFileExists(filename string) bool {

	list, err := filepath.Glob(filename)
	if err != nil {
		return false
	}
	return list != nil

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

func RandomIPAddress() string {

	var IPfields string
	rand.Seed(time.Now().Unix())
	IPfields += strconv.Itoa(rand.Intn(254)) + "."
	IPfields += strconv.Itoa(rand.Intn(254)) + "."
	IPfields += strconv.Itoa(rand.Intn(254)) + "."
	IPfields += strconv.Itoa(rand.Intn(254))

	return IPfields

}

func ShortenString(s string, i int) string {
	runes := []rune(s)
	if len(runes) > i {
		return string(runes[:i])
	}
	return s
}
