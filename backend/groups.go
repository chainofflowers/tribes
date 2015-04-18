package backend

import (
	"log"
	"path/filepath"
	"sort"
	"strconv"
    "strings"
)

func GetNumFilesByGroup(groupname string) string {

	if files, err := filepath.Glob(messages_folder + "*" + groupname + "*"); err != nil {
		log.Printf("[SOB] No messages for group %s ", groupname)
		return "0"
	} else {
        msg_num := len(files)/2    // 'cause we save two files for each message
		log.Printf("[WOW] %d messages for group %s ", msg_num, groupname)
		return strconv.Itoa(msg_num)

	}

}

func GetFirstNumByGroup(groupname string) string {

	if files, err := filepath.Glob(messages_folder + "*" + groupname + "*"); err != nil {
		log.Printf("[SOB] No messages for group %s ", groupname)
		return "0"
	} else {
		sort.Strings(files)
		pieces := strings.Split(files[0], "-")
		log.Printf("[WOW] %s first message for %s ", pieces[2], groupname)
		return pieces[2]
	}

}

func GetLastNumByGroup(groupname string) string {
	if files, err := filepath.Glob(messages_folder + "*" + groupname + "*"); err != nil {
		log.Printf("[SOB] No messages for group %s ", groupname)
		return "0"
	} else {
		sort.Strings(files)
		pieces := strings.Split(files[len(files)-1], "-")
		log.Printf("[WOW] %s last message for %s ", pieces[2], groupname)
		return pieces[2]
	}

}

func ResponseToNNTPGROUP(groupname string) string {
	response := "211 " + GetNumFilesByGroup(groupname) + " " + GetFirstNumByGroup(groupname) + " " + GetLastNumByGroup(groupname) + " group is now " + groupname
	log.Printf("[OK] answering back %s ", response)
	return response
}
