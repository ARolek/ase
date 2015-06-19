# ase
ASE decoder

Package for decoding ASE (Adobe Swatch Exchange) files into a struct.

The ASE specification can be found [here](http://www.selapa.net/swatches/colors/fileformats.php#adobe_ase).

Note: this was one of my first projects in Go as well as with binary decoding. It's in desperate need of a refactor. Will hopefully have time to make updates soon. 

# Install

`$ go get github.com/arolek/ase`

# Getting started

the ase package exposes a Decode method. You simply create a new ASE struct and pass a string to the Decode method of a local file location. For example:

```go
package main

import (
	"log"

	"github.com/arolek/ase"
)

func main() {
	filePath := "/path/to/test.ase"

	aseDecoder := ase.ASE{}
	if err := aseDecoder.Decode(filePath, false); err != nil {
		log.Println(err)
	}

	log.Printf("%+v\n", aseDecoder)
}

```

# todo

- clean up logs
- implement Encode (struct -> .ase file)
- implement JSON encoding
- update documentation
- implement tests
