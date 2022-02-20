package decoder

import (
	"bytes"
	"encoding/binary"
)

type chunk struct {
	ID   string
	Size uint32
}

func newChunk(r *bytes.Reader) (*chunk, error) {
	var (
		id   [4]byte
		size [4]byte
	)

	c := &chunk{}

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
