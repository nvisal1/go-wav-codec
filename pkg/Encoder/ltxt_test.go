package Encoder

import (
	"bytes"
	"encoding/binary"
	"testing"
	"wav-concat/pkg/Decoder"
)

func TestLabeledText_WriteTo(t *testing.T) {
	lt := LabeledText{
		CuePointID:   "AAAA",
		SampleLength: 0,
		PurposeID:    "BBBB",
		Country:      "CC",
		Language:     "DD",
		Dialect:      "EE",
		CodePage:     "FF",
		Data:         "GGG",
	}

	var b bytes.Buffer
	_, err := lt.WriteTo(&b)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	b32 := make([]byte, 4)
	_, err = r.Read(b32)

	if string(b32[:]) != TEXT_CHUNK_ID {
		t.Errorf("first 4 bytes are not %s", TEXT_CHUNK_ID)
	}

	_, err = r.Read(b32)

	dltc, err := Decoder.ReadLTXTChunk(r, binary.LittleEndian.Uint32(b32))
	if err != nil {
		t.Error(err.Error())
	}

	if dltc.CuePointID != "AAAA" {
		t.Error("ltxt cue point ID is incorrect")
	}

	if dltc.SampleLength != 0 {
		t.Error("ltxt sample length is incorrect")
	}

	if dltc.PurposeID != "BBBB" {
		t.Error("ltxt purpose ID is incorrect")
	}

	if dltc.Country != "CC" {
		t.Error("ltxt country is incorrect")
	}

	if dltc.Language != "DD" {
		t.Error("ltxt language is incorrect")
	}

	if dltc.Dialect != "EE" {
		t.Error("ltxt dialect is incorrect")
	}

	if dltc.CodePage != "FF" {
		t.Error("ltxt code page is incrrect")
	}

	if dltc.Data != "GGG" {
		t.Error("ltxt data is incorrect")
	}
}
