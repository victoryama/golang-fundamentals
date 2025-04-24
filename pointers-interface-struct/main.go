package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 123
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	b := &a
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
	c := *b
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))
}
