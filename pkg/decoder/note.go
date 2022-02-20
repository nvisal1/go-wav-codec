package decoder

import (
	"bytes"
	"encoding/binary"
	"io"
)

type note struct {
	CuePointID string
	Data       string
}

func readNoteChunk(r *bytes.Reader, size uint32) (*note, error) {
	n := &note{}
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
		bytesRead++
		if uint32(bytesRead) == size {
			break
		}
	}

	return n, nil
}
