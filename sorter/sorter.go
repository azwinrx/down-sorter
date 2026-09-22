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
			os.Rename(fileSelected, "Document/txt/"+fileSelected)
		case ".docx", ".doc", ".odt":
			os.Rename(fileSelected, "Document/word/"+fileSelected)
		case ".pptx", ".ppt", ".odp":
			os.Rename(fileSelected, "Document/ppt/"+fileSelected)
		case ".xlsx", ".xls", ".ods":
			os.Rename(fileSelected, "Document/excel/"+fileSelected)
		case ".pdf":
			os.Rename(fileSelected, "Document/pdf/"+fileSelected)

		// Audio
		case ".aac":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".ac3":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".aiff":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".amr":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".au":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".flac":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".mid":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".mka":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".mp3":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".ogg":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".ra":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".voc":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".wav":
			os.Rename(fileSelected, "Audio/"+fileSelected)
		case ".wma":
			os.Rename(fileSelected, "Audio/"+fileSelected)

		// Video
		case ".avi":
			os.Rename(fileSelected, "Video/"+fileSelected)
		case ".flv":
			os.Rename(fileSelected, "Video/"+fileSelected)
		case ".mkv":
			os.Rename(fileSelected, "Video/"+fileSelected)
		case ".mov":
			os.Rename(fileSelected, "Video/"+fileSelected)
		case ".mp4":
			os.Rename(fileSelected, "Video/"+fileSelected)
		case ".mpg":
			os.Rename(fileSelected, "Video/"+fileSelected)
		case ".swf":
			os.Rename(fileSelected, "Video/"+fileSelected)
		case ".webm":
			os.Rename(fileSelected, "Video/"+fileSelected)
		case ".wmv":
			os.Rename(fileSelected, "Video/"+fileSelected)

		// Image
		case ".ai":
			os.Rename(fileSelected, "Image/"+fileSelected)
		case ".ico":
			os.Rename(fileSelected, "Image/"+fileSelected)
		case ".jpeg":
			os.Rename(fileSelected, "Image/"+fileSelected)
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
		case ".tiff":
			os.Rename(fileSelected, "Image/"+fileSelected)
		case ".webp":
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

		// Font
		case ".ttf":
			os.Rename(fileSelected, "Font/"+fileSelected)
		case ".otf":
			os.Rename(fileSelected, "Font/"+fileSelected)
		case ".woff":
			os.Rename(fileSelected, "Font/"+fileSelected)

		// Other
		default:
			os.Rename(fileSelected, "Other/"+fileSelected)
		}

	}
}
