package main

import (
	"fmt"
	"os"
)

// echo2
func main() {

	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

}
