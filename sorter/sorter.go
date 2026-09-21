package sorter

import (
	"log"
	"os"
	"path/filepath"
)

func SorterFile() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue 
		}
		log.Println(file.Name())
		
	}
}
