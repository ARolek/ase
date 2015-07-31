package ase

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	ErrInvalidFile = errors.New("ase: file not an ASE file")
)

//	ASE File Spec http://www.selapa.net/swatches/colors/fileformats.php#adobe_ase

//	pass in the path to the ASE file to read and wheater or not to group
//	if group is set to false, colors in groups will be listed the same way
//	as colors outside of groups.
//	if group is set to true, colors will be nested in a group as named in the
//	ASE file
func Decode(reader io.Reader, group bool) (ase ASE, err error) {

	//	signature
	if err = ase.readSignature(reader); err != nil {
		return
	}

	//	file version
	if err = ase.readVersion(reader); err != nil {
		return
	}

	//	block count
	if err = ase.readNumBlock(reader); err != nil {
		return
	}

	groupName := ""
	//	itereate based on our block count
	for i := 0; i < int(ase.NumBlocks[0]); i++ {
		block := new(Block)
		block.Read(reader)
		switch fmt.Sprintf("%x", block.Type) {
		case "0001":
			color := new(Color)
			if err = color.Read(reader); err != nil {
				return
			}

			if groupName != "" && group == true {
				ase.Groups[groupName] = append(ase.Groups[groupName], *color)
			} else {
				ase.Colors = append(ase.Colors, *color)
			}
			break
		case "c001":
			//	if this is our first group entry, we need to init our groups map
			if ase.Groups == nil {
				ase.Groups = make(map[string][]Color)
			}

			group := new(Group)
			group.Read(reader)
			groupName = group.Name

			break
		case "c002":
			groupName = ""
			break
		default:
			log.Println("invalid block type")
		}
	}

	return
}

//	read a file on a file system
func DecodeFile(file string, group bool) (ase ASE, err error) {
	//	open the file
	f, err := os.Open(testFile)
	if err != nil {
		return
	}

	//	decode the file
	return Decode(f, group)
}

type ASE struct {
	Signature string
	Version   [2]int16
	NumBlocks [1]int32
	Colors    []Color
	Groups    map[string][]Color
}

//	ASE Files start are signed with ASEF at the beginning of the file
//	let's make sure this is an ASE file
func (ase *ASE) readSignature(file io.Reader) (err error) {
	signature := make([]uint8, 4)

	if err = binary.Read(file, binary.BigEndian, &signature); err != nil {
		return
	}

	ase.Signature = string(signature[0:])
	if ase.Signature != "ASEF" {
		return ErrInvalidFile
	}

	return
}

//	ASE version. Should be 1.0
func (ase *ASE) readVersion(file io.Reader) error {
	return binary.Read(file, binary.BigEndian, &ase.Version)
}

//	Total number of blocks in the ASE file
func (ase *ASE) readNumBlock(file io.Reader) error {
	return binary.Read(file, binary.BigEndian, &ase.NumBlocks)
}
