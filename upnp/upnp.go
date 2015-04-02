package upnp

import (
	"github.com/prestonTao/upnp"
	"log"
)

func AllUpnpOpen() {

	mapping_dht := new(upnp.Upnp)
	if err := mapping_dht.AddPortMapping(20000, 20000, "TCP"); err == nil {
		log.Printf("[INFO] %s", "UPnP redirect 20000 successful")
	} else {
		log.Printf("[WARNING] %s", "No UPnP on port 20000: network UPnP-agnostic")
	}

}
