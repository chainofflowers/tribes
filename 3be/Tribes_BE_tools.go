package tribe

// this is going to contain all the BE functionalities

import (
	"bufio"
	"fmt"
	"os"
)

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
