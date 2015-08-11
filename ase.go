package ase

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
	"log"
	"unicode/utf16"
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

		ase.Blocks = append(ase.Blocks, b)

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

type ASE struct {
	signature [4]uint8
	version   [2]int16
	numBlocks int32
	Colors    []Color
	Groups    []Group
	Blocks    []block
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

func Encode(ase ASE, w io.Writer) (err error) {

	if err = ase.writeSignature(w); err != nil {
		return
	}

	if err = ase.writeVersion(w); err != nil {
		return
	}

	if err = ase.writeNumBlock(w); err != nil {
		return
	}



	//	itereate based on our block count
	for i := 0; i < int(ase.numBlocks); i++ {
		b := ase.Blocks[i]
		if err = ase.writeBlockType(w, ase.Blocks[i]); err != nil {
			return
		}

		if err = ase.writeBlockLength (w, ase.Blocks[i]); err != nil {
			return
		}




		switch b.Type {
			case color:
				if err = ase.writeColorName(w, ase.Colors[i]); err != nil {
					return
				}

				if err = ase.writeColorModel(w, ase.Colors[i]); err != nil {
					return
				}

				if err = ase.writeColorValues(w, ase.Colors[i]); err != nil {
					return
				}

				if err = ase.writeColorType(w, ase.Colors[i]); err != nil {
					return
				}


				break
			case groupStart:


				break
			case groupEnd:

				break
			default:
				err = ErrInvalidBlockType
				return
		}

	}



	var numOfColors = len(ase.Colors)
	var numOfGroups = len(ase.Groups)
	log.Print("Number of colors ", numOfColors)
	log.Print("Number of groups ", numOfGroups)
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

func (ase *ASE) writeSignature(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, ase.signature)
}

func (ase *ASE) writeVersion(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, ase.version)
}

func (ase *ASE) writeNumBlock(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, ase.numBlocks)
}

func (ase *ASE) writeBlockType(w io.Writer, b block) error {
	return binary.Write(w, binary.BigEndian, b.Type)
}

func (ase *ASE) writeBlockLength(w io.Writer, b block) error {
	return binary.Write(w, binary.BigEndian, b.Length)
}

func (ase *ASE) writeColorName(w io.Writer, c Color) error {
	colorNameSlice := []rune(c.Name)
	colorNameSlice = append(colorNameSlice, 0)
	colorName := utf16.Encode(colorNameSlice)
	return binary.Write(w, binary.BigEndian, colorName)
}

func (ase *ASE) writeColorModel(w io.Writer, c Color) error {
	colorModelSlice := []rune(c.Model)
	colorModel := utf16.Encode(colorModelSlice)
	return binary.Write(w,binary.BigEndian, colorModel)
}

func (ase *ASE) writeColorValues(w io.Writer, c Color) error {
	var err error
	for _, cv := range c.Values {
		err = binary.Write(w,binary.BigEndian, cv)

		if(err != nil){
			return err
		}
	}
	return err
}

func (ase *ASE) writeColorType(w io.Writer, c Color) error {
	var cType int16
	switch {
		case c.Type == "Global":
			cType = 0
			break
		case c.Type == "Spot":
			cType = 1
			break
		case c.Type ==  "Normal":
			cType = 2
			break
		default:
			return ErrInvalidColorType

	}
	return binary.Write(w, binary.BigEndian, cType)
}


