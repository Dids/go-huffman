package main

// Import the testing package
import (
	"testing"
)

func testThings(t *testing.T) {
	if len(testResult) != 0 {
		t.Fail()
	}

	Init()

	TestHuffman()

	if len(testResult) == 0 {
		t.Fail()
	}
}
