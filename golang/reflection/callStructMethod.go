package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	u := &User{"Tom", 20}
	println(u.GetName())

	// call the method by reflection
	v := reflect.ValueOf(u)
	getName := v.MethodByName("GetName")
	ret := getName.Call(nil)
	fmt.Println(ret)

	fmt.Println("Before set name:", u.Name)

	setName := v.MethodByName("SetName")
	setName.Call([]reflect.Value{reflect.ValueOf("Jerry")})

	fmt.Println("After set name:", u.Name)
}
