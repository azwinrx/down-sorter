package folder

// init all folder that needed in first run
func InitAllFolder() {
	// Document
	MakeFolder("Document")
	MakeFolder("Document/txt")
	MakeFolder("Document/word")
	MakeFolder("Document/ppt")
	MakeFolder("Document/excel")
	MakeFolder("Document/pdf")

	// Audio
	MakeFolder("Audio")

	// Video
	MakeFolder("Video")

	// Image
	MakeFolder("Image")

	// Application
	MakeFolder("Application")

	// Compressed
	MakeFolder("Compressed")
	MakeFolder("Compressed/ZIP")
	MakeFolder("Compressed/RAR")
	MakeFolder("Compressed/7z")
	MakeFolder("Compressed/TAR")
	MakeFolder("Compressed/GZ")

	// Font
	MakeFolder("Font")

	// Other
	MakeFolder("Other")
}
