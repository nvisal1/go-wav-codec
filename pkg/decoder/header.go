package decoder

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type wavFileHeader struct {
	FileSize uint32
}

func readWavFileHeader(r *bytes.Reader) (*wavFileHeader, error) {
	c, err := newChunk(r)

	if c.ID != riffChunkID {
		return nil, errors.New("File descriptor is not RIFF")
	}

	if c.Size <= 0 {
		return nil, errors.New("file size is less than or equal to 0")
	}

	rt := make([]byte, 4)
	err = binary.Read(r, binary.BigEndian, &rt)
	if err != nil {
		return nil, err
	}
	if string(rt[:]) != waveFileFormat {
		return nil, errors.New("File format is not WAVE. Actual " + string(rt[:]))
	}

	wfh := &wavFileHeader{
		FileSize: c.Size,
	}

	return wfh, nil
}
