package tribe

// this is going to contain all the BE functionalities

import (
	"../backend/"
	"../tools"
	"bufio"
	"encoding/base64" // we will need this also
	"encoding/json"   // commented to avoid compiler error in coding phase
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Empty Tribes_Execute_POST(mybuffer)

type TribesJsonPost struct {
	Command   string // a Command field is mandatory for any communication
	MessageID string
	Group     string
	Headers   string
	Body      string
	Xover     string
}

var (
	user_home       string
	messages_folder string
)

func init() {
	var user_home = tools.GetHomeDir()
	var messages_folder string = "/News/messages/"
	messages_folder = filepath.Join(user_home, messages_folder)
	os.MkdirAll(filepath.Join(user_home, "News", "messages"), 0755) // overkill. Just to be sure it exists.
}

func Tribes_BE_POST(mybuffer []byte) error {

	var mypost TribesJsonPost

	err := json.Unmarshal(mybuffer, &mypost)

	if err == nil {
		log.Println("[UDP-POST] Received a: %s", mypost.Command)
	} else {
		log.Println("[UDP-POST] Wrong post format: %s", err.Error())
		return err
	}

	// now checking if the post exists already

	if _, err := filepath.Glob(messages_folder + "/*" + mypost.MessageID + "*"); err == nil {
		log.Printf("[UDP-POST] We have %s already, discarding", mypost.MessageID)
		return nil
	}

	// converting from pure-strings
	// in the future I will add encryption here

	body_hex, _ := base64.StdEncoding.DecodeString(mypost.Body)
	header_hex, _ := base64.StdEncoding.DecodeString(mypost.Headers)
	xover_hex, _ := base64.StdEncoding.DecodeString(mypost.Xover)
	mypost.Body = string(body_hex)
	mypost.Headers = string(header_hex)
	mypost.Xover = string(xover_hex)

	// creating the filenames according with our convention
	// first the message number
	num_message, _ := strconv.Atoi(backend.GetLastNumByGroup(mypost.Group))
	num_message++
	msgnum_str := fmt.Sprintf("%05d", num_message)
	//
	// create the complete messageID
	const layout = "0601021504"
	orario := time.Now()
	id_message := mypost.MessageID + "@" + orario.Format(layout)
	//
	// then generating file names
	header_file := filepath.Join(messages_folder, "h-"+mypost.Group+"-"+msgnum_str+"-"+id_message)
	body_file := filepath.Join(messages_folder, "b-"+mypost.Group+"-"+msgnum_str+"-"+id_message)
	xover_file := filepath.Join(messages_folder, "x-"+mypost.Group+"-"+msgnum_str+"-"+id_message)
	//
	// now simply shooting strings to file
	//
	// this is for headers
	err = ShootStringToFile(mypost.Headers, header_file)
	if err != nil {
		log.Printf("[UDP-POST] Problem saving %s for %s", header_file, mypost.MessageID)
		return err
	}
	// this is for body
	err = ShootStringToFile(mypost.Body, body_file)
	if err != nil {
		log.Printf("[UDP-POST] Problem saving %s for %s", body_file, mypost.MessageID)
		return err
	}
	// this is for xover
	err = ShootStringToFile(mypost.Xover, xover_file)
	if err != nil {
		log.Printf("[UDP-POST] Problem saving %s for %s", xover_file, mypost.MessageID)
		return err
	}

	// done, so no errors

	return nil

}

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
