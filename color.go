package ase

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"strings"
	"unicode/utf16"
)

type Color struct {
	NameLen uint16
	Name    string
	Model   string
	Values  []float32
	Type    string
}

func (color *Color) Read(file *bytes.Reader) error {
	var err error
	err = color.ReadNameLen(file)
	if err != nil {
		return err
	}

	err = color.ReadName(file)
	if err != nil {
		return err
	}

	err = color.ReadColorModel(file)
	if err != nil {
		return err
	}

	err = color.ReadColorValues(file)
	if err != nil {
		return err
	}

	err = color.ReadColorType(file)
	if err != nil {
		return err
	}

	return nil
}

func (color *Color) ReadNameLen(file *bytes.Reader) error {
	err := binary.Read(file, binary.BigEndian, &color.NameLen)
	if err != nil {
		return err
	}

	return nil
}

func (color *Color) ReadName(file *bytes.Reader) error {
	//	make array for our color name based on block length
	name := make([]uint16, color.NameLen)
	err := binary.Read(file, binary.BigEndian, &name)
	if err != nil {
		return err
	}

	//	decode our name. we trim off the last byte since it's zero terminated
	color.Name = string(utf16.Decode(name[:len(name)-1]))

	return nil
}

func (color *Color) ReadColorModel(file *bytes.Reader) error {
	colorModel := make([]uint8, 4)
	err := binary.Read(file, binary.BigEndian, colorModel)
	if err != nil {
		return err
	}

	color.Model = strings.TrimSpace(string(colorModel[0:]))
	//log.Println(color.Model)

	return nil
}

func (color *Color) ReadColorValues(file *bytes.Reader) error {
	var err error
	switch color.Model {
	case "RGB":
		log.Println("RGB")
		rgb := make([]float32, 3)
		err = binary.Read(file, binary.BigEndian, &rgb)
		if err != nil {
			return err
		}
		color.Values = rgb
		log.Println(color.Values)
		break
	case "LAB":
		log.Println("LAB")
		lab := make([]float32, 3)
		err = binary.Read(file, binary.BigEndian, &lab)
		if err != nil {
			return err
		}
		color.Values = lab
		log.Println(lab)
		break
	case "CMYK":
		log.Println("CMYK")
		cmyk := make([]float32, 4)
		err = binary.Read(file, binary.BigEndian, &cmyk)
		if err != nil {
			return err
		}
		color.Values = cmyk
		log.Println(cmyk)
		break
	case "Gray":
		log.Println("Gray")
		gray := make([]float32, 1)
		err = binary.Read(file, binary.BigEndian, &gray)
		if err != nil {
			return err
		}
		color.Values = gray
		log.Println(gray)
		break
	}

	return nil
}

func (color *Color) ReadColorType(file *bytes.Reader) error {
	var err error
	colorType := make([]int16, 1)
	err = binary.Read(file, binary.BigEndian, colorType)
	if err != nil {
		return err
	}
	//	log.Println(color.Type)

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
		return errors.New("invalid color type")
	}

	return nil
}
