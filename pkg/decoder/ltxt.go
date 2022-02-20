package decoder

import (
	"bytes"
	"encoding/binary"
	"io"
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
	bytesRead := 0
	B32 := make([]byte, 4)
	B16 := make([]byte, 2)
	if err := binary.Read(r, binary.BigEndian, &B32); err != nil {
		return nil, err
	}
	bytesRead += 4
	l.CuePointID = string(B32[:])
	if err := binary.Read(r, binary.LittleEndian, &l.SampleLength); err != nil {
		return nil, err
	}
	bytesRead += 4
	if err := binary.Read(r, binary.BigEndian, &B32); err != nil {
		return nil, err
	}
	bytesRead += 4
	l.PurposeID = string(B32[:])
	if err := binary.Read(r, binary.BigEndian, &B16); err != nil {
		return nil, err
	}
	bytesRead += 2
	l.Country = string(B16[:])
	if err := binary.Read(r, binary.BigEndian, &B16); err != nil {
		return nil, err
	}
	bytesRead += 2
	l.Language = string(B16[:])
	if err := binary.Read(r, binary.BigEndian, &B16); err != nil {
		return nil, err
	}
	bytesRead += 2
	l.Dialect = string(B16[:])
	if err := binary.Read(r, binary.BigEndian, &B16); err != nil {
		return nil, err
	}
	bytesRead += 2
	l.CodePage = string(B16[:])

	if uint32(bytesRead) == size {
		return l, nil
	}

	for {
		data := make([]byte, 1)
		if err := binary.Read(r, binary.BigEndian, &data); err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
			return nil, err
		}
		l.Data = l.Data + string(data[:])
		bytesRead++
		if uint32(bytesRead) == size {
			break
		}
	}

	return l, nil
}
