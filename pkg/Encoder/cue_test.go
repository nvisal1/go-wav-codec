package Encoder

import (
	"bytes"
	"testing"
	"wav-concat/pkg/Decoder"
)

func TestCueChunk_WriteTo(t *testing.T) {
	cp := &CuePoint{
		ID:           "AAAA",
		Position:     0,
		DataChunkID:  "AAAA",
		ChunkStart:   0,
		BlockStart:   0,
		SampleOffset: 0,
	}
	cps := make([]*CuePoint, 0)
	cps = append(cps, cp)

	cc := CueChunk{
		cps: cps,
	}

	var b bytes.Buffer
	_, err := cc.WriteTo(&b)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	cueID := make([]byte, 4)
	_, err = r.Read(cueID)

	if string(cueID[:]) != CUE_CHUNK_ID {
		t.Errorf("First 4 bytes are not %s", CUE_CHUNK_ID)
	}
	_, err = r.Read(cueID)

	dcps, err := Decoder.ReadCueChunk(r)
	if err != nil {
		t.Error(err.Error())
	}

	if len(dcps) != 1 {
		t.Error("cue point length is incorrect")
	}

	if dcps[0].ID != "AAAA" {
		t.Error("cue point ID is incorrect")
	}

	if dcps[0].Position != 0 {
		t.Error("cue point position is incorrect")
	}

	if dcps[0].DataChunkID != "AAAA" {
		t.Error("cue point data chunk ID is incorrect")
	}

	if dcps[0].ChunkStart != 0 {
		t.Error("cue point chunk start is incorrect")
	}

	if dcps[0].BlockStart != 0 {
		t.Error("cue point block start is incorrect")
	}

	if dcps[0].SampleOffset != 0 {
		t.Error("cue point sample offset is incorrect")
	}
}
