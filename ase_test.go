package ase

import "testing"

func TestDecode(t *testing.T) {

	// Load and decode the ase test file
	filePath := "samples/test.ase"
	ase := ASE{}
	ase.Decode(filePath, false)

	// Check that each respective field is correctly populated
	expectedSignature := "ASEF"
	if ase.Signature != expectedSignature {
		t.Error("expected signature of ASEF, got ", ase.Signature)
	}

	expectedVersion := [2]int16{1, 0}
	if ase.Version != expectedVersion {
		t.Error("expected version of ", expectedSignature,
			" got ", ase.Signature)
	}

	expectedNumBlocks := [1]int32{10}
	if ase.NumBlocks != expectedNumBlocks {
		t.Error("expected NumBlocks of ", expectedNumBlocks,
			" got ", ase.NumBlocks)
	}

	// Colors
	expectedColor := Color{
		NameLen: 4,
		Name:    "RGB",
		Model:   "RGB",
		Values:  []float32{1, 1, 1},
		Type:    "Normal",
	}

	firstColor := ase.Colors[0]

	if firstColor.NameLen != expectedColor.NameLen {
		t.Error("expected initial color of name length ", expectedColor.NameLen,
			"got ", firstColor.NameLen)
	}

	if firstColor.Name != expectedColor.Name {
		t.Error("expected initial color with name ", expectedColor.Name,
			"got ", firstColor.Name)
	}

	if firstColor.Model != expectedColor.Model {
		t.Error("expected initial color of Model ", expectedColor.Model,
			"got ", firstColor.Model)
	}

	for i, _ := range expectedColor.Values {
		if firstColor.Values[i] != expectedColor.Values[i] {
			t.Error("expected color value ", expectedColor.Values[i],
				"got ", firstColor.Values[i])
		}
	}

	if firstColor.Type != expectedColor.Type {
		t.Error("expected color type ", expectedColor.Type,
			"got ", firstColor.Type)
	}

}
