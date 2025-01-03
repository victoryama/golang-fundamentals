package main

import (
	"fmt"
	"os"
	"strings"
)

// echo3
func main() {

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}
