package ase

import (
	"encoding/binary"
	"errors"
	"io"
	"strings"
	"unicode/utf16"
)

var (
	ErrInvalidColorType = errors.New("ase: invalid color type")
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

func (color *Color) readNameLen(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, &color.nameLen)
}

func (color *Color) readName(r io.Reader) (err error) {
	//	make array for our color name based on block length
	name := make([]uint16, color.nameLen)
	if err = binary.Read(r, binary.BigEndian, &name); err != nil {
		return
	}

	//	decode our name. we trim off the last byte since it's zero terminated
	color.Name = string(utf16.Decode(name[:len(name)-1]))

	return
}

func (color *Color) readColorModel(r io.Reader) (err error) {
	colorModel := make([]uint8, 4)
	if err = binary.Read(r, binary.BigEndian, colorModel); err != nil {
		return
	}

	color.Model = strings.TrimSpace(string(colorModel[0:]))

	return
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
	}

	return
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

func (color *Color) write(w io.Writer) (err error) {
	if err = color.writeColorNameLength(w); err != nil {
		return
	}

	if err = color.writeColorName(w); err != nil {
		return
	}

	if err = color.writeColorModel(w); err != nil {
		return
	}

	if err = color.writeColorValues(w); err != nil {
		return
	}

	if err = color.writeColorType(w); err != nil {
		return
	}

	return
}

func (color *Color) writeColorName(w io.Writer) error {
	colorNameSlice := []rune(color.Name)
	colorNameSlice = append(colorNameSlice, 0)
	colorName := utf16.Encode(colorNameSlice)
	return binary.Write(w, binary.BigEndian, colorName)
}

func (color *Color) writeColorModel(w io.Writer) error {
	colorSlice := []byte(color.Model)

	//If color model is either RGB or LAB append space (0x20)
	if color.Model == "RGB" || color.Model == "LAB" {
		colorSlice = append(colorSlice, 0x20)
	}

	return binary.Write(w,binary.BigEndian, colorSlice)
}

func (color *Color) writeColorValues(w io.Writer) error {
	var err error
	for _, cv := range color.Values {

		err = binary.Write(w,binary.BigEndian, cv)

		if(err != nil){
			return err
		}
	}
	return err
}

func (color *Color) writeColorType(w io.Writer) error {
	var cType int16
	switch color.Type {
		case "Global":
			cType = 0
			break
		case "Spot":
			cType = 1
			break
		case "Normal":
			cType = 2
			break
		default:
			return ErrInvalidColorType

	}
	return binary.Write(w, binary.BigEndian, cType)
}

func (color *Color) writeColorNameLength(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, color.nameLen)
}