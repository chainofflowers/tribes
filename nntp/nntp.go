package nntp

import (
	"bufio"
	"log"
	"net"
	"regexp"
    "strings"
    "../backend/"
)

var capab_out string = "101 Capability list:\nVERSION 2\nREADER\nPOST\nIHAVE\nLIST ACTIVE NEWSGROUPS OVERVIEW.FMT\n"


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


    var current_group string = "garbage"
// for future use    var current_messg string = "null"

	remote_client := conn.RemoteAddr()
    greetings := "200 averno.node AVERNO Version 01 beta, S0, posting OK"
    conn.Write([]byte(greetings + "\n"))
	for {

		linea, _,_ := bufio.NewReader(conn).ReadLine()

		message := string(linea)


		// decides WTF to do with the string

		if matches, _ := regexp.MatchString("(?i)^QUIT.*", message); matches == true {
			log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
            conn.Write([]byte("205 closing connection - goodbye!"))
			conn.Close()
            break
		}

		if matches, _ := regexp.MatchString("(?i)^GROUP[ ].+", message); matches == true {
            log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
            sinta := strings.Split(message," ")
            current_group = sinta[1]
            conn.Write([]byte(backend.ResponseToNNTPGROUP(current_group)))
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^LIST.*", message); matches == true {
            conn.Write([]byte("215 list of newsgroups follows"+ "\n"))
			log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
            backend.Trasmit_Active_NG(conn)
            conn.Write([]byte("."+ "\n"))
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^HEAD.*", message); matches == true {
			log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^BODY.*", message); matches == true {
			log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^ARTICLE.*", message); matches == true {
			log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^POST.*", message); matches == true {
			log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
            if current_group == "garbage" {
                                            conn.Write([]byte("412 no newsgroup has been selected\n"))
                                            continue
                                                       }
            backend.NNTP_POST_ReadAndSave(conn , current_group)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^IHAVE.*", message); matches == true {
			log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^CAPABILITIES.*", message); matches == true {
			log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
			conn.Write([]byte(capab_out))
			continue

		}
		if matches, _ := regexp.MatchString("(?i)^MODE.*READER.*", message); matches == true {
            conn.Write([]byte(greetings + "\n"))
			log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^AUTHINFO.*", message); matches == true {
			log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^NEWGROUPS.*", message); matches == true {
			log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
            conn.Write([]byte("231 New newsgroups since whatever follow"+ "\n"))
            backend.Trasmit_New_NG(conn)
            conn.Write([]byte("\n."+ "\n"))
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^OVER.*", message); matches == true {
			log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
            conn.Write([]byte("502 no permission, and BTW not a RFC977 command\n"))
			continue
		}
		if matches, _ := regexp.MatchString("(?i)^XOVER.*", message); matches == true {
			log.Printf("[INFO] NNTP %s from %s ", message,  remote_client)
            conn.Write([]byte("502 no permission, and BTW not a RFC977 command\n"))
			continue
		}

		
        if message == "" {continue}

        log.Printf("[INFO] NNTP BULLSHIT >%s< from %s ", message,  remote_client)



	}
	conn.Close()
}
