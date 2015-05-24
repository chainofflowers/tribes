package tribe

import (
	"log"
	"net"
)

type TribeServer struct {
	TSAddr *net.UDPAddr
	TSConn *net.UDPConn
	TSPort string
	TSErr  error
}

type TribePayload struct {
	TPbuffer []byte
	TPsender *net.UDPAddr
	TPsize   int
	TPErr    error
}

func (this *TribeServer) Udp_Server() {

	this.TSAddr, this.TSErr = net.ResolveUDPAddr("udp", this.TSPort)
	if this.TSErr != nil {
		log.Printf("[UDP] Cannot resolve UDP address : %s", this.TSErr)

	}

	/* Now listen at selected port */
	this.TSConn, this.TSErr = net.ListenUDP("udp", this.TSAddr)

	if this.TSErr != nil {
		log.Printf("[UDP] Cannot bind on listening port : %s", this.TSErr)

	}

	defer this.TSConn.Close()

	// buf := make([]byte, 65507) // 0xffff - (sizeof(IP Header) + sizeof(UDP Header)) = 65535-(20+8) = 65507

	for {
		var Payload TribePayload
		Payload.TPbuffer = make([]byte, 65507) // 0xffff - (sizeof(IP Header) + sizeof(UDP Header)) = 65535-(20+8) = 65507
		Payload.TPsize, Payload.TPsender, Payload.TPErr = this.TSConn.ReadFromUDP(Payload.TPbuffer)

		if Payload.TPErr != nil {
			log.Printf("[UDP] Cannot read on UDP socket : %s", Payload.TPErr)
		} else {
			log.Printf("[UDP] Received ", string(Payload.TPbuffer[0:Payload.TPsize]), " from ", Payload.TPsender, "bytes: ", Payload.TPsize)
		}

		//		Tribes_Interpreter(string(buf[0:n]))

	}
}
