package backend

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
	to "tribes/tools"
)

// gets the active NG and sends them to the given sockets

func Trasmit_Active_NG(conn net.Conn) error {
	file, err := os.Open(to.ActiveNgFile)
	if err != nil {
		log.Printf("[BE-FS] can't open file %s", to.ActiveNgFile)
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
	file, err := os.Open(to.NewNgFile)
	if err != nil {
		log.Printf("[BE-FS] can't open file %s", to.NewNgFile)
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

func Transmit_Article(conn net.Conn, FileName string) {
	log.Printf("[BE-FS] asked to open %s", FileName)
	file, err := os.Open(FileName)
	if err != nil {
		log.Printf("[BE-FS] can't open file %s", FileName)

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
