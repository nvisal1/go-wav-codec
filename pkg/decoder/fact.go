package decoder

import (
	"bytes"
	"encoding/binary"
)

type factChunk struct {
	NumberOfSamples uint32
}

func readFactChunk(r *bytes.Reader) (*factChunk, error) {
	f := &factChunk{}

	if err := binary.Read(r, binary.LittleEndian, &f.NumberOfSamples); err != nil {
		return nil, err
	}
	return f, nil
}
