package ase

import (
	"log"
	"os"
	"testing"
	"reflect"
)

func TestDecode(t *testing.T) {
	testFile := "samples/test.ase"

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
	testFile := "samples/test.ase"

	ase, err := DecodeFile(testFile)
	if err != nil {
		t.Error(err)
	}

	//	log our output
	log.Printf("%+v\n", ase)
}

func TestSignature(t *testing.T) {
	testFile := "samples/test.ase"

	ase, err := DecodeFile(testFile)
	if err != nil {
		t.Error(err)
	}

	if ase.Signature() != "ASEF" {
		t.Error("file signature is invalid")
	}
}


func TestEncode(t *testing.T) {
	//  open the file
	f, err := os.Open("samples/test.ase")
	if err != nil {
		t.Error(err)
	}

	output, err := Decode(f)
	if err != nil {
		t.Error(err)
	}

	f.Close()

	fo, err := os.Create("samples/test-4.ase")

	err = Encode(output, fo)

	if err != nil {
		t.Error(err)
	}

	fo.Close()


	//Decode previously created *.ase file
	fTest4, err := os.Open("samples/test-4.ase")

	if err != nil {
		t.Error(err)
	}

	outputTest4, err := Decode(fTest4)
	if err != nil {
		t.Error(err)
	}

	//Ensure they are equal
	if isEqual := reflect.DeepEqual(output, outputTest4); isEqual == false {
		t.Error("Did not produce identical structs")
	}

	fTest4.Close()

}
