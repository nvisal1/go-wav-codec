package Encoder

import (
	"bytes"
	"testing"
	"wav-concat/pkg/Decoder"
)

func TestPLSTChunk_WriteTo(t *testing.T) {
	ps := &PlstSegment{
		CuePointID:      "AAAA",
		Length:          0,
		NumberOfRepeats: 0,
	}

	psl := make([]*PlstSegment, 0)
	psl = append(psl, ps)

	pc := PLSTChunk{plsts: psl}

	var b bytes.Buffer
	_, err := pc.WriteTo(&b)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	b32 := make([]byte, 4)
	_, err = r.Read(b32)

	if string(b32[:]) != PLST_CHUNK_ID {
		t.Errorf("first 4 bytes are not %s", PLST_CHUNK_ID)
	}

	_, err = r.Read(b32)

	dpc, err := Decoder.ReadPlstChunk(r)
	if err != nil {
		t.Error(err.Error())
	}

	if len(dpc) != 1 {
		t.Error("length of plst segments is incorrect")
	}

	if dpc[0].CuePointID != "AAAA" {
		t.Error("plst segment cue point ID is incorrect")
	}

	if dpc[0].Length != 0 {
		t.Error("plst segment length is incorrect")
	}

	if dpc[0].NumberOfRepeats != 0 {
		t.Error("plst segment number of repeats is incorrect")
	}
}
