package Decoder

import (
	"bytes"
	"testing"
)

func TestReadLTXTChunk(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72,
		0x67, 0x6e, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x54, 0x65, 0x78, 0x74, 0x20, 0x44, 0x61,
		0x74, 0x61}

	r := bytes.NewReader(b)

	l, err := ReadLTXTChunk(r, uint32(len(b)))
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if l.SampleLength != 176400 {
		t.Errorf("Error: did not find the correct sample length. Found %d", l.SampleLength)
	}

	if l.PurposeID != "rgn " {
		t.Errorf("Error: did not find the correct purpose ID. Found %s", l.PurposeID)
	}

	if l.Data != "Text Data" {
		t.Errorf("Error: did not find the correct data. Found %s", l.Data)
	}
}
