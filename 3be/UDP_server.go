package tribe

import (
	"../config/"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func Udp_Server() {

	udp_port := ":" + strconv.Itoa(config.GetClusterPort())

	ServerAddr, err := net.ResolveUDPAddr("udp", udp_port)
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 65507) // 0xffff - (sizeof(IP Header) + sizeof(UDP Header)) = 65535-(20+8) = 65507

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr, "bytes: ", n)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
