package ase

import "testing"

func TestGroupStruct(t *testing.T) {
	g := new(Group)

	var defaultNameLen uint16
	var defaultName string

	if g.NameLen != defaultNameLen {
		t.Error("expected block to have default Name Len of", defaultNameLen)
	}

	if g.Name != defaultName {
		t.Error("expected block to have default Name of", defaultName)
	}

	if len(g.Colors) != 0 {
		t.Error("expected Colors to be initially empty")
	}
}

func TestGroupRead(t *testing.T) {
	t.SkipNow()
}
