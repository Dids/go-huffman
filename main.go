package main

import (
	"fmt"
)

// Variable that holds our test result
var testResult string

/* Capitalized functions are exported and should include a comment */

// Init ...
func Init() {
	testResult = "main.go initialized"
}

// Primary program
func main() {
	Init()
	testResult = "main.go starting"
	fmt.Println("Yo, main.go says hi.")
	TestHuffman()
	testResult = "main.go finished"
}
