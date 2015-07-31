package ase

import (
	"log"
	"os"
	"testing"
)

const testFile = "testfiles/test.ase"

func TestDecode(t *testing.T) {
	//	open our test file
	f, err := os.Open(testFile)
	if err != nil {
		t.Error(err)
	}

	//	test our decode
	ase, err := Decode(f, false)
	if err != nil {
		t.Error(err)
	}

	//	log our output
	log.Printf("%+v\n", ase)
}

func TestDecodeFile(t *testing.T) {
	ase, err := DecodeFile(testFile, false)
	if err != nil {
		t.Error(err)
	}

	//	log our output
	log.Printf("%+v\n", ase)
}
