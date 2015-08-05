package ase

import (
	"bytes"
	"log"
	"os"
	"testing"
)

var testColors = []Color{
	Color{
		Name:   "RGB",
		Model:  "RGB",
		Values: []float32{1, 1, 1},
		Type:   "Normal",
	},
	Color{
		Name:   "Grayscale",
		Model:  "CMYK",
		Values: []float32{0, 0, 0, 0.47},
		Type:   "Spot",
	},
	Color{
		Name:   "cmyk",
		Model:  "CMYK",
		Values: []float32{0, 1, 0, 0},
		Type:   "Spot",
	},
	Color{
		Name:   "LAB",
		Model:  "RGB",
		Values: []float32{0, 0.6063648, 0.524658},
		Type:   "Global",
	},
	Color{
		Name:   "PANTONE P 1-8 C",
		Model:  "LAB",
		Values: []float32{0.9137255, -5, 94},
		Type:   "Spot",
	},
	Color{
		Name:   "Red",
		Model:  "RGB",
		Values: []float32{1, 0, 0},
		Type:   "Global",
	},
	Color{
		Name:   "Green",
		Model:  "RGB",
		Values: []float32{0, 1, 0},
		Type:   "Global",
	},
	Color{
		Name:   "Blue",
		Model:  "RGB",
		Values: []float32{0, 0, 1},
		Type:   "Global",
	},
}

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
	sampleAse.Colors = testColors

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

	if ase.numBlocks != 8 {
		t.Error("ase: expected 8 blocks to be present")
	}

}
