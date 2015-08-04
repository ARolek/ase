# ase
ASE decoder

Package for decoding ASE (Adobe Swatch Exchange) files into a struct.

The ASE specification can be found [here](http://www.selapa.net/swatches/colors/fileformats.php#adobe_ase).

Note: this was one of my first projects in Go as well as with binary decoding. It's in desperate need of a refactor. Will hopefully have time to make updates soon. 

# Install

`$ go get github.com/arolek/ase`

# Getting started

the ASE package exposes a Decode method. You simply pass an io.Reader interface to ase.Decode and it will return an ASE struct of the decoded data. For example:

```go
package main

import (
	"log"

	"github.com/arolek/ase"
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

# todo

- clean up logs
- implement Encode (struct -> .ase file)
- implement JSON encoding
- update documentation
- implement tests
