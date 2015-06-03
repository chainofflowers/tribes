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

type TribesJsonGroups struct {
	Command string // a Command field is mandatory for any communication
	Groups  string // base64 encoded list of groups, one per line
	Fill    string // 32 random char to reach the AES block
}

var (
	groups_folder      string
	groups_active_file string //actually something equivalent to dht cache
)

func init() {
	user_home = tools.GetHomeDir()
	groups_folder = "/News/groups/"
	groups_folder = filepath.Join(user_home, groups_folder)
	os.MkdirAll(groups_folder, 0755) // overkill. Just to be sure it exists.
	groups_active_file = filepath.Join(user_home, groups_folder, "ng.active")

}

func Tribes_BE_Groups(mybuffer []byte) error {

	var mypost TribesJsonGroups

	err := json.Unmarshal(mybuffer, &mypost)

	if err == nil {
		log.Println("[DHT-GRP] Received a: %s", mypost.Command)
	} else {
		log.Println("[DHT-GRP] Wrong post format: %s", err.Error())
		return err
	}

	// base64 encoded so decode
	mypost_Groups, _ := base64.StdEncoding.DecodeString(mypost.Groups)

	// then split in single lines
	mygroups := SplitStringInLines(string(mypost_Groups))

	for groupID := range mygroups {

		err = AddLineToFile(mygroups[groupID], groups_active_file)
		if err == nil {
			log.Println("[DHT-GRP] Written the groups to: %s", groups_active_file)
		} else {
			log.Println("[DHT-GRP] Can't write groups to %s: %s", groups_active_file, err.Error())
			return err
		}

	}

	return nil

}
