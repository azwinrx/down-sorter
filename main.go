package main

import (
	"down-sorter/folder"
	"down-sorter/sorter"
)

func main() {
	folder.InitAllFolder()
	sorter.SorterFile()
}
