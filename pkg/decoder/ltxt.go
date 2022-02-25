package decoder

import (
	"bytes"
	"encoding/binary"
)

type labeledText struct {
	CuePointID   string
	SampleLength uint32
	PurposeID    string
	Country      string
	Language     string
	Dialect      string
	CodePage     string
	Data         string
}

func readLTXTChunk(r *bytes.Reader, size uint32) (*labeledText, error) {

	l := &labeledText{}
	B32 := make([]byte, 4)
	B16 := make([]byte, 2)
	if err := binary.Read(r, binary.BigEndian, &B32); err != nil {
		return nil, err
	}
	l.CuePointID = string(B32[:])
	if err := binary.Read(r, binary.LittleEndian, &l.SampleLength); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &B32); err != nil {
		return nil, err
	}
	l.PurposeID = string(B32[:])
	if err := binary.Read(r, binary.BigEndian, &B16); err != nil {
		return nil, err
	}
	l.Country = string(B16[:])
	if err := binary.Read(r, binary.BigEndian, &B16); err != nil {
		return nil, err
	}
	l.Language = string(B16[:])
	if err := binary.Read(r, binary.BigEndian, &B16); err != nil {
		return nil, err
	}
	l.Dialect = string(B16[:])
	if err := binary.Read(r, binary.BigEndian, &B16); err != nil {
		return nil, err
	}
	l.CodePage = string(B16[:])

	// subtract 20 from the total size to account for the fields that come before
	// the data string
	b := make([]byte, size-20)
	if err := binary.Read(r, binary.BigEndian, &b); err != nil {
		return nil, err
	}
	l.Data = removeNullCharacters(string(b[:]))

	return l, nil
}
