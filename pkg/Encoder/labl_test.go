package Encoder

import (
	"bytes"
	"encoding/binary"
	"testing"
	"wav-concat/pkg/Decoder"
)

func TestLabel_WriteTo(t *testing.T) {
	l := Label{
		CuePointID: "AAAA",
		Data:       "this is labl data",
	}

	var b bytes.Buffer
	_, err := l.WriteTo(&b)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	b32 := make([]byte, 4)
	_, err = r.Read(b32)

	if string(b32[:]) != LABL_CHUNK_ID {
		t.Errorf("first four bytes are not %s", LABL_CHUNK_ID)
	}

	_, err = r.Read(b32)

	dlc, err := Decoder.ReadLABLChunk(r, binary.LittleEndian.Uint32(b32))
	if err != nil {
		t.Error(err.Error())
	}

	if dlc.CuePointID != "AAAA" {
		t.Error("labl cue point ID is incorrect")
	}

	if dlc.Data != "this is labl data" {
		t.Error("labl data is incorrect")
	}
}
