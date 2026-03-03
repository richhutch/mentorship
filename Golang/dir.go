// Dir go-command written by Richard Hutchinson

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
	items, err := os.ReadDir(directory)
	if err != nil {
		// Do something with the error
		fmt.Printf("Failed to read directory %q: %s\n", directory, err.Error) // Print a helpful error message to the user
		os.Exit(1) // Exit the program with an exit code of 1, indicating failure.
	}

	//Print them so we see what we get
	for _, item := range items {
		fmt.Println(item.Name())
	}
}

//You can use underscore _ to ignore something without throwing an error
//"I know this exists but I dont want to use it"
