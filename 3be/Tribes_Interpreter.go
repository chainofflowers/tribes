package tribe

import (
	"log"
	"tribes/cripta"
)

type DhtHeaders map[string]string

// returns the field "Command" in a GPG payload
func GetGPGCommand(mybuffer string) string {

	var tmp_map DhtHeaders
	tmp_map = make(DhtHeaders)
	tmp_map = cripta.GpgGetHeaders(mybuffer)
	if val, ok := tmp_map[TRIBES_H_CMD]; ok {
		return val
	}
	return "NOOP"
}

// Receives a JSON payload and decides what to do, looking at the "Command" field.
// Please notice the payload is encrypted and zipped.
func Tribes_Interpreter(payload string) {

	mycommand := GetGPGCommand(payload)

	log.Print("[DHT-PGP] Got headers from payload")

	switch mycommand {

	case TRIBES_NOOP:
		break
		// doing nothing
		//
	case TRIBES_BODY:
		DhtReceiveBody(payload)

	case TRIBES_HEADER:
		DhtReceiveHeaders(payload)

	case TRIBES_NEWGROUP:
		DhtReceiveGroup(payload)

	case TRIBES_XOVER:
		DhtReceiveXover(payload)

	// whatever else is lost
	default:
		break

	}

}
