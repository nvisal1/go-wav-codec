package encoder

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestWriteDataChunkBuffer_Success_8(t *testing.T) {
	p := []int{0, 0, 0, 0}
	numChannels := uint16(2)
	bitsPerSample := uint16(8)

	var b bytes.Buffer

	_, _, err := writeDataChunkBuffer(&b, p, numChannels, bitsPerSample)
	if err != nil {
		t.Error(err.Error())
	}

	if len(b.Bytes()) != 4 {
		t.Error("buffer length is incorrect")
	}
}

func TestWriteDataChunkBuffer_Success_16(t *testing.T) {
	p := []int{0, 0, 0, 0}
	numChannels := uint16(2)
	bitsPerSample := uint16(16)

	var b bytes.Buffer

	_, _, err := writeDataChunkBuffer(&b, p, numChannels, bitsPerSample)
	if err != nil {
		t.Error(err.Error())
	}

	if len(b.Bytes()) != 8 {
		t.Error("buffer length is incorrect")
	}
}

func TestWriteDataChunkBuffer_Success_32(t *testing.T) {
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

func TestWriteDataChunkBuffer_Fail_Unsupported(t *testing.T) {
	p := []int{0, 0, 0, 0}
	numChannels := uint16(2)
	bitsPerSample := uint16(13)

	var b bytes.Buffer

	_, _, err := writeDataChunkBuffer(&b, p, numChannels, bitsPerSample)
	if err == nil {
		t.Error("returned error is nil")
	}

	if err.Error() != "unsupported bits per sample. expected one of [8, 16, 32]. received 13" {
		t.Errorf("expected \"unsupported bits per sample. expected one of [8, 16, 32]. received 13\". received \"%s\"", err.Error())
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
