package folder

// init all folder that needed in first run
func InitAllFolder() {
	MakeFolder("Document")
	MakeFolder("Audio")
	MakeFolder("Video")
	MakeFolder("Image")
	MakeFolder("Application")
	MakeFolder("Compressed")
	MakeFolder("Compressed/RAR")
	MakeFolder("Compressed/7z")
	MakeFolder("Compressed/ZIP")
	MakeFolder("Unknown")
}
