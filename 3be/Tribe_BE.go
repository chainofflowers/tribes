package tribe

// this is going to contain all the BE functionalities

import (
	// "encoding/json" // commented to avoid compiler error in coding phase
	// "encoding/base64" // we will need this also
	"log"
)

// Empty Tribes_Execute_POST(mybuffer)

func Tribes_BE_POST(mybuffer []byte) error {

	// it will unmarshal the JSON and write it to files.

	log.Println("[UDP-POST] Received a POST")
	return nil

}
