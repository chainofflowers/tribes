package backend

import (
	"log"
	"path/filepath"
	"sort"
)

func GetNumFilesByGroup(groupname string) int {

	if files, err := filepath.Glob(backend.messages_folder + "*" + groupname + "*"); err != nil {
		log.Printf("[SOB] No messages for group %s ", groupname)
		return 0
	} else {
		log.Printf("[WOW] %d messages for group %s ", len[files], groupname)
		return len(files)

	}

}

func GetFirstNumByGroup(groupname string) string {

	if files, err := filepath.Glob(backend.messages_folder + "*" + groupname + "*"); err != nil {
		log.Printf("[SOB] No messages for group %s ", groupname)
		return "0"
	} else {
		sort.Strings(files)
		pieces = strings.Split(files[0], "-")
		log.Printf("[WOW] %s first message for %s ", pieces[2], groupname)
		return pieces[2]
	}

}

func GetLastNumByGroup(groupname string) string {
	if files, err := filepath.Glob(backend.messages_folder + "*" + groupname + "*"); err != nil {
		log.Printf("[SOB] No messages for group %s ", groupname)
		return "0"
	} else {
		sort.Strings(files)
		pieces = strings.Split(files[len(files)-1], "-")
		log.Printf("[WOW] %s last message for %s ", pieces[2], groupname)
		return pieces[2]
	}

}

func ResponseToNNTPGROUP(groupname string) string {
	response = "211 " + GetNumFilesByGroup(groupname) + " " + GetFirstNumByGroup(groupname) + " " + GetLastNumByGroup(groupname) + " group is now " + groupname
	log.Printf("[OK] answering back %s ", response)
	return response
}
