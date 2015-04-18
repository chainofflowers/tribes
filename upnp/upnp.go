package upnp

import (
	"../config/"
	"github.com/prestonTao/upnp"
	"log"
)

var ClusterPort int

func init() {
    ClusterPort := config.GetClusterPort()
}



func AllUpnpOpen() {



	mapping_http := new(upnp.Upnp)
	log.Printf("[INFO] UPnP on TCP %d...", ClusterPort)
	if err := mapping_http.AddPortMapping(ClusterPort, ClusterPort, "TCP"); err == nil {
		log.Printf("[INFO] UPnP redirect TCP %d successful", ClusterPort)
	} else {
		log.Printf("[WARNING] No UPnP on TCP %d: network UPnP-agnostic", ClusterPort)
	}

	log.Printf("[INFO] UPnP on UDP %d...", ClusterPort)

	if err := mapping_http.AddPortMapping(ClusterPort, ClusterPort, "UDP"); err == nil {
		log.Printf("[INFO] UPnP redirect UDP %d successful", ClusterPort)
	} else {
		log.Printf("[WARNING] No UPnP on UDP %d: network UPnP-agnostic", ClusterPort)
	}

}
