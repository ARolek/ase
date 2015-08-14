# ase
Golang package for decoding and encoding ASE (Adobe Swatch Exchange) files.

The ASE specification can be found [here](http://www.selapa.net/swatches/colors/fileformats.php#adobe_ase).

# Install

`$ go get github.com/arolek/ase`

# Getting started

The ASE package exposes a Decode and Encode method. You simply pass an io.Reader interface to ase.Decode and it will return an ASE struct of the decoded data. For convenience, a DecodeFile method is available to decode an existing ASE file. For encoding, simply initialize an ASE struct and populate it with the appropriate Groups and Colors data.

## Examples

### Decoding
```go
package main

import (
	"log"

	"github.com/ARolek/ase"
)

func main() {
	//	open the file
	f, err := os.Open("/path/to/test.ase")
	if err != nil {
		log.Println(err)
	}

	//	decode can take in any io.Reader
	ase, err := ase.Decode(f)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%+v\n", ase)
}
```

## Encoding
```go
package main

import (
	"bufio"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Initialize a sample ASE
	sampleAse := ase.ASE{}
	sampleAse.Colors = testColors
	sampleAse.Groups = append(sampleAse.Groups, testGroup)

	// Create the file to write the encoded ASE
	f, err := os.Create("./encoded.ase")
	check(err)

	// It’s idiomatic to defer a Close immediately after opening a file.
	defer f.Close()

	// Create a new writer
	w := bufio.NewWriter(f)
	
	// Encode the ASE 
	ase.Encode(sampleAse, w)

	// Use Flush to ensure all buffered operations have been
	// applied to the underlying writer.
	w.Flush()
}
```

# todo
- implement JSON encoding
