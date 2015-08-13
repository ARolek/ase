package ase

import (
	"encoding/binary"
	"io"
	"unicode/utf16"
)

type Group struct {
	nameLen uint16
	Name    string
	Colors  []Color
}

func (group *Group) read(r io.Reader) (err error) {
	if err = group.readNameLen(r); err != nil {
		return
	}

	return group.readName(r)
}

func (group *Group) readNameLen(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, &group.nameLen)
}

func (group *Group) readName(r io.Reader) (err error) {
	//	make array for our color name based on block length
	name := make([]uint16, group.nameLen)
	if err = binary.Read(r, binary.BigEndian, &name); err != nil {
		return
	}

	//	decode our name. we trim off the last byte since it's zero terminated
	group.Name = string(utf16.Decode(name[:len(name)-1]))

	return
}

func (group *Group) write(w io.Writer) (err error){
	if err = group.writeName(w); err != nil {
		return
	}

	if err = group.writeNameLen(w); err != nil {
		return
	}

	return
}

func (group *Group) writeNameLen(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, group.nameLen)
}

func (group *Group) writeName(w io.Writer) error {
	groupNameSlice := []rune(group.Name)
	groupNameSlice = append(groupNameSlice, 0)
	groupName := utf16.Encode(groupNameSlice)
	return binary.Write(w, binary.BigEndian, groupName)
}
