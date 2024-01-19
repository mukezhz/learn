package main

import (
	"github.com/mukezhz/golang/design_patterns/creational/singleton"
)

func main() {
	s := singleton.GetInstance()
	s.SetName("Mukesh")
	println(s.GetName())
}
