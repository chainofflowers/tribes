package punchhole

import (
	"../config/"
	"../tools/"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type MyPunchHole struct {
	MyLocalAddr  net.UDPAddr
	MyRemoteAddr net.UDPAddr
	MyUdpConn    *net.UDPConn
	MyError      error
}

func (h *MyPunchHole) OpenNatUDPport() {

	h.MyLocalAddr.IP = net.ParseIP(tools.ReadIpFromHost())
	h.MyLocalAddr.Port = config.GetClusterPort()
	h.MyRemoteAddr.IP = net.ParseIP(RandomIPAddress())
	h.MyRemoteAddr.Port = config.GetClusterPort()

	h.MyUdpConn, h.MyError = net.DialUDP("udp", &h.MyLocalAddr, &h.MyRemoteAddr)

	if h.MyError != nil {
		log.Printf("[NATUDP] Cannot open UDP from %s:%d to %s:%d", h.MyLocalAddr.IP, h.MyLocalAddr.Port, h.MyRemoteAddr.IP, h.MyRemoteAddr.Port)
		return
	}

	log.Printf("[NATUDP] UDP ready from %s:%d to %s:%d", h.MyLocalAddr.IP, h.MyLocalAddr.Port, h.MyRemoteAddr.IP, h.MyRemoteAddr.Port)

	h.MyUdpConn.Close()
	h.MyUdpConn, h.MyError = net.ListenUDP("udp", &h.MyLocalAddr)

	if h.MyError != nil {
		log.Printf("[NATUDP] UDP cannot listen on %s:%d", h.MyLocalAddr.IP, h.MyLocalAddr.Port)
	} else {
		log.Printf("[NATUDP] UDP listening on %s:%d", h.MyLocalAddr.IP, h.MyLocalAddr.Port)

	}
}

func (h *MyPunchHole) RefreshPunchHole() {

	log.Printf("[NATUDP] Starting UDP HolePunch Engine")

	for {

		h.OpenNatUDPport()
		time.Sleep(2 * time.Minute)
		log.Printf("[NATUDP] Closing the old port")
		h.MyUdpConn.Close()
		log.Printf("[NATUDP] Refreshing the Hole Punch")

	}

}

func RandomIPAddress() string {

	var IPfields string
	rand.Seed(time.Now().Unix())
	IPfields += strconv.Itoa(rand.Intn(254)) + "."
	IPfields += strconv.Itoa(rand.Intn(254)) + "."
	IPfields += strconv.Itoa(rand.Intn(254)) + "."
	IPfields += strconv.Itoa(rand.Intn(254))

	return IPfields

}
