package tribe

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	dht "tribes/cripta"
	to "tribes/tools"
)

func DhtReceiveBody(dhtPayload string) {

	var MyHeaders map[string]string
	MyHeaders = make(map[string]string)

	MyHeaders = dht.GpgGetHeaders(dhtPayload)

	messageId := MyHeaders[TRIBES_H_MID]
	groupname := MyHeaders[TRIBES_H_GID]

	dhtContent := dht.GpgDecrypt(dhtPayload)

	num_message, _ := strconv.Atoi(CreateSerialByGroup(groupname))
	num_message++

	msgnum_str := fmt.Sprintf("%05d", num_message)

	body_file := filepath.Join(to.MessagesFolder, "b-"+groupname+"-"+msgnum_str+"-"+messageId)

	if to.TheFileExists(body_file) == false {
		ShootStringToFile(dhtContent, body_file)
	} else {
		log.Printf("[DHT-3BE] we have %s already: doing nothing on body", messageId)
	}

}

func DhtReceiveHeaders(dhtPayload string) {

	var user_home = to.GetHomeDir()
	var messages_folder string = "/News/messages/"
	messages_folder = filepath.Join(user_home, messages_folder)

	var MyHeaders map[string]string
	MyHeaders = make(map[string]string)

	MyHeaders = dht.GpgGetHeaders(dhtPayload)

	messageId := MyHeaders[TRIBES_H_MID]
	groupname := MyHeaders[TRIBES_H_GID]

	dhtContent := dht.GpgDecrypt(dhtPayload)

	num_message, _ := strconv.Atoi(CreateSerialByGroup(groupname))
	num_message++

	msgnum_str := fmt.Sprintf("%05d", num_message)

	body_file := filepath.Join(messages_folder, "h-"+groupname+"-"+msgnum_str+"-"+messageId)

	if to.TheFileExists(body_file) == false {
		ShootStringToFile(dhtContent, body_file)
	} else {
		log.Printf("[DHT-3BE] we have %s already: doing nothing on headers", messageId)
	}

}

func DhtReceiveXover(dhtPayload string) {

	var user_home = to.GetHomeDir()
	var messages_folder string = "/News/messages/"
	messages_folder = filepath.Join(user_home, messages_folder)

	var MyHeaders map[string]string
	MyHeaders = make(map[string]string)

	MyHeaders = dht.GpgGetHeaders(dhtPayload)

	messageId := MyHeaders[TRIBES_H_MID]
	groupname := MyHeaders[TRIBES_H_GID]

	dhtContent := dht.GpgDecrypt(dhtPayload)

	num_message, _ := strconv.Atoi(CreateSerialByGroup(groupname))
	num_message++

	msgnum_str := fmt.Sprintf("%05d", num_message)

	body_file := filepath.Join(messages_folder, "x-"+groupname+"-"+msgnum_str+"-"+messageId)

	if to.TheFileExists(body_file) == false {
		ShootStringToFile(dhtContent, body_file)
	} else {
		log.Printf("[DHT-3BE] we have %s already: doing nothing on xover", messageId)
	}

}

func DhtReceiveGroup(dhtPayload string) {

	var user_home = to.GetHomeDir()
	var active_ng_file string = "/News/groups/ng.active"
	var all_ng_file string = "/News/groups/ng.all"
	active_ng_file = filepath.Join(user_home, active_ng_file)
	all_ng_file = filepath.Join(user_home, all_ng_file)

	dhtContent := dht.GpgDecrypt(dhtPayload)

	AddLineToFile(dhtContent, active_ng_file)
	AddLineToFile(dhtContent, all_ng_file)

}
