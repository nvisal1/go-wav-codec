package decoder

import (
	"bytes"
	"encoding/binary"
	"strings"
)

type cuePoint struct {
	ID           string
	Position     uint32
	DataChunkID  string
	ChunkStart   uint32
	BlockStart   uint32
	SampleOffset uint32
}

func readCueChunk(r *bytes.Reader) ([]*cuePoint, error) {
	var cueCount uint32

	if err := binary.Read(r, binary.LittleEndian, &cueCount); err != nil {
		return nil, err
	}

	cues := make([]*cuePoint, 0, cueCount)

	if cueCount > 0 {
		str := make([]byte, 4)

		for i := uint32(0); i < cueCount; i++ {
			c := &cuePoint{}

			if err := binary.Read(r, binary.BigEndian, &str); err != nil {
				return nil, err
			}

			c.ID = strings.ReplaceAll(string(str[:]), "\u0000", "")

			if err := binary.Read(r, binary.LittleEndian, &c.Position); err != nil {
				return nil, err
			}

			if err := binary.Read(r, binary.BigEndian, &str); err != nil {
				return nil, err
			}

			c.DataChunkID = string(str[:])

			if err := binary.Read(r, binary.LittleEndian, &c.ChunkStart); err != nil {
				return nil, err
			}

			if err := binary.Read(r, binary.LittleEndian, &c.BlockStart); err != nil {
				return nil, err
			}

			if err := binary.Read(r, binary.LittleEndian, &c.SampleOffset); err != nil {
				return nil, err
			}

			cues = append(cues, c)
		}
	}

	return cues, nil
}
