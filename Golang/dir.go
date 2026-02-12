package main

//import the correct functions
import (
	"fmt"
	"os"
)

//func main() {
//	fmt.Println("Hello Dir Command!")
//}

func main() {
	//If user didnt provide directory use documents folder
	var directory string
	if len(os.Args) < 2 {
		homeDir, _ := os.UserHomeDir()
		directory = homeDir + "/Documents"
	} else {
		directory = os.Args[1]
	}

	//Read everything in the directory
	items, _ := os.ReadDir(directory)

	//Print them so we see what we get
	for _, item := range items {
		fmt.Println(item.Name())
	}
}

//You can use underscore _ to ignore something without throwing an error
//"I know this exists but I dont want to use it"
