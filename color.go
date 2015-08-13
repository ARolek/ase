package ase

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strings"
	"unicode/utf16"
)

var (
	ErrInvalidColorType  = errors.New("ase: invalid color type")
	ErrInvalidColorValue = errors.New("ase: invalid color value")
	ErrInvalidColorModel = errors.New("ase: invalid color model")
)

type Color struct {
	nameLen uint16
	Name    string
	Model   string // CMYK, RGB, LAB or Gray
	Values  []float32
	Type    string // Global, Spot, Normal
}

// Read decodes a color's fields from an io.Reader interface
// in the order specified by the ASE specification.
func (color *Color) read(r io.Reader) (err error) {

	if err = color.readNameLen(r); err != nil {
		return err
	}

	if err = color.readName(r); err != nil {
		return err
	}

	if err = color.readModel(r); err != nil {
		return err
	}

	if err = color.readValues(r); err != nil {
		return err
	}

	if err = color.readType(r); err != nil {
		return err
	}

	return nil
}

// Reads the color's name length.
func (color *Color) readNameLen(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, &color.nameLen)
}

// Reads the color's name.
func (color *Color) readName(r io.Reader) (err error) {
	//	make array for our color name based on block length
	name := make([]uint16, color.nameLen) // assumes the nameLen was already defined.
	if err = binary.Read(r, binary.BigEndian, &name); err != nil {
		return err
	}

	//	decode our name. we trim off the last byte since it's zero terminated
	// utf16.Decode returns a slice of runes from an input of []uint16
	color.Name = string(utf16.Decode(name[:len(name)-1]))

	return nil
}

// Reads the color's model.
func (color *Color) readModel(r io.Reader) (err error) {
	// make array for our color model, where four is the max possible
	// amount of characters (RGB, LAB, CMYK, Gray).
	colorModel := make([]uint8, 4)
	if err = binary.Read(r, binary.BigEndian, colorModel); err != nil {
		return err
	}

	// Assign the string version of the `colorModel`
	color.Model = strings.TrimSpace(string(colorModel[0:]))

	return nil
}

// Reads the color's values.
func (color *Color) readValues(r io.Reader) (err error) {
	switch color.Model {
	case "RGB":
		rgb := make([]float32, 3)

		//	read into rbg array
		if err = binary.Read(r, binary.BigEndian, &rgb); err != nil {
			return err
		}
		color.Values = rgb
		break
	case "LAB":
		lab := make([]float32, 3)

		//	read into lab array
		if err = binary.Read(r, binary.BigEndian, &lab); err != nil {
			return err
		}

		color.Values = lab
		break
	case "CMYK":
		cmyk := make([]float32, 4)

		//	read into cmyk array
		if err = binary.Read(r, binary.BigEndian, &cmyk); err != nil {
			return err
		}

		color.Values = cmyk
		break
	case "Gray":
		gray := make([]float32, 1)

		//	read into gray array
		if err = binary.Read(r, binary.BigEndian, &gray); err != nil {
			return err
		}

		color.Values = gray
		break
	default:
		return ErrInvalidColorValue
	}

	return nil
}

// Read the color's type.
func (color *Color) readType(r io.Reader) (err error) {

	colorType := make([]int16, 1)

	//	read into colorType array
	if err = binary.Read(r, binary.BigEndian, colorType); err != nil {
		return err
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

	return nil
}

// Encodes a color's attributes according to the ASE specification.
func (color *Color) write(w io.Writer) (err error) {

	// Write the block type
	if err = color.writeBlockType(w); err != nil {
		return err
	}

	// Write the block length
	if err = color.writeBlockLength(w); err != nil {
		return err
	}

	// Write the color data
	if err = color.writeNameLen(w); err != nil {
		return err
	}

	if err = color.writeName(w); err != nil {
		return err
	}

	if err = color.writeModel(w); err != nil {
		return err
	}

	if err = color.writeValues(w); err != nil {
		return err
	}

	if err = color.writeType(w); err != nil {
		return err
	}

	return nil
}

// Write color's block length as a part of the ASE encoding.
func (color *Color) writeBlockLength(w io.Writer) (err error) {
	blockLength, err := color.calculateBlockLength()
	if err != nil {
		return err
	}
	if err = binary.Write(w, binary.BigEndian, blockLength); err != nil {
		return err
	}
	return nil
}

// Calculates the block length to be written based on the color's attributes.
func (color *Color) calculateBlockLength() (blockLength int32, err error) {
	buf := new(bytes.Buffer)

	if err = color.writeNameLen(buf); err != nil {
		return 0, err
	}
	if err = color.writeName(buf); err != nil {
		return 0, err
	}
	if err = color.writeModel(buf); err != nil {
		return 0, err
	}
	if err = color.writeValues(buf); err != nil {
		return 0, err
	}
	if err = color.writeType(buf); err != nil {
		return 0, err
	}

	return int32(buf.Len()), nil
}

// Encode the color's name length.
func (color *Color) writeNameLen(w io.Writer) (err error) {
	// Adding one to the name length accounts for the zero-terminated character.
	return binary.Write(w, binary.BigEndian, color.NameLen()+1)
}

// Encode the color's name as a slice of uint16.
func (color *Color) writeName(w io.Writer) (err error) {
	name := utf16.Encode([]rune(color.Name))
	name = append(name, uint16(0))
	return binary.Write(w, binary.BigEndian, name)
}

// Encode the color's model as a of slice of uint8.
func (color *Color) writeModel(w io.Writer) (err error) {
	model := make([]uint8, 4)

	// Populate model with the uint8 version of the each character in the string.
	for i, _ := range model {
		colorModel := color.Model

		// Our iterator over model needs to match the number of elements in color.Model.
		// If color.Model has a length less than four, append an empty space.
		if len(color.Model) < 4 {
			for len(colorModel) < 4 {
				colorModel += " "
			}
		}
		model[i] = uint8([]rune(colorModel)[i])
	}

	return binary.Write(w, binary.BigEndian, model)
}

// Encode color's values.
func (color *Color) writeValues(w io.Writer) (err error) {
	return binary.Write(w, binary.BigEndian, color.Values)
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

// Helper function that returns the length of a color's name.
func (color *Color) NameLen() uint16 {
	return uint16(len(color.Name))
}

// Write color's block header as a part of the ASE encoding.
func (color *Color) writeBlockType(w io.Writer) (err error) {
	return binary.Write(w, binary.BigEndian, colorEntry)
}
