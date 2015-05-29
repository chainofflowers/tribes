package tribe

// this is going to contain all the BE functionalities

import (
	"encoding/json" // commented to avoid compiler error in coding phase
	"fmt"
	"log"
)

type TribesJsonRegister struct {
	Command string // a Command field is mandatory for any communication
	Proof   string // Fibo 01123581321345589144233377610987 encrypted with TribeID
}

func Tribes_BE_REG(mypayload TribePayload) error {

	var mypost TribesJsonRegister

	err := json.Unmarshal(mypayload.TPbuffer[0:mypayload.TPsize], &mypost)

	if err == nil {
		log.Println("[UDP-REG] Received a: %s", mypost.Command)
	} else {
		log.Println("[UDP-REG] Wrong post format: %s", err.Error())
		return err
	}

	// Decrypt the Proof

	if ProofIsOk(mypost.Proof) == false {
		err := fmt.Errorf("Not our tribe")
		return err
	}

	// Write the peer in the active peers file

	err = AddLineToFile(mypayload.TPsender.String(), peers_active_file)
	if err != nil {
		return err
	}

	return nil

}
