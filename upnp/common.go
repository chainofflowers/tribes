package upnp

import (
	// "log"
	"log"
	"net"
	"strings"
)

func GetLocalIntenetIp() string {

	conn, err := net.Dial("udp", "example.com:80")
	if err != nil {
		log.Printf("[UPnP] GetLocalInternetIP: cannot use UDP")
		return "0.0.0.0"
	}
	defer conn.Close()
	torn := strings.Split(conn.LocalAddr().String(), ":")
	return torn[0]
}

// This returns the list of local ip addresses which other hosts can connect
// to (NOTE: Loopback ip is ignored).
func GetLocalIPs() ([]*net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	ips := make([]*net.IP, 0)
	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		if ipnet.IP.IsLoopback() {
			continue
		}

		ips = append(ips, &ipnet.IP)
	}

	return ips, nil
}
