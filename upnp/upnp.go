package upnp

import (
	"../legione/"
	"github.com/prestonTao/upnp"
	"log"
)

func AllUpnpOpen() {

	ClusterPort := legion.GetClusterPort()

	mapping_dht := new(upnp.Upnp)
	log.Printf("[INFO] UPnP on TCP %d...", ClusterPort)
	if err := mapping_dht.AddPortMapping(ClusterPort, ClusterPort, "TCP"); err == nil {
		log.Printf("[INFO] UPnP redirect %d successful", ClusterPort)
	} else {
		log.Printf("[WARNING] No UPnP on port %d: network UPnP-agnostic", ClusterPort)
	}

	log.Printf("[INFO] UPnP on UDP %d...", ClusterPort)

	if err := mapping_dht.AddPortMapping(ClusterPort, ClusterPort, "UDP"); err == nil {
		log.Printf("[INFO] UPnP redirect %d successful", ClusterPort)
	} else {
		log.Printf("[WARNING] No UPnP on port %d: network UPnP-agnostic", ClusterPort)
	}

}
