package decoder

import (
	"bytes"
	"testing"
)

func TestReadPlstChunk(t *testing.T) {

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
