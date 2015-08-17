package ase

import (
	"bytes"
	"encoding/binary"
	"io"
	"unicode/utf16"
)

type Group struct {
	nameLen uint16
	Name    string
	Colors  []Color
}

// Decode an ASE group.
func (group *Group) read(r io.Reader) (err error) {
	if err = group.readNameLen(r); err != nil {
		return
	}

	return group.readName(r)
}

// Decode a group's name length.
func (group *Group) readNameLen(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, &group.nameLen)
}

// Decode a group's name.
func (group *Group) readName(r io.Reader) (err error) {
	if group.nameLen == 0 {
		return
	}

	//	make array for our color name based on block length
	name := make([]uint16, group.nameLen)
	if err = binary.Read(r, binary.BigEndian, &name); err != nil {
		return
	}

	//	decode our name. we trim off the last byte since it's zero terminated
	group.Name = string(utf16.Decode(name[:len(name)-1]))

	return
}

// Encode a group's block headers (starting and ending), metadata and colors.
func (group *Group) write(w io.Writer) (err error) {

	// Write group start headers (block entry, block length, nameLen, name)
	if err = group.writeBlockStart(w); err != nil {
		return
	}

	if err = group.writeBlockLength(w); err != nil {
		return
	}

	if err = group.writeNameLen(w); err != nil {
		return
	}
	if err = group.writeName(w); err != nil {
		return
	}

	// Encode the group's color data.
	for _, color := range group.Colors {
		if err = color.write(w); err != nil {
			return
		}
	}

	// Write the group's closing headers
	if err = group.writeBlockEnd(w); err != nil {
		return
	}

	return
}

// Wrapper around writing a group start header.
func (group *Group) writeBlockStart(w io.Writer) (err error) {
	return binary.Write(w, binary.BigEndian, groupStart)
}

// Wrapper around writing a group end header.
func (group *Group) writeBlockEnd(w io.Writer) (err error) {
	// First writes the groupEnd block followed by two terminating zeroes.
	if err = binary.Write(w, binary.BigEndian, groupEnd); err != nil {
		return
	}
	if err = binary.Write(w, binary.BigEndian, uint16(0x0000)); err != nil {
		return
	}
	return binary.Write(w, binary.BigEndian, uint16(0x0000))
}

// Encode the color's name length.
func (group *Group) writeNameLen(w io.Writer) (err error) {
	// Adding one to the name length accounts for the zero-terminated character.
	return binary.Write(w, binary.BigEndian, group.NameLen()+1)
}

// Encode the group's name.
func (group *Group) writeName(w io.Writer) (err error) {
	name := utf16.Encode([]rune(group.Name))
	name = append(name, uint16(0))
	return binary.Write(w, binary.BigEndian, name)
}

// Helper function that returns the length of a group's name.
func (group *Group) NameLen() uint16 {
	return uint16(len(group.Name))
}

// Write color's block length as a part of the ASE encoding.
func (group *Group) writeBlockLength(w io.Writer) (err error) {
	blockLength, err := group.calculateBlockLength()
	if err != nil {
		return
	}
	return binary.Write(w, binary.BigEndian, blockLength)
}

// Calculates the block length to be written based on the color's attributes.
func (group *Group) calculateBlockLength() (val int32, err error) {
	buf := new(bytes.Buffer)
	if err = group.writeNameLen(buf); err != nil {
		return
	}
	if err = group.writeName(buf); err != nil {
		return
	}

	val = int32(buf.Len())

	return
}
