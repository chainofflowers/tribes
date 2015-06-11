package tools

import (
	"log"
	"os"
	"path/filepath"
)

// here I will group everything about filenames.
// this package will only be imported.
// each and every package willing to use this values shall import "tribes/filesys"

var (
	ActiveNgFile   string = "/News/groups/ng.active"
	NewNgFile      string = "/News/groups/ng.local"
	AllNgFile      string = "/News/groups/ng.all"
	MessagesFolder string = "/News/messages/"
	UserHome       string
	TribesHome     string
	ConfigPath     string
	ConfigFile     string
	LogFolder      string = "/News/logs"
	GroupsFolder   string = "/News/groups"
)

// here we init and create all the folders and the variables we need. This operation
// SHALL NOT be done again

func init() {

	UserHome = GetHomeDir()
	TribesHome = Hpwd()
	ActiveNgFile = filepath.Join(UserHome, ActiveNgFile)
	NewNgFile = filepath.Join(UserHome, NewNgFile)
	AllNgFile = filepath.Join(UserHome, AllNgFile)
	MessagesFolder = filepath.Join(UserHome, MessagesFolder)
	LogFolder = filepath.Join(UserHome, LogFolder)
	GroupsFolder = filepath.Join(UserHome, GroupsFolder)
	ConfigFile = filepath.Join(UserHome, "News", "config.toml")
	ConfigPath = filepath.Join(UserHome, "News")

	// just in case, making sure the folder are existing
	os.MkdirAll(MessagesFolder, 0755)
	os.MkdirAll(GroupsFolder, 0755)
	os.MkdirAll(LogFolder, 0755)
	// initializing some files

	if TheFileExists(ActiveNgFile) == false {
		log.Printf("[BE-FS] creating file %s", ActiveNgFile)
		os.Create(ActiveNgFile)
	}

	if TheFileExists(NewNgFile) == false {
		log.Printf("[BE-FS] creating file %s", NewNgFile)
		os.Create(NewNgFile)
	}

	if TheFileExists(AllNgFile) == false {
		log.Printf("[BE-FS] creating file %s", AllNgFile)
		os.Create(AllNgFile)
	}

}
