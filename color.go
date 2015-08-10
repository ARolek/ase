package ase

import (
	"encoding/binary"
	"errors"
	"io"
	"strings"
	"unicode/utf8"
	"unicode/utf16"
)

var (
	ErrInvalidColorType = errors.New("ase: invalid color type")
	ErrInvalidColorValue = errors.New("ase: invalid color value")
)

type Color struct {
	nameLen uint16
	Name    string
	Model   string // CMYK, RGB, LAB or Gray
	Values  []float32
	Type    string //	Global, Spot, Normal
}

func (color *Color) read(r io.Reader) (err error) {

	if err = color.readNameLen(r); err != nil {
		return
	}

	if err = color.readName(r); err != nil {
		return
	}

	if err = color.readColorModel(r); err != nil {
		return
	}

	if err = color.readColorValues(r); err != nil {
		return
	}

	return color.readColorType(r)
}

func (color *Color) write(w io.Writer) (err error) {

	if err = color.writeBlockHeader(w); err != nil {
		return
	}

	if err = color.writeNameLen(w); err != nil {
		return
	}

	if err = color.writeName(w); err != nil {
		return
	}

	if err = color.writeModel(w); err != nil {
		return
	}

	if err = color.writeValues(w); err != nil {
		return
	}

	if err = color.writeType(w); err != nil {
		return
	}

	return
}

func (color *Color) readNameLen(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, &color.nameLen)
}

// Write color's nameLen
func (color *Color) writeNameLen(w io.Writer) (err error) {
	return binary.Write(w, binary.BigEndian, color.NameLen())
}


func (color *Color) readName(r io.Reader) (err error) {
	//	make array for our color name based on block length
	name := make([]uint16, color.nameLen)
	if err = binary.Read(r, binary.BigEndian, &name); err != nil {
		return
	}

	//	decode our name. we trim off the last byte since it's zero terminated
	// utf16.Decode returns a slice of runes from an input of []uint16
	color.Name = string(utf16.Decode(name[:len(name)-1]))

	return
}

// Encode the color's name as a slice of uint16.
func (color *Color) writeName(w io.Writer) (err error) {
	name := utf16.Encode([]rune(color.Name))
	name = append(name, uint16(0))
	return binary.Write(w, binary.BigEndian, name)
}

func (color *Color) readColorModel(r io.Reader) (err error) {
	colorModel := make([]uint8, 4)
	if err = binary.Read(r, binary.BigEndian, colorModel); err != nil {
		return
	}

	color.Model = strings.TrimSpace(string(colorModel[0:]))

	return
}

// Encode the color's model as a of slice of uint8.
func (color *Color) writeModel(w io.Writer) (err error) {
	model := utf8.Encode([]rune(color.Model))
	return binary.Write(w, binary.BigEndian, model)
}

func (color *Color) readColorValues(r io.Reader) (err error) {
	switch color.Model {
	case "RGB":
		rgb := make([]float32, 3)

		//	read into rbg array
		if err = binary.Read(r, binary.BigEndian, &rgb); err != nil {
			return
		}
		color.Values = rgb
		break
	case "LAB":
		lab := make([]float32, 3)

		//	read into lab array
		if err = binary.Read(r, binary.BigEndian, &lab); err != nil {
			return
		}

		color.Values = lab
		break
	case "CMYK":
		cmyk := make([]float32, 4)

		//	read into cmyk array
		if err = binary.Read(r, binary.BigEndian, &cmyk); err != nil {
			return
		}

		color.Values = cmyk
		break
	case "Gray":
		gray := make([]float32, 1)

		//	read into gray array
		if err = binary.Read(r, binary.BigEndian, &gray); err != nil {
			return
		}

		color.Values = gray
		break
	default:
		return ErrInvalidColorValue
	}

	return
}

// TODO: Write color's values
func (color *Color) writeValues(w io.Writer) (err error) {
	return binary.Write(w, binary.BigEndian, color.Values)
}


func (color *Color) readColorType(r io.Reader) (err error) {

	colorType := make([]int16, 1)

	//	read into colorType array
	if err = binary.Read(r, binary.BigEndian, colorType); err != nil {
		return
	}

	switch colorType[0] {
	case 0:
		color.Type = "Global"
		break
	case 1:
		color.Type = "Spot"
		break
	case 2:
		color.Type = "Normal"
		break
	default:
		return ErrInvalidColorType
	}

	return
}

// Encode the color's type.
func (color *Color) writeType(w io.Writer) (err error) {

	var colorType int16

	switch color.Type {
	case "Global":
		colorType = 0
		break
	case "Spot":
		colorType = 1
		break
	case "Normal":
		colorType = 2
		break
	default:
		return ErrInvalidColorType
	}

	return binary.Write(w, binary.BigEndian, []int16{colorType})
}

func (color *Color) NameLen() uint16 {
	return uint16(len(color.Name))
}

// Write color's block header as a part of encoding
func (color *Color) writeBlockHeader(w io.Writer) (err error) {
	colorEntry := uint16(0x0001)
	return binary.Write(w, binary.BigEndian, colorEntry)
}
