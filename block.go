package ase

import (
	"encoding/binary"
	"io"
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

// Decode an ASE block.
func (b *block) Read(r io.Reader) (err error) {
	if err = b.readType(r); err != nil {
		return
	}

	return b.readLength(r)
}

// Reads the block's type.
// Can either be a group's start block, a group's end block, or a color entry.
func (block *block) readType(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, &block.Type)
}

// Reads the block's length.
func (block *block) readLength(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, &block.Length)
}
