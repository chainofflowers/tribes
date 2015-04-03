package upnp

import (
	"github.com/prestonTao/upnp"
	"log"
)

func AllUpnpOpen() {

	mapping_dht := new(upnp.Upnp)
	log.Printf("[INFO] %s", "UPnP on TCP 20000...")
	if err := mapping_dht.AddPortMapping(20000, 20000, "TCP"); err == nil {
		log.Printf("[INFO] %s", "UPnP redirect 20000 successful")
	} else {
		log.Printf("[WARNING] %s", "No UPnP on port 20000: network UPnP-agnostic")
	}

	log.Printf("[INFO] %s", "UPnP on UDP 20000...")

	if err := mapping_dht.AddPortMapping(20000, 20000, "UDP"); err == nil {
		log.Printf("[INFO] %s", "UPnP redirect 20000 successful")
	} else {
		log.Printf("[WARNING] %s", "No UPnP on port 20000: network UPnP-agnostic")
	}


}
