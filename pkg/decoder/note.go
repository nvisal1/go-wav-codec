package decoder

import (
	"bytes"
	"encoding/binary"
)

type note struct {
	CuePointID string
	Data       string
}

func readNoteChunk(r *bytes.Reader, size uint32) (*note, error) {
	n := &note{}
	B32 := make([]byte, 4)

	if err := binary.Read(r, binary.BigEndian, &B32); err != nil {
		return nil, err
	}
	n.CuePointID = removeNullCharacters(string(B32[:]))

	// subtract 4 from the total size to account for the cue point ID
	b := make([]byte, size-4)
	if err := binary.Read(r, binary.BigEndian, &b); err != nil {
		return nil, err
	}
	n.Data = removeNullCharacters(string(b[:]))

	return n, nil
}
