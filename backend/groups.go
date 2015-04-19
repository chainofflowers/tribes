package backend

import (
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func GetNumFilesByGroup(groupname string) string {

	if files, err := filepath.Glob(messages_folder + "/h-*" + groupname + "*"); err != nil {
		log.Printf("[SOB] No messages for group %s, %s ")
		return "0"
	} else {
		msg_num := len(files)
		resp := strconv.Itoa(msg_num)
		log.Printf("[WOW] %s messages for group %s : %s ", resp, groupname, files)
		return resp

	}

}

func GetFirstNumByGroup(groupname string) string {

	if files, err := filepath.Glob(messages_folder + "/h-*" + groupname + "*"); err != nil {
		log.Printf("[SOB] No first message for group %s ", groupname)
		return "0"
	} else {

		if files == nil {
			files = append(files, "bh-ng-0-sh1")
		}

		sort.Strings(files)
		pieces := strings.Split(files[0], "-")
		log.Printf("[WOW] %s first message for %s ", pieces[2], groupname)
		return pieces[2]
	}

}

func GetLastNumByGroup(groupname string) string {
	if files, err := filepath.Glob(messages_folder + "/h-*" + groupname + "*"); err != nil {
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

func ResponseToNNTPGROUP(groupname string) string {

	if strings.Contains(groupname, "-") {
		response := "411 no such news group\n"
		log.Printf("[ERR] invalid group name %s ", groupname)
        return response
	} else {

		response := "211 " + GetNumFilesByGroup(groupname) + " " + GetFirstNumByGroup(groupname) + " " + GetLastNumByGroup(groupname) + " group is now " + groupname + " \n"
        log.Printf("[OK] answering back %s ", response)
        return response
	}

}
