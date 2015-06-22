package ase

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

//	ASE File Spec http://www.selapa.net/swatches/colors/fileformats.php#adobe_ase

type ASE struct {
	Signature string
	Version   [2]int16
	NumBlocks [1]int32
	Colors    []Color
	Groups    map[string][]Color
}

//	pass in the path to the ASE file to read and wheater or not to group
//	if group is set to false, colors in groups will be listed the same way
//	as colors outside of groups.
//	if group is set to true, colors will be nested in a group as named in the
//	ASE file
func (ase *ASE) Decode(aseFile string, group bool) error {
	var err error
	file, err := ioutil.ReadFile(aseFile)
	if err != nil {
		return err
	}

	fileBuf := bytes.NewReader(file)

	err = ase.ReadSignature(fileBuf)
	if err != nil {
		return err
	}

	err = ase.ReadVersion(fileBuf)
	if err != nil {
		return err
	}

	err = ase.ReadNumBlock(fileBuf)
	if err != nil {
		return err
	}

	groupName := ""
	//	itereate based on our block count
	for i := 0; i < int(ase.NumBlocks[0]); i++ {
		block := new(Block)
		block.Read(fileBuf)
		switch fmt.Sprintf("%x", block.Type) {
		case "0001":
			color := new(Color)
			err = color.Read(fileBuf)
			if err != nil {
				return err
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
			group.Read(fileBuf)
			groupName = group.Name

			break
		case "c002":
			groupName = ""
			break
		default:
			log.Println("INVALID BLOCK TYPE")
		}
	}

	return nil
}

//	ASE Files start are signed with ASEF at the beginning of the file
//	let's make sure this is an ASE file
func (ase *ASE) ReadSignature(file *bytes.Reader) error {
	signature := make([]uint8, 4)

	err := binary.Read(file, binary.BigEndian, &signature)
	if err != nil {
		return err
	}

	ase.Signature = string(signature[0:])
	if ase.Signature != "ASEF" {
		return errors.New("File not an ASE file")
	}

	return nil
}

//	ASE version. Should be 1.0
func (ase *ASE) ReadVersion(file *bytes.Reader) error {
	err := binary.Read(file, binary.BigEndian, &ase.Version)
	if err != nil {
		return err
	}

	return nil
}

//	Total number of blocks in the ASE file
func (ase *ASE) ReadNumBlock(file *bytes.Reader) error {
	err := binary.Read(file, binary.BigEndian, &ase.NumBlocks)
	if err != nil {
		return err
	}

	return nil
}
