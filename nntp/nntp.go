package nntp

import (
	"bufio"

	"log"
	"net"
	"regexp"
)

var capab_out string = "101 Capability list:\nVERSION 2\nREADER\nPOST\nIHAVE\nOVER\nXOVER\nLIST ACTIVE NEWSGROUPS OVERVIEW.FMT\n"

func NNTP_Frontend() {

	// setting up the tcp connection

	ln, err := net.Listen("tcp", "127.0.0.1:11119")
	if err == nil {
		log.Printf("[INFO] TCP listening at %s ", "127.0.0.1:11119")
	} else {
		log.Printf("[WTF] TCP CANNOT listen at %s. SYSADMIIIIN!!", "127.0.0.1:11119")
	}

	defer ln.Close()

	for {

		// start listening at it

		server, err := ln.Accept()
		tcp_client := server.RemoteAddr()

		if err == nil {

			log.Printf("[INFO] NNTP accepted connection from %s ", tcp_client)
		} else {
			log.Printf("[WTF] NNTP something went wrong at %s. SYSADMIIIIN!!", "127.0.0.1:11119")
		}

		go NNTP_Interpret(server)

	}

}

func NNTP_Interpret(conn net.Conn) {

	remote_client := conn.RemoteAddr()

	for {

		linea, _ := bufio.NewReader(conn).ReadString('\n')

		message := string(linea)
		// output message received

		log.Printf("[DEBUG] NNTP %s from %s ", message, remote_client)

		// decides WTF to do with the string

		if matches, _ := regexp.MatchString("(?i)^QUIT.*", message); matches == true {
			log.Printf("[INFO] NNTP QUIT from %s ", remote_client)
			conn.Close()
		}

		if matches, _ := regexp.MatchString("(?i)^GROUP.*", message); matches == true {
			log.Printf("[INFO] NNTP GROUP from %s ", remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^LIST.*", message); matches == true {
			log.Printf("[INFO] NNTP LIST from %s ", remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^HEAD.*", message); matches == true {
			log.Printf("[INFO] NNTP HEAD from %s ", remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^BODY.*", message); matches == true {
			log.Printf("[INFO] NNTP BODY from %s ", remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^ARTICLE.*", message); matches == true {
			log.Printf("[INFO] NNTP ARTICLE from %s ", remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^POST.*", message); matches == true {
			log.Printf("[INFO] NNTP POST from %s ", remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^IHAVE.*", message); matches == true {
			log.Printf("[INFO] NNTP IHAVE from %s ", remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^CAPABILITIES.*", message); matches == true {
			log.Printf("[INFO] NNTP CAPABILITIES from %s ", remote_client)
			conn.Write([]byte(capab_out))
			continue

		}
		if matches, _ := regexp.MatchString("(?i)^MODE.*", message); matches == true {
			log.Printf("[INFO] NNTP MODE from %s ", remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^AUTHINFO.*", message); matches == true {
			log.Printf("[INFO] NNTP AUTHINFO from %s ", remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^NEWSGROUPS.*", message); matches == true {
			log.Printf("[INFO] NNTP NEWSGROUPS from %s ", remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^OVER.*", message); matches == true {
			log.Printf("[INFO] NNTP OVER from %s ", remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^XOVER.*", message); matches == true {
			log.Printf("[INFO] NNTP XOVER from %s ", remote_client)
			continue
		}

		log.Printf("[INFO] NNTP BULLSHIT %s , closing connection ", remote_client)
		conn.Close()
	}
}
