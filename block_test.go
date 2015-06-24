package ase

import (
	"testing"
)

func TestBlockStruct(t *testing.T) {
	var b Block
	var defaultType [2]uint8
	var defaultLength [1]int32

	if b.Type != defaultType {
		t.Error("expected block to have default Type of", defaultType)
	}

	if b.Length != defaultLength {
		t.Error("expected block to have default Type of", defaultLength)
	}
}

// The main challenge here is to deconstruct
// what is assumed in how ase is calling this.

// Right now, ase uses block in the following way
// It reads the filebuffer from bytes.NewReader(file)
// and puts that data into the block itself
// from ther block.Type is assigned
// and it should also be that block has a length as well
// how about doing a direct read for now and see where it goes?
func TestBlockRead(t *testing.T) {
    t.SkipNow()
}