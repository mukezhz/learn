package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello world from Go Web Assembly")
	js.Global().Set("goLang", "Go Lang")
}
