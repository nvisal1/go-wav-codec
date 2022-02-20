package decoder

import (
	"bytes"
	"testing"
)

func TestReadInstChunk(t *testing.T) {
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
