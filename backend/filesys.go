package backend

import (
	"bufio"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"tribes/tools"
)

var (
	active_ng_file  string = "/News/groups/ng.active"
	new_ng_file     string = "/News/groups/ng.local"
	all_ng_file     string = "/News/groups/ng.all"
	messages_folder string = "/News/messages/"
)

// initializes everything

func init() {

	var user_home = tools.GetHomeDir()
	active_ng_file = filepath.Join(user_home, active_ng_file)
	new_ng_file = filepath.Join(user_home, new_ng_file)
	all_ng_file = filepath.Join(user_home, all_ng_file)
	messages_folder = filepath.Join(user_home, messages_folder)

	os.MkdirAll(filepath.Join(user_home, "News", "groups"), 0755)
	os.MkdirAll(filepath.Join(user_home, "News", "messages"), 0755)
	os.MkdirAll(filepath.Join(user_home, "News", "logs"), 0755)

	if tools.TheFileExists(active_ng_file) == false {
		log.Printf("[BE-FS] creating file %s", active_ng_file)
		os.Create(active_ng_file)
	}
	if tools.TheFileExists(new_ng_file) == false {
		log.Printf("[BE-FS] Creating File %s", new_ng_file)
		os.Create(new_ng_file)
	}
	if tools.TheFileExists(all_ng_file) == false {
		log.Printf("[BE-FS] Creating File %s", all_ng_file)
		os.Create(new_ng_file)
	}

}

// gets the active NG and sends them to the given sockets

func Trasmit_Active_NG(conn net.Conn) error {
	file, err := os.Open(active_ng_file)
	if err != nil {
		log.Printf("[BE-FS] can't open file %s", active_ng_file)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line string = scanner.Text()
		line = strings.Replace(line, "-", "_", -1)
		response := line + " " + GetLastNumByGroup(line) + " " + GetFirstNumByGroup(line) + " y"
		conn.Write([]byte(response + "\r\n"))
		log.Printf("[BE-FS] NNTP print: %s ", response)
	}
	file.Close()
	return scanner.Err()
}

// transmits NEW newgroups (here "local") to the given socket

func Trasmit_New_NG(conn net.Conn) error {
	file, err := os.Open(new_ng_file)
	if err != nil {
		log.Printf("[BE-FS] can't open file %s", new_ng_file)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line string = scanner.Text()
		conn.Write([]byte(line + "\r\n"))
		log.Printf("[BE-FS] NNTP print: %s ", line)
	}
	file.Close()
	return scanner.Err()
}

func Transmit_Article(conn net.Conn, file_id string) {
	log.Printf("[BE-FS] asked to open %s", file_id)
	file, err := os.Open(file_id)
	if err != nil {
		log.Printf("[BE-FS] can't open file %s", file_id)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := conn.Write([]byte(line + "\r\n"))
		log.Printf("[BE-FS] NNTP print: %s [%d BYTES]", line, n)
	}

	file.Close()

}
