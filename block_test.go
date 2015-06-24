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

func TestBlockRead(t *testing.T) {
    t.SkipNow()
}