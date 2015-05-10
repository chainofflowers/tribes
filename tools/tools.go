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

type tribesfile struct {
	filename string
	logfile  *os.File
}

func init() {

	var mylogfile tribesfile
	mylogfile.SetLogFolder()
	go mylogfile.RotateLogFolder()

}

func ReadIpFromHost() string {

	conn, err := net.Dial("udp", "example.com:80")
	if err != nil {
		log.Printf("[TOOLS] SYSADMIIIIIN : cannot use UDP")
		return "0.0.0.0"
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

// sets the log folder

func (f tribesfile) RotateLogFolder() {

	log.Println("[TOOLS] LogRotation engine started")

	for {

		time.Sleep(1 * time.Hour)
		if f.logfile != nil {
			err := f.logfile.Close()
			log.Println("[TOOLS] close logfile returned: ", err)
		}

		f.SetLogFolder()

	}

}

func (f tribesfile) SetLogFolder() {

	const layout = "2006-Jan-02.15"

	orario := time.Now()

	var user_home = GetHomeDir()
	f.filename = filepath.Join(user_home, "News", "logs", "averno."+orario.Format(layout)+"00.log")
	log.Println("[TOOLS] Logfile is: " + f.filename)

	f.logfile, _ = os.Create(f.filename)

	log.SetPrefix("TRIBES> ")
	log.SetOutput(f.logfile)

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
