package tribe

// this is going to contain all the BE functionalities

import (
	"encoding/json" // commented to avoid compiler error in coding phase
	"fmt"
	"log"
	"tribes/config"
	"tribes/cripta"
)

type TribesJsonRegister struct {
	Command string // a Command field is mandatory for any communication
	Proof   string // Fibo 01123581321345589144233377610987 encrypted with TribeID
}

func Tribes_BE_REG(mypayload TribePayload) error {

	var mypost TribesJsonRegister
	const proof string = "01123581321345589144233377610987"

	err := json.Unmarshal(mypayload.TPbuffer[0:mypayload.TPsize], &mypost)

	if err == nil {
		log.Println("[UDP-PEER] Received a: %s", mypost.Command)
	} else {
		log.Println("[UDP-PEER] Wrong post format: %s", err.Error())
		return err
	}

	mykey := config.GetTribeID()
	// Decrypt the Proof

	if cripta.EasyDeCrypt(mypost.Proof, mykey) != proof {
		err := fmt.Errorf("Wrong proof, not our tribe")
		return err
	}

	err = AddPeerToFile(mypayload.TPsender.String(), peers_active_file)

	if err != nil {
		return err
	}

	return nil

}
