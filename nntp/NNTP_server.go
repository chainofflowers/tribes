package nntp

import (
	"log"
	"net"
)

var CapabResponse string = "101 Capability list:\nVERSION 2\nREADER\nPOST\nSTAT\nXOVER\nOVER\nLIST ACTIVE NEWSGROUPS OVERVIEW.FMT\n"

func init() {

	go NNTP_Frontend()

}

func NNTP_Frontend() {

	// setting up the tcp connection

	ln, err := net.Listen("tcp", "127.0.0.1:11119")
	if err == nil {
		log.Printf("[NNTP] TCP listening at %s ", "127.0.0.1:11119")

	} else {
		log.Printf("[WTF] TCP CANNOT listen at %s. SYSADMIIIIN!!", "127.0.0.1:11119")
		return
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

		// start the NNTP interpreter in background.
		go NNTP_Interpret(server)

	}

}

func NNTP_Engine_Start() {

	log.Printf("[NNTP] NNTP Engine Starting")

}
