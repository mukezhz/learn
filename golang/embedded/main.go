package main

import "github.com/mukezhz/learn/golang/embedded/pkg"

func main() {

	print("Embed file to variable[string]: ", pkg.FileString)
	print("Embed file to variable[bytes]: ", string(pkg.FileByte))

	// pkg.Folder -> temporary file system
	content1, _ := pkg.Folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := pkg.Folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
