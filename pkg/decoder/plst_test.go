package decoder

import (
	"bytes"
	"io"
	"testing"
)

func TestReadPlstChunk_Success(t *testing.T) {

	b := []byte{
		0x01, 0x00, 0x00, 0x00,
		0x49, 0x44, 0x20, 0x20,
		0x01, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	s, err := readPlstChunk(r)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if len(s) != 1 {
		t.Errorf("Error: did not find the correct number of segments. Found %d", len(s))
	}

	if s[0].CuePointID != "ID  " {
		t.Errorf("Error: did not find the correct cue point ID. Found %s", s[0].CuePointID)
	}

	if s[0].Length != 1 {
		t.Errorf("Error: did not find the correct length. Found %d", s[0].Length)
	}

	if s[0].NumberOfRepeats != 1 {
		t.Errorf("Error: did not find the correct number of repeats. Found %d", s[0].NumberOfRepeats)
	}

}

func TestReadPlstChunk_Fail_No_NumSegments(t *testing.T) {
	b := []byte{}

	r := bytes.NewReader(b)

	s, err := readPlstChunk(r)

	if s != nil {
		t.Error("returned plst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadPlstChunk_Fail_Short_NumSegments(t *testing.T) {
	b := []byte{0x01}

	r := bytes.NewReader(b)

	s, err := readPlstChunk(r)

	if s != nil {
		t.Error("returned plst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadPlstChunk_Success_NumSegments_Is_0(t *testing.T) {
	b := []byte{0x00, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	s, err := readPlstChunk(r)

	if err != nil {
		t.Error(err.Error())
	}

	if s != nil {
		t.Error("returned segments is not nil")
	}

	if len(s) != 0 {
		t.Errorf("expected 0 segments. received %d segments", len(s))
	}
}

func TestReadPlstChunk_Fail_No_Segment_Cue_Point_ID(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	s, err := readPlstChunk(r)

	if s != nil {
		t.Error("returned plst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadPlstChunk_Fail_Short_Segment_Cue_Point_ID(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00,
		0x49}

	r := bytes.NewReader(b)

	s, err := readPlstChunk(r)

	if s != nil {
		t.Error("returned plst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadPlstChunk_Fail_No_Segment_Length(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00,
		0x49, 0x44, 0x20, 0x20}

	r := bytes.NewReader(b)

	s, err := readPlstChunk(r)

	if s != nil {
		t.Error("returned plst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadPlstChunk_Fail_Short_Segment_Length(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00,
		0x49, 0x44, 0x20, 0x20,
		0x01}

	r := bytes.NewReader(b)

	s, err := readPlstChunk(r)

	if s != nil {
		t.Error("returned plst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadPlstChunk_Fail_No_Segment_Position(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00,
		0x49, 0x44, 0x20, 0x20,
		0x01, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	s, err := readPlstChunk(r)

	if s != nil {
		t.Error("returned plst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadPlstChunk_Fail_Short_Segment_Position(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00,
		0x49, 0x44, 0x20, 0x20,
		0x01, 0x00, 0x00, 0x00,
		0x01}

	r := bytes.NewReader(b)

	s, err := readPlstChunk(r)

	if s != nil {
		t.Error("returned plst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}
