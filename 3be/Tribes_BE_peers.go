package tribe

// this is going to contain all the BE functionalities

import (
	"encoding/base64"
	"encoding/json" // commented to avoid compiler error in coding phase

	"log"
	"os"
	"path/filepath"
	"tribes/tools"
)

type TribesJsonPeers struct {
	Command string // a Command field is mandatory for any communication
	Peers   string // base64 encoded list of peers, one per line
	Fill    string // 32 random chars to reach the AES blocksize
}

var (
	peers_folder      string
	peers_active_file string //actually something equivalent to dht cache
)

func init() {
	user_home = tools.GetHomeDir()
	peers_folder = "/News/peers/"
	peers_folder = filepath.Join(user_home, peers_folder)
	os.MkdirAll(peers_folder, 0755) // overkill. Just to be sure it exists.
	peers_active_file = filepath.Join(user_home, peers_folder, "peers.active")

}

func Tribes_BE_PEERS(mybuffer []byte) error {

	var mypost TribesJsonPeers

	err := json.Unmarshal(mybuffer, &mypost)

	if err == nil {
		log.Println("[UDP-PEER] Received a: %s", mypost.Command)
	} else {
		log.Println("[UDP-PEER] Wrong post format: %s", err.Error())
		return err
	}

	// base64 encoded
	mypost_Peers, _ := base64.StdEncoding.DecodeString(mypost.Peers)

	mypeers := SplitStringInLines(string(mypost_Peers))

	for peername := range mypeers {

		err = AddLineToFile(mypeers[peername], peers_active_file)
		if err == nil {
			log.Println("[UDP-PEER] Written the peers to: %s", peers_active_file)
		} else {
			log.Println("[UDP-PEER] Can't write peers to %s: %s", peers_active_file, err.Error())
			return err
		}

	}

	return nil

}
