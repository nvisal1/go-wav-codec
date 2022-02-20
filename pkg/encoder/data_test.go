package encoder

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestWriteDataChunkBuffer(t *testing.T) {
	p := []int{0, 0, 0, 0}
	numChannels := uint16(2)
	bitsPerSample := uint16(32)

	var b bytes.Buffer

	_, _, err := writeDataChunkBuffer(&b, p, numChannels, bitsPerSample)
	if err != nil {
		t.Error(err.Error())
	}

	if len(b.Bytes()) != 16 {
		t.Error("buffer length is incorrect")
	}
}

func TestWriteDataChunkID(t *testing.T) {
	var b bytes.Buffer

	_, err := writeDataChunkID(&b)
	if err != nil {
		t.Error(err.Error())
	}

	temp := make([]byte, 4)
	_, err = b.Read(temp)
	if err != nil {
		t.Error(err.Error())
	}

	if string(temp[:]) != dataChunkID {
		t.Errorf("first 4 bytes are not %s", dataChunkID)
	}

	_, err = b.Read(temp)
	if err != nil {
		t.Error(err.Error())
	}

	if binary.LittleEndian.Uint32(temp) != 0 {
		t.Error("size is not 0")
	}
}
