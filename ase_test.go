package ase

import "testing"

func TestDecode(t *testing.T) {

	filePath := "samples/test.ase"
	ase := ASE{}
	ase.Decode(filePath, false)

	expectedNumBlocks := [1]int32{10}
	if ase.NumBlocks != expectedNumBlocks {
		t.Error("expected NumBlocks of ", expectedNumBlocks,
			" got ", ase.NumBlocks)
	}

	// Colors
	expectedColors := []Color{
		Color{
			NameLen: 4,
			Name:    "RGB",
			Model:   "RGB",
			Values:  []float32{1, 1, 1},
			Type:    "Normal",
		},
		Color{
			NameLen: 10,
			Name:    "Grayscale",
			Model:   "CMYK",
			Values:  []float32{0, 0, 0, 0.47},
			Type:    "Spot",
		},
		Color{
			NameLen: 5,
			Name:    "cmyk",
			Model:   "CMYK",
			Values:  []float32{0, 1, 0, 0},
			Type:    "Spot",
		},
		Color{
			NameLen: 4,
			Name:    "LAB",
			Model:   "RGB",
			Values:  []float32{0, 0.6063648, 0.524658},
			Type:    "Global",
		},
		Color{
			NameLen: 16,
			Name:    "PANTONE P 1-8 C",
			Model:   "LAB",
			Values:  []float32{0.9137255, -5, 94},
			Type:    "Spot",
		},
		Color{
			NameLen: 4,
			Name:    "Red",
			Model:   "RGB",
			Values:  []float32{1, 0, 0},
			Type:    "Global",
		},
		Color{
			NameLen: 6,
			Name:    "Green",
			Model:   "RGB",
			Values:  []float32{0, 1, 0},
			Type:    "Global",
		},
		Color{
			NameLen: 5,
			Name:    "Blue",
			Model:   "RGB",
			Values:  []float32{0, 0, 1},
			Type:    "Global",
		},
	}

	expectedNumColors := len(expectedColors)
	actualNumColors := len(ase.Colors)

	if actualNumColors != expectedNumColors {
		t.Error("expected number of colors to be", expectedNumColors,
			"got ", actualNumColors)
	}

	for i, color := range ase.Colors {
		expectedColor := expectedColors[i]

		if color.NameLen != expectedColor.NameLen {
			t.Error("expected initial color of name length ", expectedColor.NameLen,
				"got ", color.NameLen)
		}

		if color.Name != expectedColor.Name {
			t.Error("expected initial color with name ", expectedColor.Name,
				"got ", color.Name)
		}

		if color.Model != expectedColor.Model {
			t.Error("expected initial color of Model ", expectedColor.Model,
				"got ", color.Model)
		}

		for j, _ := range expectedColor.Values {
			if color.Values[j] != expectedColor.Values[j] {
				t.Error("expected color value ", expectedColor.Values[j],
					"got ", color.Values[j])
			}
		}

		if color.Type != expectedColor.Type {
			t.Error("expected color type ", expectedColor.Type,
				"got ", color.Type)
		}
	}

	// Groups
	expectedGroupLen := 0
	actualGroupLen := len(ase.Groups)

	if expectedGroupLen != actualGroupLen {
		t.Error("expected group length of ", expectedGroupLen,
			"got ", actualGroupLen)
	}

}

