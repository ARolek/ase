package ase

import (
	"bytes"
	"encoding/binary"
	"unicode/utf16"
)

type Group struct {
	NameLen uint16
	Name    string
	Colors  []Color
}

func (group *Group) Read(file *bytes.Reader) error {
	var err error
	err = group.readNameLen(file)
	if err != nil {
		return err
	}

	err = group.readName(file)
	if err != nil {
		return err
	}

	return nil
}

func (group *Group) readNameLen(file *bytes.Reader) error {
	err := binary.Read(file, binary.BigEndian, &group.NameLen)
	if err != nil {
		return err
	}

	return nil
}

func (group *Group) readName(file *bytes.Reader) error {
	//	make array for our color name based on block length
	name := make([]uint16, group.NameLen)
	err := binary.Read(file, binary.BigEndian, &name)
	if err != nil {
		return err
	}

	//	decode our name. we trim off the last byte since it's zero terminated
	group.Name = string(utf16.Decode(name[:len(name)-1]))

	return nil
}
