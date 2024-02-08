package main

import "fmt"

// generic in interface
type Data interface {
	int | float32
}

type User[T Data] struct {
	Data T
	Name string
	Age  int
}

// generic with struct
type IDisplay[T string] interface {
	Display(T)
}
type CustomStruct[T any] struct {
	Data T
}

func (c CustomStruct[T]) Display(s T) {
	fmt.Println(s)
}

// COMPLEX EXAMPLES
// generic in interface
type Getter[T any] interface {
	Get() T
}

type MyStruct struct {
	Val string
}

// implements Getter[string]
func (m MyStruct) Get() string {
	return m.Val
}

// we can return MyStruct since MyStruct has implemented Getter[string]
func foo() Getter[string] {
	return MyStruct{
		Val: "Hello",
	}
}

// generic in map
type CustomMap[K int | float32, V string] map[K]V

func main() {
	u := User[int]{Name: "John", Age: 20, Data: 10}
	fmt.Println(u)
	u2 := User[float32]{Name: "John", Age: 20, Data: 10.5}
	fmt.Println(u2)

	m := CustomMap[int, string]{1: "a", 2: "b"}
	fmt.Println(m)
	m2 := CustomMap[float32, string]{1.1: "a", 2.2: "b"}
	fmt.Println(m2)

	c := CustomStruct[int]{Data: 100}
	fmt.Println(c)

	c2 := CustomStruct[float32]{Data: 100.5}
	fmt.Println(c2)

	c3 := CustomStruct[string]{Data: "Hello"}
	c3.Display("Hey")
	fmt.Println(c3)

	c4 := CustomStruct[bool]{Data: true}
	c4.Display(true)

	// COMPLEX EXAMPLES
	z := foo()
	fmt.Println(z.Get())
}
