package decoder

import (
	"bytes"
	"io"
	"testing"
)

func TestReadInstChunk_Success(t *testing.T) {
	b := []byte{0x3c, 0x00, 0x00, 0x00, 0x7f, 0x01, 0x7f, 0x00}
	r := bytes.NewReader(b)

	i, err := readInstChunk(r)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if i.UnshiftedNote != 60 {
		t.Errorf("Error: did not find the correct unshifted note. Found %d", i.UnshiftedNote)
	}

	if i.FineTuneDB != 0 {
		t.Errorf("Error: did not find the correct fine tune db. Found %d", i.FineTuneDB)
	}

	if i.Gain != 0 {
		t.Errorf("Error: did not find the correct gain. Found %d", i.Gain)
	}

	if i.LowNote != 0 {
		t.Errorf("Error: did not find the correct low note. Found %d", i.LowNote)
	}

	if i.HighNote != 127 {
		t.Errorf("Error: did not find the correct high note. Found %d", i.HighNote)
	}

	if i.LowVelocity != 127 {
		t.Errorf("Error: did not find the correct low velocity. Found %d", i.LowVelocity)
	}

	if i.HighVelocity != 1 {
		t.Errorf("Error: did not find the correct high velocity. Found %d", i.HighVelocity)
	}

}

func TestReadInstChunk_Fail_No_UnshiftedNote(t *testing.T) {
	b := []byte{}
	r := bytes.NewReader(b)

	i, err := readInstChunk(r)

	if i != nil {
		t.Error("returned inst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadInstChunk_Fail_No_FineTuneDB(t *testing.T) {
	b := []byte{0x3c}
	r := bytes.NewReader(b)

	i, err := readInstChunk(r)

	if i != nil {
		t.Error("returned inst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadInstChunk_Fail_No_Gain(t *testing.T) {
	b := []byte{0x3c, 0x00}
	r := bytes.NewReader(b)

	i, err := readInstChunk(r)

	if i != nil {
		t.Error("returned inst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadInstChunk_Fail_No_LowNote(t *testing.T) {
	b := []byte{0x3c, 0x00, 0x00}
	r := bytes.NewReader(b)

	i, err := readInstChunk(r)

	if i != nil {
		t.Error("returned inst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadInstChunk_Fail_No_HighNote(t *testing.T) {
	b := []byte{0x3c, 0x00, 0x00, 0x00}
	r := bytes.NewReader(b)

	i, err := readInstChunk(r)

	if i != nil {
		t.Error("returned inst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadInstChunk_Fail_No_LowVelocity(t *testing.T) {
	b := []byte{0x3c, 0x00, 0x00, 0x00, 0x7f}
	r := bytes.NewReader(b)

	i, err := readInstChunk(r)

	if i != nil {
		t.Error("returned inst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadInstChunk_Fail_No_HighVelocity(t *testing.T) {
	b := []byte{0x3c, 0x00, 0x00, 0x00, 0x7f, 0x01}
	r := bytes.NewReader(b)

	i, err := readInstChunk(r)

	if i != nil {
		t.Error("returned inst chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}
