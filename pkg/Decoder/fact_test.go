package Decoder

import (
	"bytes"
	"testing"
)

func TestReadFactChunk(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	f, err := ReadFactChunk(r)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if f.NumberOfSamples != 1 {
		t.Errorf("Error: did not find the correct number of samples. Found %d", f.NumberOfSamples)
	}
}
