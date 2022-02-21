package decoder

import (
	"bytes"
	"io"
	"testing"
)

func TestReadCueChunk_Success(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xa0, 0x0f, 0x00, 0x00, 0x64, 0x61, 0x74, 0x61, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa0, 0x0f, 0x00, 0x00}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if len(c) != 1 {
		t.Errorf("Error: did not find the correct number of cues. Found %d cues", len(c))
	}

	if c[0].ID != "\u0001" {
		t.Errorf("ID was not the correct value")
	}

	if c[0].Position != 4000 {
		t.Errorf("Error: did not get the correct position. Found position %d", c[0].Position)
	}

	if c[0].DataChunkID != "data" {
		t.Errorf("Error: did not get the correct data ID. Found data ID %s", c[0].DataChunkID)
	}

	if c[0].ChunkStart != 0 {
		t.Errorf("Error: did not get the correct chunk start. Found chunk start %d", c[0].ChunkStart)
	}

	if c[0].BlockStart != 0 {
		t.Errorf("Error: did not get the correct block start. Found block start %d", c[0].BlockStart)
	}

	if c[0].SampleOffset != 4000 {
		t.Errorf("Error: did not get the correct sample offset. Found sample offset %d", c[0].SampleOffset)
	}
}

func TestReadCueChunk_Success_No_Cue_Points(t *testing.T) {
	b := []byte{0x00, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	c, err := readCueChunk(r)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if len(c) != 0 {
		t.Errorf("Expected cue point length to be %d. Received %d", 0, len(c))
	}

}

func TestReadCueChunk_Fail_No_Cue_Count(t *testing.T) {
	b := []byte{}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.EOF {
		t.Error("Returned error is not EOF")
	}
}

func TestReadCueChunk_Fail_Short_Cue_Count(t *testing.T) {
	b := []byte{0x01}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Error("Returned error is not Unexpected EOF")
	}

}

func TestReadCueChunk_Fail_No_Cue_ID(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.EOF {
		t.Error("Returned error is not EOF")
	}

}

func TestReadCueChunk_Fail_Short_Cue_ID(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x00}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Error("Returned error is not Unexpected EOF")
	}

}

func TestReadCueChunk_Fail_No_Cue_Position(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.EOF {
		t.Error("Returned error is not EOF")
	}

}

func TestReadCueChunk_Fail_Short_Cue_Position(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xa0}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Error("Returned error is not Unexpected EOF")
	}

}

func TestReadCueChunk_Fail_No_Cue_Data_Chunk_ID(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xa0, 0x0f, 0x00, 0x00}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.EOF {
		t.Error("Returned error is not EOF")
	}

}

func TestReadCueChunk_Fail_Short_Cue_Data_Chunk_ID(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xa0, 0x0f, 0x00, 0x00, 0x64}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Error("Returned error is not Unexpected EOF")
	}

}

func TestReadCueChunk_Fail_No_Cue_Chunk_Start(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xa0, 0x0f, 0x00, 0x00, 0x64, 0x61, 0x74, 0x61}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.EOF {
		t.Error("Returned error is not EOF")
	}

}

func TestReadCueChunk_Fail_Short_Cue_Chunk_Start(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xa0, 0x0f, 0x00, 0x00, 0x64, 0x61, 0x74, 0x61, 0x00}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Error("Returned error is not Unexpected EOF")
	}

}

func TestReadCueChunk_Fail_No_Cue_Block_Start(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xa0, 0x0f, 0x00, 0x00, 0x64, 0x61, 0x74, 0x61, 0x00, 0x00, 0x00, 0x00}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.EOF {
		t.Error("Returned error is not EOF")
	}

}

func TestReadCueChunk_Fail_Short_Cue_Block_Start(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xa0, 0x0f, 0x00, 0x00, 0x64, 0x61, 0x74, 0x61, 0x00, 0x00, 0x00, 0x00, 0x00}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Error("Returned error is not Unexpected EOF")
	}

}

func TestReadCueChunk_Fail_No_Cue_Sample_Offset(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xa0, 0x0f, 0x00, 0x00, 0x64, 0x61, 0x74, 0x61, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.EOF {
		t.Error("Returned error is not EOF")
	}

}

func TestReadCueChunk_Fail_Short_Cue_Sample_Offset(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xa0, 0x0f, 0x00, 0x00, 0x64, 0x61, 0x74, 0x61, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa0}
	r := bytes.NewReader(b)

	c, err := readCueChunk(r)

	if c != nil {
		t.Error("Returned []*CuePoint is not nil")
	}

	if err == nil {
		t.Error("Returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Error("Returned error is not Unexpected EOF")
	}

}
