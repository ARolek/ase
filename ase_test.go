package ase

import (
	"bytes"
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

func TestVersion(t *testing.T) {
	testFile := "testfiles/test.ase"

	ase, err := DecodeFile(testFile)
	if err != nil {
		t.Error(err)
	}

	if ase.Version() != "1.0" {
		t.Error("did not get version 1.0, got:", ase.Version())
	}
}

func TestEncode(t *testing.T) {

	// Initialize a sample ASE
	sampleAse := ASE{}

	// Encode the sampleAse into the buffer and immediately decode it.
	b := new(bytes.Buffer)
	Encode(sampleAse, b)
	ase, _ := Decode(b)

	// Check the ASE's decoded values.
	if string(ase.signature[0:]) != "ASEF" {
		t.Error("ase: file not an ASE file")
	}

	if ase.version[0] != 1 && ase.version[1] != 0 {
		t.Error("ase: version is not 1.0")
	}
}
