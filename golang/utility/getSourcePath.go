package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	fmt.Println(GetSourcePath())
}

func GetSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
