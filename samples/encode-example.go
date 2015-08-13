package main

import (
	"bufio"
	"os"

	"github.com/francismakes/ase"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Load the test.ase
	testFile := "./test.ase"
	sampleAse, err := DecodeFile(testFile)
	
	// Create a file
	f, err := os.Create("./encoded.ase")
	check(err)

	// It’s idiomatic to defer a Close immediately after opening a file.
	defer f.Close()

	// Create a new writer
	w := bufio.NewWriter(f)
	ase.Encode(sampleAse, w)

	// Use Flush to ensure all buffered operations have been
	// applied to the underlying writer.
	w.Flush()
}
