package ase

import "testing"

func TestColorStruct(t *testing.T) {
    c := new(Color)
    
	var defaultNameLen uint16
	var defaultName    string
	var defaultModel   string
	var defaultType    string
	
	if c.NameLen != defaultNameLen {
	    t.Error("expected default Name Length of", defaultNameLen,
	        " got ", c.NameLen)
	}
	
	if c.Name != defaultName {
	    t.Error("expected default Name of", defaultName,
	        " got ", c.Name)
	}
	
	if c.Model != defaultModel {
	    t.Error("expected default Model of", defaultModel,
	        " got ", c.Model)
	}
	
	if len(c.Values) != 0 {
	    t.Error("expected empty set of Values by default")   
	}
	
	if c.Type != defaultType {
	    t.Error("expected default Type of", defaultType,
	        " got ", c.Type)
	}
}

func TestColorRead(t *testing.T) {
	t.SkipNow()
}