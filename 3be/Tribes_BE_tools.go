package tribe

// this is going to contain all the BE functionalities

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// just a functional for saving file quickly
// to move in tools, maybe

func ShootStringToFile(mystring string, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	fmt.Fprint(w, mystring)

	return w.Flush()

}

// functional for retrieving peers from the file, adding one and saving back

func AddPeerToFile(peer string, filename string) error {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	content_string := string(content)

	if strings.Contains(content_string, peer) == false {

		content_string += "\n" + peer

	}

	err = ioutil.WriteFile(filename, []byte(content_string), 0755)
	if err != nil {
		return err
	}

	return nil

}

func RetrieveStringFromFile(filename string) string {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}

	return string(content)

}
