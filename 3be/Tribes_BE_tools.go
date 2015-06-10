package tribe

// this is going to contain all the BE functionalities

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// ShootStringToFile : just a quick way for saving string in new files
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

// AddLineToFile : appends one line to the given file.
// only when the line doesn't exists already
func AddLineToFile(peer string, filename string) error {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	contentString := string(content)

	if strings.Contains(contentString, peer) == false {

		contentString += "\n" + peer

	}

	err = ioutil.WriteFile(filename, []byte(contentString), 0755)
	if err != nil {
		return err
	}

	return nil

}

// RetrieveStringFromFile returns a file into a single string
func RetrieveStringFromFile(filename string) string {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "EMPTY FILE"
	}

	return string(content)

}

// SplitStringInLines splits a string in lines
// regardless EOL is Windows or UNIX
func SplitStringInLines(myblock string) []string {

	// create a splitter because "split" adds an empty line after the last \n
	splitter := func(c rune) bool {
		return (c == '\n' || c == '\r') // this is both for windows and unix-like EOL
	}

	lines := strings.FieldsFunc(myblock, splitter)

	return lines

}

// CreateSerialByGroup : duplicate, creates a serial ID for a given group
// to move in tools after the Tag #1
func CreateSerialByGroup(groupname string) string {
	if files, err := filepath.Glob(messages_folder + "/h-" + groupname + "-*"); err != nil {
		log.Printf("[SOB] No messages for group %s ", groupname)
		return "0"
	} else {
		if files == nil {
			files = append(files, "bh-ng-0-sh1")
		}
		sort.Strings(files)
		pieces := strings.Split(files[len(files)-1], "-")
		log.Printf("[WOW] %s last message for %s ", pieces[2], groupname)
		return pieces[2]
	}

}
