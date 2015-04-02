package averno

import (
	"fmt"
	"os"
)

// just to avoid running it as root

func init() {
	
		if (os.Getuid() == 0) || ((os.Getgid() == 0) {
    		fmt.Println("AAAARGH! ROOT! ROOT! ROOOOOT! ")
    		fmt.Println("This is not a tree! We need no roots!")
    		os.Exit(1)
                        }
	
}


func main() {
	
	os.exit(0)	
	
}
