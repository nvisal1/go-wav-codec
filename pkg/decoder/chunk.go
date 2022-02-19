package Decoder

import (
	"bytes"
	"encoding/binary"
)

type Chunk struct {
	ID   string
	Size uint32
}

func NewChunk(r *bytes.Reader) (*Chunk, error) {
	var (
		id   [4]byte
		size [4]byte
	)

	c := &Chunk{}

	err := binary.Read(r, binary.BigEndian, &id)
	if err != nil {
		return nil, err
	}

	c.ID = string(id[:])

	err = binary.Read(r, binary.LittleEndian, &size)
	if err != nil {
		return nil, err
	}

	c.Size = binary.LittleEndian.Uint32(size[:])

	return c, nil
}
