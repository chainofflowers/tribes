package tribe

// this is going to contain all the BE functionalities

import (
	"../tools"
	"encoding/json" // commented to avoid compiler error in coding phase
	"log"
	"os"
	"path/filepath"
	"strings"
)

type TribesJsonPeers struct {
	Command string // a Command field is mandatory for any communication
	Peers   string
}

var (
	peers_folder       string
	peers_active_file  string //actually something equivalent to dht cache
	peers_initial_file string
)

func init() {
	user_home = tools.GetHomeDir()
	peers_folder = "/News/peers/"
	peers_folder = filepath.Join(user_home, peers_folder)
	os.MkdirAll(peers_folder, 0755) // overkill. Just to be sure it exists.
	peers_active_file = filepath.Join(user_home, peers_folder, "peers.active")
	peers_initial_file = filepath.Join(user_home, peers_folder, "peers.initial")

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

	// create a splitter because "split" adds an empty line after the last \n
	splitter := func(c rune) bool {
		return (c == '\n' || c == '\r') // this is for windows and unix like EOL
	}

	mypeers := strings.FieldsFunc(mypost.Peers, splitter)

	for peername := range mypeers {

		err = AddPeerToFile(mypeers[peername], peers_active_file)
		if err == nil {
			log.Println("[UDP-PEER] Written the peers to: %s", peers_active_file)
		} else {
			log.Println("[UDP-PEER] Can't write peers to %s: %s", peers_active_file, err.Error())
			return err
		}

	}

	return nil

}
