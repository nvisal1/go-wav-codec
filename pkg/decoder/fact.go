package decoder

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type factChunk struct {
	NumberOfSamples uint32
}

func readFactChunk(r *bytes.Reader) (*factChunk, error) {
	f := &factChunk{}

	if err := binary.Read(r, binary.LittleEndian, &f.NumberOfSamples); err != nil {
		return nil, errors.New("An error occurred when reading the fact number of samples")
	}
	return f, nil
}
