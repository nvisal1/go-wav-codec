package Decoder

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type WavFileHeader struct {
	FileSize uint32
}

func ReadWavFileHeader(r *bytes.Reader) (*WavFileHeader, error) {
	c, err := NewChunk(r)

	if c.ID != RIFF_CHUNK_ID {
		return nil, errors.New("File descriptor is not RIFF")
	}

	if c.Size <= 0 {
		return nil, errors.New("File size is less than or equal to 0. Actual file size: " + string(c.Size))
	}

	rt := make([]byte, 4)
	err = binary.Read(r, binary.BigEndian, &rt)
	if err != nil {
		return nil, err
	}
	if string(rt[:]) != WAVE_FILE_FORMAT {
		return nil, errors.New("File format is not WAVE. Actual " + string(rt[:]))
	}

	wfh := &WavFileHeader{
		FileSize: c.Size,
	}

	return wfh, nil
}
