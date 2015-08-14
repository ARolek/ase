package main

import (
	"os"

	"github.com/ARolek/ase"
)

var testColors = []ase.Color{
	ase.Color{
		Name:   "RGB",
		Model:  "RGB",
		Values: []float32{1, 1, 1},
		Type:   "Normal",
	},
	ase.Color{
		Name:   "Grayscale",
		Model:  "CMYK",
		Values: []float32{0, 0, 0, 0.47},
		Type:   "Spot",
	},
	ase.Color{
		Name:   "cmyk",
		Model:  "CMYK",
		Values: []float32{0, 1, 0, 0},
		Type:   "Spot",
	},
	ase.Color{
		Name:   "LAB",
		Model:  "RGB",
		Values: []float32{0, 0.6063648, 0.524658},
		Type:   "Global",
	},
	ase.Color{
		Name:   "PANTONE P 1-8 C",
		Model:  "LAB",
		Values: []float32{0.9137255, -5, 94},
		Type:   "Spot",
	},
}

var testGroup = ase.Group{
	Name: "A Color Group",
	Colors: []ase.Color{
		ase.Color{
			Name:   "Red",
			Model:  "RGB",
			Values: []float32{1, 0, 0},
			Type:   "Global",
		},
		ase.Color{
			Name:   "Green",
			Model:  "RGB",
			Values: []float32{0, 1, 0},
			Type:   "Global",
		},
		ase.Color{
			Name:   "Blue",
			Model:  "RGB",
			Values: []float32{0, 0, 1},
			Type:   "Global",
		},
	},
}

func main() {
	// Initialize a sample ASE
	sampleAse := ase.ASE{}
	sampleAse.Colors = testColors
	sampleAse.Groups = append(sampleAse.Groups, testGroup)

	// Create the file to write the encoded ASE
	f, err := os.Create("./encoded.ase")
	if err != nil {
		panic(err)
	}

	// It’s idiomatic to defer a Close immediately after opening a file.
	defer f.Close()

	//	encode our ASE file
	if err = ase.Encode(sampleAse, f); err != nil {
		panic(err)
	}
}
