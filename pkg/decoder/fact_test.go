package decoder

import (
	"bytes"
	"io"
	"testing"
)

func TestReadFactChunk_Success(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	f, err := readFactChunk(r)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if f.NumberOfSamples != 1 {
		t.Errorf("Error: did not find the correct number of samples. Found %d", f.NumberOfSamples)
	}
}

func TestReadFactChunk_Fail_EOF(t *testing.T) {
	b := []byte{}

	r := bytes.NewReader(b)

	f, err := readFactChunk(r)

	if f != nil {
		t.Error("returned fact chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("returned error is not \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadFactChunk_Fail_Unexpected_EOF(t *testing.T) {
	b := []byte{0x00}

	r := bytes.NewReader(b)

	f, err := readFactChunk(r)

	if f != nil {
		t.Error("returned fact chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("returned error is not \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}
