package punchhole

import (
	"../config/"
	"../tools/"
	"log"
	"net"
	"time"
)

type MyPunchHole struct {
	MyLocalAddr  net.UDPAddr
	MyRemoteAddr net.UDPAddr
	MyUdpConn    *net.UDPConn
	MyError      error
}

func (this *MyPunchHole) OpenNatUDPport() {

	this.MyLocalAddr.IP = net.ParseIP(tools.ReadIpFromHost())
	this.MyLocalAddr.Port = config.GetClusterPort()
	this.MyRemoteAddr.IP = net.ParseIP(tools.RandomIPAddress())
	this.MyRemoteAddr.Port = config.GetClusterPort()

	this.MyUdpConn, this.MyError = net.DialUDP("udp", &this.MyLocalAddr, &this.MyRemoteAddr)

	if this.MyError != nil {
		log.Printf("[NATUDP] Cannot open UDP from %s:%d to %s:%d", this.MyLocalAddr.IP, this.MyLocalAddr.Port, this.MyRemoteAddr.IP, this.MyRemoteAddr.Port)
		log.Printf("[NATUDP] Error: " + this.MyError.Error())
		return
	}

	_, this.MyError = this.MyUdpConn.Write([]byte(tools.RandSeq(16)))

	if this.MyError == nil {
		log.Printf("[NATUDP] UDP ready from %s:%d to %s:%d", this.MyLocalAddr.IP, this.MyLocalAddr.Port, this.MyRemoteAddr.IP, this.MyRemoteAddr.Port)
	} else {
		log.Printf("[NATUDP] UDP BLOCKED from %s:%d to %s:%d", this.MyLocalAddr.IP, this.MyLocalAddr.Port, this.MyRemoteAddr.IP, this.MyRemoteAddr.Port)
	}

	this.MyUdpConn.Close()

}

func (this *MyPunchHole) RefreshPunchHole() {

	log.Printf("[NATUDP] Starting UDP HolePunch Engine")

	for {

		this.OpenNatUDPport()
		time.Sleep(2 * time.Minute)
		log.Printf("[NATUDP] Refreshing the Hole Punch...")

	}

}
