package Encoder

import (
	"encoding/binary"
	"errors"
	"io"
)

type CueChunk struct {
	cps []*CuePoint
}

type CuePoint struct {
	ID           string
	Position     uint32
	DataChunkID  string
	ChunkStart   uint32
	BlockStart   uint32
	SampleOffset uint32
}

func (c CueChunk) WriteTo(w io.Writer) (int, error) {
	bytesWritten := 0
	cueCount := uint32(len(c.cps))

	b := bytesFromString(CUE_CHUNK_ID)
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromUINT32(uint32(24 * len(c.cps)))
	err = binary.Write(w, binary.LittleEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromUINT32(cueCount)
	if err := binary.Write(w, binary.LittleEndian, &b); err != nil {
		return 0, errors.New("An error occurred when reading the number of cues")
	}
	bytesWritten += len(b)

	if cueCount > 0 {

		for i := uint32(0); i < cueCount; i++ {

			b = bytesFromString(c.cps[i].ID)
			if err := binary.Write(w, binary.BigEndian, &b); err != nil {
				return bytesWritten, errors.New("An error occurred when reading the cue ID")
			}
			bytesWritten += len(b)

			b = bytesFromUINT32(c.cps[i].Position)
			if err := binary.Write(w, binary.LittleEndian, &b); err != nil {
				return bytesWritten, errors.New("An error occurred when reading the cue position")
			}
			bytesWritten += len(b)

			b = bytesFromString(c.cps[i].DataChunkID)
			if err := binary.Write(w, binary.BigEndian, &b); err != nil {
				return bytesWritten, errors.New("An error occurred when reading the cue data chunk ID")
			}
			bytesWritten += len(b)

			b = bytesFromUINT32(c.cps[i].ChunkStart)
			if err := binary.Write(w, binary.LittleEndian, &b); err != nil {
				return bytesWritten, errors.New("An error occurred when reading the cue chunk start")
			}
			bytesWritten += len(b)

			b = bytesFromUINT32(c.cps[i].BlockStart)
			if err := binary.Write(w, binary.LittleEndian, &b); err != nil {
				return bytesWritten, errors.New("An error occurred when reading the cue block start")
			}
			bytesWritten += len(b)

			b = bytesFromUINT32(c.cps[i].SampleOffset)
			if err := binary.Write(w, binary.LittleEndian, &b); err != nil {
				return bytesWritten, errors.New("An error occurred when reading the cue sample offset")
			}
			bytesWritten += len(b)
		}
	}

	return bytesWritten, nil
}
