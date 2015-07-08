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
	NumBlocks [1]int32
	Colors    []Color
	Groups    map[string][]Color
}


// Decode processes an encoded ASE file into an ASE struct.
// Setting the `group` option to:
// (true) nests colors in their appropriate group as named in the ASE file
// (false) pools group colors into the Colors struct.
func (ase *ASE) Decode(aseFile string, group bool) error {
	var err error
	file, err := ioutil.ReadFile(aseFile)
	if err != nil {
		return err
	}

	fileBuf := bytes.NewReader(file)

	err = ase.readSignature(fileBuf)
	if err != nil {
		return err
	}

	err = ase.readVersion(fileBuf)
	if err != nil {
		return err
	}

	err = ase.readNumBlock(fileBuf)
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

// readSignature ensures that the file being read has the appropriate ASE file
// signature of ASEF.
func (ase *ASE) readSignature(file *bytes.Reader) error {
	signature := make([]uint8, 4)

	err := binary.Read(file, binary.BigEndian, &signature)
	if err != nil {
		return err
	}

	expectedSignature := "ASEF"
	if string(signature[0:]) != expectedSignature {
		return errors.New("File not an ASE file: expected signature of " + expectedSignature)
	}

	return nil
}

// readVersion ensures that the file being read has the appropriate ASE version
// of 1.0.
func (ase *ASE) readVersion(file *bytes.Reader) error {
	version := make([]int16, 2)

	err := binary.Read(file, binary.BigEndian, &version)
	if err != nil {
		return err
	}

	for i, v := range version {
		expectedVersion := [2]int16{1, 0}
		if expectedVersion[i] != v {
			return errors.New("File not an ASE file")
		}
	}

	return nil
}

//	readNumBlock stores the total number of blocks in the ASE file.
func (ase *ASE) readNumBlock(file *bytes.Reader) error {
	err := binary.Read(file, binary.BigEndian, &ase.NumBlocks)
	if err != nil {
		return err
	}

	return nil
}
