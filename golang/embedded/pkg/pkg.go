package pkg

import (
	"embed"
)

//go:embed folder/single_file.txt
var FileString string

//go:embed folder/single_file.txt
var FileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var Folder embed.FS
