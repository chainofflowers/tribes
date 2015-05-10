package upnp

import (
	"../config/"
	"log"
	"strings"
	"time"
)

func AllUpnpOpen() {

	upnp_renew := time.NewTicker(5 * time.Minute)
	mapping_http := new(Upnp)

	for {

		log.Printf("[UPnP] Renew the UPnP lease")

		ClusterPort := config.GetClusterPort()

		err := mapping_http.SearchGateway()
		if err != nil {
			log.Printf("[UPnP] Problem getting the gateway:  %s...", err.Error())
		} else {
			log.Printf("[UPnP] Local ip address: %s", mapping_http.LocalHost)
			torn := strings.Split(mapping_http.Gateway.Host, ":")
			log.Printf("[UPnP] Gateway ip address is %s on port: %s", torn[0], torn[1])
		}

		err = mapping_http.ExternalIPAddr()
		if err != nil {
			log.Printf("[UPnP] Problem getting my external IP:  %s...", err.Error())

		} else {
			log.Printf("[UPnP] WAN ip address: %s", mapping_http.GatewayOutsideIP)

		}

		log.Printf("[UPnP] UPnP on TCP %d...", ClusterPort)
		if err = mapping_http.AddPortMapping(ClusterPort, ClusterPort, "TCP"); err == nil {
			log.Printf("[UpNp] UPnP redirect TCP %d : no errors from %s", ClusterPort, mapping_http.Gateway.Host)
		} else {
			log.Printf("[UPnP] No UPnP on TCP %d:  %s", ClusterPort, err.Error())
		}

		log.Printf("[UPnP] UPnP on UDP %d...", ClusterPort)

		if err = mapping_http.AddPortMapping(ClusterPort, ClusterPort, "UDP"); err == nil {
			log.Printf("[UPnP] UPnP redirect UDP %d : no errors from %s", ClusterPort, mapping_http.Gateway.Host)
		} else {
			log.Printf("[UPnP] No UPnP on UDP %d:  %s", ClusterPort, err.Error())
		}

		<-upnp_renew.C

		mapping_http.

	}

}
