package upnp

import (
	"github.com/prestonTao/upnp"
	"log"
)

func AllUpnpOpen() {

	mapping_nntp := new(upnp.Upnp)
	if err := mapping_nntp.AddPortMapping(11119, 11119, "TCP"); err == nil {
		log.Printf("[INFO] %s", "UPnP redirect 11119 successful")
	} else {
		log.Printf("[WARNING] %s", "UPnP redirect failed for port 11119: router is UPnP-agnostic")
	}

	mapping_dht := new(upnp.Upnp)
	if err := mapping_dht.AddPortMapping(20000, 20000, "TCP"); err == nil {
		log.Printf("[INFO] %s", "UPnP redirect 20000 successful")
	} else {
		log.Printf("[WARNING] %s", "UPnP redirect failed for port 20000: router is UPnP-agnostic")
	}

}
