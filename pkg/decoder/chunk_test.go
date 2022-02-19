package decoder

import (
	"bytes"
	"testing"
)

func TestNewChunk(t *testing.T) {

	b := []byte{0x6c, 0x61, 0x62, 0x6c, 0x0b, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	c, err := newChunk(r)
	if err != nil {
		t.Error(err.Error())
	}

	if c.ID != "labl" {
		t.Errorf("chunk ID is incorrect. Got %s", c.ID)
	}

	if c.Size != 11 {
		t.Errorf("chunk size is incorrect. Got %d", c.Size)
	}

}
