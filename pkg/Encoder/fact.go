package Encoder

import (
	"encoding/binary"
	"errors"
	"io"
)

type FactChunk struct {
	NumberOfSamples uint32
}

func (fc FactChunk) WriteTo(w io.Writer) (int, error) {
	bytesWritten := 0

	b := bytesFromString(FACT_CHUNK_ID)
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromUINT32(uint32(4))
	err = binary.Write(w, binary.LittleEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	if err := binary.Write(w, binary.LittleEndian, &fc.NumberOfSamples); err != nil {
		return bytesWritten, errors.New("An error occurred when reading the fact number of samples")
	}
	return bytesWritten, nil
}
