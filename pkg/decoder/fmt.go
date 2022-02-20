package decoder

import (
	"bytes"
	"encoding/binary"
)

type fmtChunk struct {
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
}

func readFMTChunk(r *bytes.Reader) (*fmtChunk, error) {

	f := &fmtChunk{}

	if err := binary.Read(r, binary.LittleEndian, &f.AudioFormat); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &f.NumChannels); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &f.SampleRate); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &f.ByteRate); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &f.BlockAlign); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &f.BitsPerSample); err != nil {
		return nil, err
	}

	return f, nil
}
