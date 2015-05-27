package tribe

import (
	"../config"
	"../tools/"
	"log"
	"net"
	"strconv"
	"time"
)

type TribeServer struct {
	TSLAddr *net.UDPAddr // local udp address
	TSRAddr *net.UDPAddr // remote udp address
	TSConn  *net.UDPConn // connection in use
	TSRun   bool
	TSPort  string
	TSErr   error
}

type TribePayload struct {
	TPbuffer []byte
	TPsender *net.UDPAddr
	TPsize   int
	TPErr    error
}

func init() {

	var TribeSrv TribeServer
	TribeSrv.TSRun = false
	TribeSrv.TSPort = tools.ReadIpFromHost() + ":" + strconv.Itoa(config.GetClusterPort())
	go TribeSrv.RefreshPunchHole()

}

func Start3beEngine() {

	log.Printf("[UDP] Tribes Engine starting")

}

func (this *TribeServer) Udp_Server() {

	this.TSLAddr, this.TSErr = net.ResolveUDPAddr("udp", this.TSPort)

	if this.TSErr != nil {
		log.Printf("[UDP] Cannot resolve UDP address : %s", this.TSErr.Error())
	}

	/* Now listen at selected port */
	this.TSConn, this.TSErr = net.ListenUDP("udp", this.TSLAddr)

	if this.TSErr != nil {
		log.Printf("[UDP] Cannot bind on listening port : %s", this.TSErr.Error())
		this.TSRun = false
	} else {
		log.Printf("[UDP] Bound on listening port %s", this.TSPort)
		this.TSRun = true
	}

	defer this.TSConn.Close()

	// buf := make([]byte, 65507) // 0xffff - (sizeof(IP Header) + sizeof(UDP Header)) = 65535-(20+8) = 65507

	for {
		var Payload TribePayload
		Payload.TPbuffer = make([]byte, 65507) // 0xffff - (sizeof(IP Header) + sizeof(UDP Header)) = 65535-(20+8) = 65507
		Payload.TPsize, Payload.TPsender, Payload.TPErr = this.TSConn.ReadFromUDP(Payload.TPbuffer)

		if Payload.TPErr != nil {
			log.Printf("[UDP] Cannot read on UDP socket : %s", Payload.TPErr.Error())
		} else {
			log.Printf("[UDP] Received %d bytes", Payload.TPsize)
			log.Printf("[UDP] Received from: %v", Payload.TPsender)
			log.Printf("[UDP] Received string follows: %s", string(Payload.TPbuffer[0:Payload.TPsize]))
		}

		this.Tribes_Interpreter(Payload)

	}

}

func (this *TribeServer) OpenNatUDPport() {

	// this will be changed. Instead of using random IP, it will hit peers
	// doing a REGISTER

	this.TSRAddr, this.TSErr = net.ResolveUDPAddr("udp", tools.RandomIPAddress()+":"+strconv.Itoa(config.GetClusterPort()))

	if this.TSErr != nil {
		log.Printf("[NATUDP] Cannot resolve UDP address : %s", this.TSErr.Error())
		return
	}

	_, this.TSErr = this.TSConn.WriteToUDP([]byte(tools.RandSeq(16)), this.TSRAddr)

	if this.TSErr == nil {
		log.Printf("[NATUDP] UDP ready from %s to %s", this.TSLAddr, this.TSRAddr)
	} else {
		log.Printf("[NATUDP] UDP BLOCKED from %s to %s: %s", this.TSLAddr, this.TSRAddr, this.TSErr.Error())
		return
	}

}

func (this *TribeServer) RefreshPunchHole() {

	go this.Udp_Server()

	log.Printf("[NATUDP] Starting UDP HolePunch Engine")

	for {

		if this.TSRun == true {
			this.OpenNatUDPport()
		}
		time.Sleep(2 * time.Minute)
		log.Printf("[NATUDP] Refreshing the Hole Punch...")

	}

}

// sends a JSON []byte packet to someone. The Tribe Payload is used to know the address
// the Tribeserver is needed to know the UDP server port and IP

func (this *TribeServer) Shoot_JSON(recv TribePayload, bullet []byte) {

	_, err := this.TSConn.WriteToUDP(bullet, recv.TPsender)

	if err == nil {
		log.Printf("[JSON-UDP] UDP ready from %s to %s", this.TSLAddr, recv.TPsender)
	} else {
		log.Printf("[JSON-UDP] UDP BLOCKED from %s to %s: %s", this.TSLAddr, recv.TPsender, err.Error())
	}

}
