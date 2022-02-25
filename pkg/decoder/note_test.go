package decoder

import (
	"bytes"
	"io"
	"testing"
)

func TestReadNoteChunk_Success(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x53, 0x61, 0x6d,
		0x70, 0x6c, 0x65, 0x72, 0x20, 0x6c, 0x6f, 0x6f,
		0x70, 0x00, 0x00}

	r := bytes.NewReader(b)

	n, err := readNoteChunk(r, uint32(len(b)))
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if n.Data != "Sampler loop" {
		t.Errorf("Error: did not find the correct data. Found %s", n.Data)
	}
}

func TestReadNoteChunk_Fail_No_Cue_Point_ID(t *testing.T) {
	b := []byte{}

	r := bytes.NewReader(b)

	l, err := readNoteChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned labl chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadNoteChunk_Fail_Short_Cue_Point_ID(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00}

	r := bytes.NewReader(b)

	l, err := readNoteChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned labl chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadNoteChunk_Fail_No_Data(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	l, err := readNoteChunk(r, uint32(12))

	if l != nil {
		t.Error("returned labl chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadNoteChunk_Fail_Short_Data(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	l, err := readNoteChunk(r, uint32(12))

	if l != nil {
		t.Error("returned labl chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}
