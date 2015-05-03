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

// sets the log folder

func RotateLogFolder() {

	log.Println("[TOOLS] LogRotation engine started")

	for {

		time.Sleep(1 * time.Hour)
		SetLogFolder()

	}

}

func SetLogFolder() {

	const layout = "2006-Jan-02.15"

	orario := time.Now()

	var user_home = GetHomeDir()
	avernologfile := filepath.Join(user_home, "News", "logs", "averno."+orario.Format(layout)+"00.log")
	log.Println("[TOOLS] Logfile is: " + avernologfile)

	f, err := os.OpenFile(avernologfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("[TOOLS] Error opening logfile: %s " + avernologfile)
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
