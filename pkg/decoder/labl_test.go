package decoder

import (
	"bytes"
	"testing"
)

func TestReadLABLChunk(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x4c, 0x6f, 0x6f,
		0x70, 0x20, 0x31, 0x00, 0x00}

	r := bytes.NewReader(b)

	l, err := readLABLChunk(r, uint32(len(b)))
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if l.Data != "Loop 1" {
		t.Errorf("Error: did not find the correct data. Found %s", l.Data)
	}
}
