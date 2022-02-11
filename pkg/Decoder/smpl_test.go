package Decoder

import (
	"bytes"
	"testing"
)

func TestReadSmplChunk(t *testing.T) {
	b := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x93, 0x58, 0x00, 0x00, 0x2d, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x10, 0xb1, 0x02, 0x00, 0x1f, 0x62, 0x05, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	s, err := ReadSmplChunk(r)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if len(s.Loops) != 1 {
		t.Errorf("Error: did not find the correct number of loops. Found %d", len(s.Loops))
	}

}
