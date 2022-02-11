package Encoder

import (
	"bytes"
	"testing"
	"wav-concat/pkg/Decoder"
)

func TestFactChunk_WriteTo(t *testing.T) {
	fc := FactChunk{
		NumberOfSamples: 0,
	}

	var b bytes.Buffer
	_, err := fc.WriteTo(&b)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	b32 := make([]byte, 4)
	_, err = r.Read(b32)

	if string(b32[:]) != FACT_CHUNK_ID {
		t.Errorf("The first 4 bytes are not %s", FACT_CHUNK_ID)
	}

	_, err = r.Read(b32)

	dfc, err := Decoder.ReadFactChunk(r)
	if err != nil {
		t.Error(err.Error())
	}

	if dfc.NumberOfSamples != 0 {
		t.Error("Fact number of samples is incorrect")
	}

}
