package decoder

import (
	"bytes"
	"testing"
)

func TestReadCueChunk(t *testing.T) {
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
