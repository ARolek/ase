package ase

import (
	"bytes"
	"encoding/binary"
)

type Block struct {
	Type   [2]uint8
	Length [1]int32
}

func (block *Block) Read(file *bytes.Reader) error {
	var err error
	err = block.ReadType(file)
	if err != nil {
		return err
	}

	err = block.ReadLength(file)
	if err != nil {
		return err
	}

	return nil
}

//	0xc001 ⇒ Group start
//	0xc002 ⇒ Group end
//	0x0001 ⇒ Color entry
func (block *Block) ReadType(file *bytes.Reader) error {
	err := binary.Read(file, binary.BigEndian, &block.Type)
	if err != nil {
		return err
	}

	return nil
}

func (block *Block) ReadLength(file *bytes.Reader) error {
	err := binary.Read(file, binary.BigEndian, &block.Length)
	if err != nil {
		return err
	}

	return nil
}
