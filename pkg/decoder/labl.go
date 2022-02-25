package decoder

import (
	"bytes"
	"encoding/binary"
)

type label struct {
	CuePointID string
	Data       string
}

func readLABLChunk(r *bytes.Reader, size uint32) (*label, error) {
	l := &label{}

	B32 := make([]byte, 4)

	if err := binary.Read(r, binary.BigEndian, &B32); err != nil {
		return nil, err
	}
	l.CuePointID = removeNullCharacters(string(B32[:]))

	// subtract 4 from the total size to account for the cue point ID
	b := make([]byte, size-4)
	if err := binary.Read(r, binary.BigEndian, &b); err != nil {
		return nil, err
	}
	l.Data = removeNullCharacters(string(b[:]))

	return l, nil
}
