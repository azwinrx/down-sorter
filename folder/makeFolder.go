package folder

import (
	"log"
	"os"
)

// make folder
func MakeFolder(directoryName string) {
	err := os.Mkdir(directoryName, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Folder created:", directoryName)
}
