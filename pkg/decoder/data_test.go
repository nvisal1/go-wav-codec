package decoder

import (
	"bytes"
	"io"
	"testing"
)

func TestReadDataChunk_Success_8_Bits_Per_Sample(t *testing.T) {
	b := []byte{0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	i, err := readDataChunk(8, 8, r)
	if err != nil {
		t.Error(err.Error())
	}

	if len(i) != 8 {
		t.Errorf("expected buffer length 16, received buffer of length %d", len(i))
	}
}

func TestReadDataChunk_Success_16_Bits_Per_Sample(t *testing.T) {
	b := []byte{0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	i, err := readDataChunk(8, 16, r)
	if err != nil {
		t.Error(err.Error())
	}

	if len(i) != 8 {
		t.Errorf("expected buffer length 16, received buffer of length %d", len(i))
	}
}

func TestReadDataChunk_Success_32_Bits_Per_Sample(t *testing.T) {
	b := []byte{0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	i, err := readDataChunk(5, 32, r)
	if err != nil {
		t.Error(err.Error())
	}

	if len(i) != 5 {
		t.Errorf("expected buffer length 16, received buffer of length %d", len(i))
	}
}

func TestReadDataChunk_Fail_size_0(t *testing.T) {
	b := []byte{0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	i, err := readDataChunk(0, 16, r)

	if i != nil {
		t.Error("returned buffer is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err.Error() != "numSamples must be 1 or greater" {
		t.Errorf("returned error is not \"numSamples must be 1 or greater\". received \"%s\"", err.Error())
	}

}

func TestReadDataChunk_Fail_Unsupported_Bits_Per_Sample(t *testing.T) {
	b := []byte{0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	i, err := readDataChunk(1, 14, r)

	if i != nil {
		t.Error("returned buffer is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err.Error() != "received unsupported bits per sample: 14 | supported bits per sample: [8, 16, 32]" {
		t.Errorf("returned error is not \"received unsupported bits per sample: 14 | supported bits per sample: [8, 16, 32]\". received \"%s\"", err.Error())
	}
}

func TestReadDataChunk_Success_EOF_Before_Size_Limit(t *testing.T) {
	b := []byte{0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	i, err := readDataChunk(16, 16, r)

	if i == nil {
		t.Error("returned buffer is nil")
	}

	if len(i) != 16 {
		t.Errorf("returned buffer is not length 16. received buffer with length %d", len(i))
	}

	if i[10] != 20 {
		t.Errorf("expected value 20 at index 10 in returned buffer. received value %d", i[10])
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("returned error is not \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadDataChunk_Fail_Unexpected_EOF_Before_Size_Limit(t *testing.T) {
	b := []byte{0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00,
		0x00}

	r := bytes.NewReader(b)

	i, err := readDataChunk(16, 16, r)

	if i != nil {
		t.Error("returned buffer is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("returned error is not \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}
