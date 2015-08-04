package ase

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

var (
	ErrInvalidFile = errors.New("ase: file not an ASE file")
)

//	ASE File Spec http://www.selapa.net/swatches/colors/fileformats.php#adobe_ase

//	Decode a valid ASE input
func Decode(r io.Reader) (ase ASE, err error) {

	//	signature
	if err = ase.readSignature(r); err != nil {
		return
	}

	//	file version
	if err = ase.readVersion(r); err != nil {
		return
	}

	//	block count
	if err = ase.readNumBlock(r); err != nil {
		return
	}

	//	if we encounter groups, store a ref here
	var g Group

	//	itereate based on our block count
	for i := 0; i < int(ase.numBlocks); i++ {
		//	new block
		b := block{}

		//	decode the block container
		b.Read(r)

		//	switch on block type
		switch b.Type {
		case color:
			c := Color{}
			if err = c.read(r); err != nil {
				return
			}

			//	if we have a group, add color to the group
			if g.Name != "" {
				g.Colors = append(g.Colors, c)
			} else {
				//	color is not in a group. add to color slice
				ase.Colors = append(ase.Colors, c)
			}

			break
		case groupStart:
			//	new group
			g = Group{}

			//	read the group
			g.read(r)

			break
		case groupEnd:
			//	add the group to our ase struct
			ase.Groups = append(ase.Groups, g)

			//	reset our group struct
			g = Group{}
			break
		default:
			err = ErrInvalidBlockType
			return
		}
	}

	return
}

//	read a file on a file system
func DecodeFile(file string) (ase ASE, err error) {
	//	open the file
	f, err := os.Open(file)
	if err != nil {
		return
	}

	//	decode the file
	return Decode(f)
}

//	TODO: complete encode method
func Encode(ase ASE, w io.Writer) (err error) {

	//	write signature

	//	write version

	//	write number of blocks

	//	write details of each block

	return
}

type ASE struct {
	signature [4]uint8
	version   [2]int16
	numBlocks int32
	Colors    []Color
	Groups    []Group
}

//	ASE Files start are signed with ASEF at the beginning of the file
//	let's make sure this is an ASE file
func (ase *ASE) readSignature(r io.Reader) (err error) {
	//	read the signature
	if err = binary.Read(r, binary.BigEndian, &ase.signature); err != nil {
		return
	}

	//	check our file signature
	if string(ase.signature[0:]) != "ASEF" {
		return ErrInvalidFile
	}

	return
}

//	ASE version. Should be 1.0
func (ase *ASE) readVersion(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, &ase.version)
}

//	Total number of blocks in the ASE file
func (ase *ASE) readNumBlock(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, &ase.numBlocks)
}

//	returns the file signature in human readable format
func (ase *ASE) Signature() string {
	return string(ase.signature[0:])
}
