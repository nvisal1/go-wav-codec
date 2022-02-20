package decoder

import (
	"bytes"
	"encoding/binary"
	"io"
)

type label struct {
	CuePointID string
	Data       string
}

func readLABLChunk(r *bytes.Reader, size uint32) (*label, error) {
	l := &label{}
	bytesRead := 0

	B32 := make([]byte, 4)

	if err := binary.Read(r, binary.BigEndian, &B32); err != nil {
		return nil, err
	}
	bytesRead += 4
	l.CuePointID = removeNullCharacters(string(B32[:]))

	for {
		data := make([]byte, 1)
		if err := binary.Read(r, binary.BigEndian, &data); err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
			return nil, err
		}
		l.Data = l.Data + removeNullCharacters(string(data[:]))
		bytesRead++
		if uint32(bytesRead) == size {
			break
		}
	}

	return l, nil
}
