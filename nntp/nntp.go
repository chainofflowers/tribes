package nntp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings" // only needed below for sample processing
)

func NNTP_Frontend() {

	ln, err := net.Listen("tcp", "127.0.0.1:11119")
	if err == nil {
		log.Printf("[INFO] TCP listening at %s ", "127.0.0.1:11119")
	} else {
		log.Printf("[WTF] TCP CANNOT listen at %s. SYSADMIIIIN!!", "127.0.0.1:11119")
	}

	for {

		conn, err := ln.Accept()

		if err == nil {
			remote_client := conn.RemoteAddr()

			log.Printf("[INFO] NNTP accepted connection from %s ", remote_client)
		} else {
			log.Printf("[WTF] NNTP something went wrong at %s. SYSADMIIIIN!!", "127.0.0.1:11119")
		}

		// will listen for message to process ending in newline (\n)

		message, _ := bufio.NewReader(conn).ReadString('\n')

		// output message received

		fmt.Print("Message Received:", string(message))

		// sample process for string received

		newmessage := strings.ToUpper(message)

		// send new string back to client

		conn.Write([]byte(newmessage + "\n"))

		conn.Close()

	}
}
