package Decoder

import (
	"bytes"
	"testing"
)

func TestReadNoteChunk(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x53, 0x61, 0x6d,
		0x70, 0x6c, 0x65, 0x72, 0x20, 0x6c, 0x6f, 0x6f,
		0x70, 0x00, 0x00}

	r := bytes.NewReader(b)

	n, err := ReadNoteChunk(r)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if n.Data != "Sampler loop" {
		t.Errorf("Error: did not find the correct data. Found %s", n.Data)
	}
}
