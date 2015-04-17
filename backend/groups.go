package backend

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetNumFilesByGroup(groupname string) int {
	files, _ := ioutil.ReadDir(backend.messages_folder)
	count := 0

	for fileline := range files {
		if strings.Contains(files[fileline], groupname) {

			count++
		}
	}
	return count
}

func GetFirstNumByGroup(groupname string) string {
	files, _ := ioutil.ReadDir(backend.messages_folder)

	current_num := "null"

	for filename := range files {

		if strings.Contains(files[filename], groupname) {

			pieces = strings.Split(filename, "-")
			if current_num > pieces[2] {
				current_num = pieces[2]
			}
		}
	}
	return current_num
}

func GetLastNumByGroup(groupname string) string {
	files, _ := ioutil.ReadDir(backend.messages_folder)

	current_num := ""

	for filename := range files {

		if strings.Contains(filename, groupname) {

			pieces = strings.Split(filename, "-")
			if current_num < pieces[2] {
				current_num = pieces[2]
			}
		}
	}
	return current_num
}
