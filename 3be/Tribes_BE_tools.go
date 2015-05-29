package tribe

// this is going to contain all the BE functionalities

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"tribes/config"
	"tribes/cripta"
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

// functional for adding peers from the file, adding one and saving back

func AddLineToFile(peer string, filename string) error {

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

// retrieves an entire file into a string

func RetrieveStringFromFile(filename string) string {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "EMPTY FILE"
	}

	return string(content)

}

// splits a string in lines
// regardless EOL is Windows or UNIX

func SplitStringInLines(myblock string) []string {

	// create a splitter because "split" adds an empty line after the last \n
	splitter := func(c rune) bool {
		return (c == '\n' || c == '\r') // this is for windows and unix like EOL
	}

	lines := strings.FieldsFunc(myblock, splitter)

	return lines

}

// Checks if the proof is ok and the tribe is ok.

func ProofIsOk(proof string) bool {

	const layout = "06010215"
	orario := time.Now()

	// the Prood will contain pad +  a random string at beginning and end.
	var pad string = "F01123581321345589144233377610987" + orario.Format(layout)

	tmp_decrypt := cripta.EasyDeCrypt(proof, config.GetTribeID())

	return strings.Contains(tmp_decrypt, pad)

}

// generating a Proof

func GenerateProof() string {

	const layout = "06010215"
	orario := time.Now()

	var pad string = "F01123581321345589144233377610987" + orario.Format(layout)

	tmp_key := config.GetTribeID()

	return cripta.EasyCrypt(pad, tmp_key)

}
