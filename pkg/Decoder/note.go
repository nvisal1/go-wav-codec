package Decoder

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Note struct {
	CuePointID string
	Data       string
}

func ReadNoteChunk(r *bytes.Reader, size uint32) (*Note, error) {
	n := &Note{}
	bytesRead := 0
	B32 := make([]byte, 4)

	if err := binary.Read(r, binary.BigEndian, &B32); err != nil {
		return nil, err
	}
	bytesRead += 4
	n.CuePointID = removeNullCharacters(string(B32[:]))

	for {
		data := make([]byte, 1)
		if err := binary.Read(r, binary.BigEndian, &data); err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
			return nil, err
		}
		n.Data = n.Data + removeNullCharacters(string(data[:]))
		bytesRead += 1
		if uint32(bytesRead) == size {
			break
		}
	}

	return n, nil
}
