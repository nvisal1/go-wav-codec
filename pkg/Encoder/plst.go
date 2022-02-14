package Encoder

import (
	"encoding/binary"
	"errors"
	"io"
)

type PLSTChunk struct {
	plsts []*PlstSegment
}

type PlstSegment struct {
	CuePointID      string
	Length          uint32
	NumberOfRepeats uint32
}

func (p PLSTChunk) WriteTo(w io.Writer) (int, error) {
	bytesWritten := 0

	b := bytesFromString(PLST_CHUNK_ID)
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = Align(bytesFromUINT32(uint32(len(p.plsts) * 12)))
	err = binary.Write(w, binary.LittleEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	numSegments := uint32(len(p.plsts))
	b = bytesFromUINT32(numSegments)
	if err := binary.Write(w, binary.LittleEndian, &b); err != nil {
		return bytesWritten, errors.New("An error occurred when reading the number of PLST segments")
	}
	bytesWritten += len(b)

	if numSegments > 0 {

		for i := uint32(0); i < numSegments; i++ {

			b = bytesFromString(p.plsts[i].CuePointID)
			if err := binary.Write(w, binary.BigEndian, &b); err != nil {
				return bytesWritten, errors.New("An error occurred when reading the number of PLST segments")
			}
			bytesWritten += len(b)

			b = bytesFromUINT32(p.plsts[i].Length)
			if err := binary.Write(w, binary.LittleEndian, &b); err != nil {
				return bytesWritten, errors.New("An error occurred when reading the number of PLST segments")
			}
			bytesWritten += len(b)

			b = bytesFromUINT32(p.plsts[i].NumberOfRepeats)
			if err := binary.Write(w, binary.LittleEndian, &b); err != nil {
				return bytesWritten, errors.New("An error occurred when reading the number of PLST segments")
			}
			bytesWritten += len(b)

		}
	}
	return bytesWritten, nil
}
