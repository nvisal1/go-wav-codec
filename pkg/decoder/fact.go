package Decoder

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type FactChunk struct {
	NumberOfSamples uint32
}

func ReadFactChunk(r *bytes.Reader) (*FactChunk, error) {
	f := &FactChunk{}

	if err := binary.Read(r, binary.LittleEndian, &f.NumberOfSamples); err != nil {
		return nil, errors.New("An error occurred when reading the fact number of samples")
	}
	return f, nil
}
