package Encoder

import (
	"testing"
)

func TestBytesFromString(t *testing.T) {
	s := "this is a string"
	b := bytesFromString(s)
	if len(b) != 16 {
		t.Error("returned value is not an array")
	}
	if b[0] != 116 {
		t.Error("returned value does not have elements of type byte")
	}
}

func TestBytesFromUINT16(t *testing.T) {
	b16 := uint16(2)
	b := bytesFromUINT16(b16)
	if len(b) != 2 {
		t.Error("returned value is not an array")
	}
	if b[0] != 2 && b[1] != 0 {
		t.Error("returned value does not have elements of type byte")
	}
}

func TestBytesFromUINT32(t *testing.T) {
	b32 := uint32(2)
	b := bytesFromUINT32(b32)
	if len(b) != 4 {
		t.Error("returned value is not an array")
	}
	if b[0] != 2 && b[1] != 0 && b[2] != 0 && b[3] != 0 {
		t.Error("returned value does not have elements of type byte")
	}
}
