package main

import (
	"fmt"
	"reflect"
)

func okay(d any) {
	dPtr := reflect.ValueOf(d)
	fmt.Printf("VALUE: %v\n", dPtr)

	if dPtr.Kind() != reflect.Ptr || dPtr.Elem().Kind() != reflect.Struct {
		fmt.Println("Expected a pointer to a struct")
		return
	}

	dV := dPtr.Elem()
	dT := reflect.TypeOf(d)
	fmt.Printf("TYPE: %v\n", dT.Kind())
	fmt.Printf("Value: %v\n", dV)
	fmt.Println("WHOLE STRUCT:", dT.Elem())

	for i := 0; i < dT.Elem().NumField(); i++ {
		field := dT.Elem().Field(i)
		fmt.Printf("FIELD: %v\n", field)

		key := field.Tag.Get("mukezhz")
		kind := field.Type.Kind()

		fmt.Printf("KEY: %v\n", key)
		fmt.Printf("KIND: %v\n", kind)

		fV := dV.Field(i)
		fmt.Println("THE VALUE:", fV)
		if kind == reflect.String {
			fV.SetString(key)
		}
	}

	dPtr = reflect.ValueOf(d)
	fmt.Printf("VALUE: %v\n", dPtr)
}

type TEST struct {
	Name string `mukezhz:"this is datas"`
}

// FOR MORE: https://articles.wesionary.team/reflections-tutorial-query-string-to-struct-parser-in-go-b2f858f99ea1
func main() {
	// MUST PASS POINTER: TO insure the value you're modifying is settable
	// reflection doesn't work in the original value
	// on passing the pointer it knows the address of the value
	test := TEST{
		Name: "Hellou",
	}
	okay(&test)
	fmt.Println("AFTER CHANGE:", test.Name)
}
