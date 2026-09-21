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
		// IsDir was used to check is that was a file or a folder, returning boolean
		if file.IsDir() {
			continue
		}

		log.Println(file.Name())
		fileSelected := file.Name()
		extension := filepath.Ext(fileSelected)

		switch extension {
		// Document
		case ".txt":
			os.Rename(fileSelected, "Documents/txt/"+fileSelected)
		case ".docx":
			os.Rename(fileSelected, "Documents/word/"+fileSelected)
		case ".pptx":
			os.Rename(fileSelected, "Documents/ppt/"+fileSelected)
		case ".xlsx":
			os.Rename(fileSelected, "Documents/excel/"+fileSelected)
		case ".pdf":
			os.Rename(fileSelected, "Documents/pdf/"+fileSelected)

		// Audio
		case ".mp3":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".mp4":
			os.Rename(fileSelected, "Video/"+fileSelected)

		// Image
		case ".jpg":
			os.Rename(fileSelected, "Image/"+fileSelected)
		case ".png":
			os.Rename(fileSelected, "Image/"+fileSelected)
		case ".gif":
			os.Rename(fileSelected, "Image/"+fileSelected)
		case ".bmp":
			os.Rename(fileSelected, "Image/"+fileSelected)
		case ".svg":
			os.Rename(fileSelected, "Image/"+fileSelected)

		// Application
		case ".exe":
			os.Rename(fileSelected, "Application/"+fileSelected)
		case ".msi":
			os.Rename(fileSelected, "Application/"+fileSelected)
		case ".apk":
			os.Rename(fileSelected, "Application/"+fileSelected)
		case ".deb":
			os.Rename(fileSelected, "Application/"+fileSelected)
		case ".rpm":
			os.Rename(fileSelected, "Application/"+fileSelected)

		//Compressed
		case ".zip":
			os.Rename(fileSelected, "Compressed/ZIP/"+fileSelected)
		case ".rar":
			os.Rename(fileSelected, "Compressed/RAR/"+fileSelected)
		case ".7z":
			os.Rename(fileSelected, "Compressed/7z/"+fileSelected)
		case ".tar":
			os.Rename(fileSelected, "Compressed/TAR/"+fileSelected)
		case ".gz":
			os.Rename(fileSelected, "Compressed/GZ/"+fileSelected)

		// Other
		default:
			os.Rename(fileSelected, "other/"+fileSelected)
		}

	}
}
