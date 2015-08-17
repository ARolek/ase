package ase

import (
	"strings"
	"testing"
)

func TestFuzzCrashers(t *testing.T) {

	var crashers = []string{
		"ASEF\x00\x01000000\xc0\x010000\x00\x00",
		"ASEF00\x00\x000000\x00\x010000\x00\x00",
	}

	for _, f := range crashers {
		Decode(strings.NewReader(f))
	}
}
