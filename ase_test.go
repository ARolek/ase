package ase

import "testing"

func TestASEStruct(t *testing.T) {
	ase := ASE{}

	var defaultSignature string
	var defaultVersion [2]int16
	var defaultNumBlocks [1]int32

	if ase.Signature != defaultSignature {
		t.Error("expected ase to have default Signature of", defaultSignature)
	}
	if ase.Version != defaultVersion {
		t.Error("expected ase to have default Version of", defaultVersion)
	}
	if ase.NumBlocks != defaultNumBlocks {
		t.Error("expected ase to have default NumBlocks of", defaultNumBlocks)
	}
	if ase.Colors != nil {
		t.Error("expected ase to have default Colors as nil")
	}
	if ase.Groups != nil {
		t.Error("expected ase to have default Groups as nil")
	}
}

func TestDecode(t *testing.T) {

	// Load and decode the ase test file
	filePath := "testfiles/test.ase"
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
		NameLen: 12,
		Name:    "PANTONE 877",
		Model:   "CMYK",
		Values:  []float32{0.22, 0.17, 0.13, 0.4},
		Type:    "Spot",
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
			"got ", expectedColor.Type)
	}
}
