package averno

import "fmt"


func init() {
    if (os.Getuid() == 0) || ((os.Getgid() == 0) {
    fmt.Println("AAAARGH! ROOT! ROOT! ROOOOOT! ")
    fmt.Println("This is not a tree! We need no root!")
    os.Exit(1)
                        }
}


func main() {
	
	
	
}
