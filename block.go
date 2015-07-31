package ase

import (
	"encoding/binary"
	"io"
)

type Block struct {
	Type   [2]uint8
	Length [1]int32
}

func (block *Block) Read(file io.Reader) (err error) {

	//	type
	if err = block.readType(file); err != nil {
		return
	}

	//	block length
	if err = block.readLength(file); err != nil {
		return
	}

	return
}

//	0xc001 ⇒ Group start
//	0xc002 ⇒ Group end
//	0x0001 ⇒ Color entry
func (block *Block) readType(file io.Reader) error {
	return binary.Read(file, binary.BigEndian, &block.Type)
}

func (block *Block) readLength(file io.Reader) error {
	return binary.Read(file, binary.BigEndian, &block.Length)
}
