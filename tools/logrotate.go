package tools

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

type tribeslogfile struct {
	filename string
	logfile  *os.File
}

func init() {

	// just the first time
	var user_home = GetHomeDir()
	os.MkdirAll(filepath.Join(user_home, "News", "logs"), 0755)
	//

	var mylogfile tribeslogfile
	mylogfile.SetLogFolder()
	go mylogfile.RotateLogFolder()

}

// rotates the log folder

func (this *tribeslogfile) RotateLogFolder() {

	for {

		time.Sleep(1 * time.Hour)
		if this.logfile != nil {
			err := this.logfile.Close()
			log.Println("[LOG] close logfile returned: ", err)
		}

		this.SetLogFolder()

	}

}

// sets the log folder

func (this *tribeslogfile) SetLogFolder() {

	const layout = "2006-Jan-02.15"

	orario := time.Now()

	var user_home = GetHomeDir()
	this.filename = filepath.Join(user_home, "News", "logs", "tribes."+orario.Format(layout)+"00.log")
	log.Println("[LOG] Logfile is: " + this.filename)

	this.logfile, _ = os.Create(this.filename)

	log.SetPrefix("TRIBES> ")
	log.SetOutput(this.logfile)

}

func Log_Engine_Start() {

	log.Println("[LOG] LogRotation engine started")

}
