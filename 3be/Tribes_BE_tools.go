package tribe

// this is going to contain all the BE functionalities

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// just a quick way for saving string in new files
// please notice it will destroy the file , if existing
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

// appends one line to tihe given file.
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

// returns a file into a single string
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
