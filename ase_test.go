package ase

import (
	"log"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	testFile := "testfiles/test.ase"

	//	open our test file
	f, err := os.Open(testFile)
	if err != nil {
		t.Error(err)
	}

	//	test our decode
	ase, err := Decode(f)
	if err != nil {
		t.Error(err)
	}

	//	log our output
	log.Printf("%+v\n", ase)
}

func TestDecodeFile(t *testing.T) {
	testFile := "testfiles/test.ase"

	ase, err := DecodeFile(testFile)
	if err != nil {
		t.Error(err)
	}

	//	log our output
	log.Printf("%+v\n", ase)
}

func TestSignature(t *testing.T) {
	testFile := "testfiles/test.ase"

	ase, err := DecodeFile(testFile)
	if err != nil {
		t.Error(err)
	}

	if ase.Signature() != "ASEF" {
		t.Error("file signature is invalid")
	}
}
