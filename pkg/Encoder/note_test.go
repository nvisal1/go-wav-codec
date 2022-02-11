package Encoder

import (
	"bytes"
	"encoding/binary"
	"testing"
	"wav-concat/pkg/Decoder"
)

func TestNote_WriteTo(t *testing.T) {
	n := Note{
		CuePointID: "AAAA",
		Data:       "this is note data",
	}

	var b bytes.Buffer
	_, err := n.WriteTo(&b)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	b32 := make([]byte, 4)
	_, err = r.Read(b32)

	if string(b32[:]) != NOTE_CHUNK_ID {
		t.Errorf("first 4 bytes are not %s", NOTE_CHUNK_ID)
	}

	_, err = r.Read(b32)

	dnc, err := Decoder.ReadNoteChunk(r, binary.LittleEndian.Uint32(b32))
	if err != nil {
		t.Error(err.Error())
	}

	if dnc.CuePointID != "AAAA" {
		t.Error("note cue point ID is incorrect")
	}

	if dnc.Data != "this is note data" {
		t.Error("note data is incorrect")
	}
}