func TestDecodeWithGroup(t *testing.T) {

	filePath := "samples/test.ase"
	ase := ASE{}
	ase.Decode(filePath, true)

	expectedNumBlocks := [1]int32{10}
	if ase.NumBlocks != expectedNumBlocks {
		t.Error("expected NumBlocks of ", expectedNumBlocks,
			" got ", ase.NumBlocks)
	}

	// Colors
	expectedColors := []Color{
		Color{
			NameLen: 4,
			Name:    "RGB",
			Model:   "RGB",
			Values:  []float32{1, 1, 1},
			Type:    "Normal",
		},
		Color{
			NameLen: 10,
			Name:    "Grayscale",
			Model:   "CMYK",
			Values:  []float32{0, 0, 0, 0.47},
			Type:    "Spot",
		},
		Color{
			NameLen: 5,
			Name:    "cmyk",
			Model:   "CMYK",
			Values:  []float32{0, 1, 0, 0},
			Type:    "Spot",
		},
		Color{
			NameLen: 4,
			Name:    "LAB",
			Model:   "RGB",
			Values:  []float32{0, 0.6063648, 0.524658},
			Type:    "Global",
		},
		Color{
			NameLen: 16,
			Name:    "PANTONE P 1-8 C",
			Model:   "LAB",
			Values:  []float32{0.9137255, -5, 94},
			Type:    "Spot",
		},
	}

	expectedNumColors := len(expectedColors)
	actualNumColors := len(ase.Colors)

	if actualNumColors != expectedNumColors {
		t.Error("expected number of colors to be", expectedNumColors,
			"got ", actualNumColors)
	}

	for i, color := range ase.Colors {
		expectedColor := expectedColors[i]

		if color.NameLen != expectedColor.NameLen {
			t.Error("expected initial color of name length ", expectedColor.NameLen,
				"got ", color.NameLen)
		}

		if color.Name != expectedColor.Name {
			t.Error("expected initial color with name ", expectedColor.Name,
				"got ", color.Name)
		}

		if color.Model != expectedColor.Model {
			t.Error("expected initial color of Model ", expectedColor.Model,
				"got ", color.Model)
		}

		for j, _ := range expectedColor.Values {
			if color.Values[j] != expectedColor.Values[j] {
				t.Error("expected color value ", expectedColor.Values[j],
					"got ", color.Values[j])
			}
		}

		if color.Type != expectedColor.Type {
			t.Error("expected color type ", expectedColor.Type,
				"got ", color.Type)
		}
	}

	// Groups
	expectedGroups := []Group{
		Group{
			NameLen: 14,
			Name:    "A Color Group",
			Colors: []Color{
				Color{
					NameLen: 4,
					Name:    "Red",
					Model:   "RGB",
					Values:  []float32{1, 0, 0},
					Type:    "Global",
				},
				Color{
					NameLen: 6,
					Name:    "Green",
					Model:   "RGB",
					Values:  []float32{0, 1, 0},
					Type:    "Global",
				},
				Color{
					NameLen: 5,
					Name:    "Blue",
					Model:   "RGB",
					Values:  []float32{0, 0, 1},
					Type:    "Global",
				},
			},
		},
	}
	expectedGroupsLen := len(expectedGroups)
	actualGroupLen := len(ase.Groups)

	if expectedGroupsLen != actualGroupLen {
		t.Error("expected group length of ", expectedGroupsLen,
			"got ", actualGroupLen)
	}

	expectedGroup := expectedGroups[0]
	expectedGroupName := "A Color Group"
	actualGroup, ok := ase.Groups[expectedGroupName]

	if !ok {
		t.Error("expected group name to exist: ", expectedGroupName)
	}

	for i, color := range actualGroup {
		expectedColor := expectedGroup.Colors[i]

		if color.NameLen != expectedColor.NameLen {
			t.Error("expected initial color of name length ", expectedColor.NameLen,
				"got ", color.NameLen)
		}

		if color.Name != expectedColor.Name {
			t.Error("expected initial color with name ", expectedColor.Name,
				"got ", color.Name)
		}

		if color.Model != expectedColor.Model {
			t.Error("expected initial color of Model ", expectedColor.Model,
				"got ", color.Model)
		}

		for j, _ := range expectedColor.Values {
			if color.Values[j] != expectedColor.Values[j] {
				t.Error("expected color value ", expectedColor.Values[j],
					"got ", color.Values[j])
			}
		}

		if color.Type != expectedColor.Type {
			t.Error("expected color type ", expectedColor.Type,
				"got ", color.Type)
		}
	}

}
