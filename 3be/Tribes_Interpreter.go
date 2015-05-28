package tribe

import (
	"encoding/json"
	"log"
)

// assuming whatever kind of message we receive, it has a "COMMAND" field in JSON
// this field will be encrypted in the future.

type BareCommand struct {
	// this is a structure only needed to check the command.
	Command string
	// this is where I expect to see a command
	Whatever string
}

func GetJSONCommand(mybuffer []byte) string {

	var JSON_command BareCommand

	err := json.Unmarshal(mybuffer, &JSON_command)
	if err != nil {
		log.Println("[UDP-JSON] Cannot marshal the Payload: %s", err.Error())
		return "NOOP"
	} else {
		return JSON_command.Command
	}
}

// now let's go with the interpreter
// it needs to know about the connection to write to, and to know about who sent the payload

func (this *TribeServer) Tribes_Interpreter(mypayload TribePayload) {

	mycommand := GetJSONCommand(mypayload.TPbuffer)

	switch mycommand {

	case "NOOP":
		break // doing nothing
		//
		// Implementation of single post exchange
	case "HEREPOST":
		// herepost just returns the requested post
		err := Tribes_BE_POST(mypayload.TPbuffer[0:mypayload.TPsize])
		if err != nil {
			log.Println("[UDP-INT] Cannot execute POST %s", err.Error())
		}
		// each function should have the full buffer when starting
		// the ones with BE are saving something.
		// the ones with FE are answeing back (so they need to know who to answer
		// all FE functions will return a []byte to shoot with Shoot_JSON
	case "GIMMEPOST":
		// gimmepost just requires to send a post back
		// giving the messageID as argument
		// those functions starting with GIMME are asked to reply to the peer
		//
		// Implementation of PEERS exchange
	case "HEREPEERS":
		err := Tribes_BE_PEERS(mypayload.TPbuffer[0:mypayload.TPsize])
		if err != nil {
			log.Println("[UDP-INT] Cannot execute PEERS %s", err.Error())
		}
		// herepeers gives a list of known peers
	case "GIMMEPEERS":
		// asks for a list of known peers
		//
		// Implementation of GROUPS exchange
	case "HEREGROUPS":
		// Sends the list of active groups
	case "GIMMEGROUPS":
		// Asks for the list of active groups
		//
		// Implementation of group index: to have a list of messageIDs for a group
	case "HEREINDEX":
		// Gives a list of MessageIDs on a specified group
	case "GIMMEINDEX":
		// Asks for a list of posts in a specified group
		//
		// Now registering peers, let's copy SIP
	case "REGISTER":
		// a peer asks to be registered

	// whatever else is lost
	default:
		break

	}

}
