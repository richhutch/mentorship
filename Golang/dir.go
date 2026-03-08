// Dir go-command written by Richard Hutchinson

package main

//import the correct functions
import (
	"fmt"
	"os"
)

func main() {
	//If user didnt provide directory use documents folder
	var directory string
	if len(os.Args) < 2 {
		homeDir, _ := os.UserHomeDir()
		directory = homeDir + "/Documents"
	} else {
		directory = os.Args[1]
	}

	// Read everything in the directory
	items, err := os.ReadDir(directory)
	if err != nil {
		// Do something with the error
		fmt.Printf("Failed to read directory %q: %s\n", directory, err.Error()) // Print a helpful error message
		os.Exit(1)                                                              //Exit the program with an exit code of 1, indicating failure
	}

	// Check if -l flag was provided
	foundFlag := false
	for _, arg := range os.Args {
		if arg == "-l" {
			foundFlag = true
			break
		}
	}

	// Use foundFlag to decide what to print
	if foundFlag {
		for _, item := range items {
			info, err := item.Info()
			if err != nil {
				fmt.Printf("Error getting info for %s: %s\n", item.Name(), err.Error())
				continue
			}

			size := info.Size()
			modTime := info.ModTime().Format("Jan 02 15:04")
			fmt.Println(item.Name(), size, modTime)
		}
	} else {
		for _, item := range items {
			fmt.Println(item.Name())
		}
	}
}

//You can use underscore _ to ignore something without throwing an error
//"I know this exists but I dont want to use it"
