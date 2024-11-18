package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	var sourceString string = "Hello, OTUS!"
	sourceString = reverse.String(sourceString)
	fmt.Println(sourceString)
}
