package ase

import (
	"encoding/binary"
	"errors"
	"io"
)

var (
	ErrInvalidBlockType = errors.New("ase: invalid block type")
)

type block struct {
	Type   uint16
	Length int32
}

const (
	groupStart = uint16(0xc001)
	groupEnd   = uint16(0xc002)
	colorEntry = uint16(0x0001)
)

func (b *block) Read(r io.Reader) (err error) {

	//	type
	if err = b.readType(r); err != nil {
		return
	}

	//	block length
	if err = b.readLength(r); err != nil {
		return
	}

	return
}

//	0xc001 ⇒ Group start
//	0xc002 ⇒ Group end
//	0x0001 ⇒ Color entry
func (block *block) readType(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, &block.Type)
}

func (block *block) readLength(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, &block.Length)
}
