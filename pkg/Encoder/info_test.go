package Encoder

import (
	"bytes"
	"testing"
	"wav-concat/pkg/Decoder"
)

func TestInfoChunk_WriteTo(t *testing.T) {
	ic := InfoChunk{
		Location:     "AAA",
		Artist:       "BBBB",
		Software:     "C",
		CreationDate: "DD",
		Copyright:    "EEEEE",
		Title:        "FFFFFF",
		Engineer:     "GGGGGGG",
		Genre:        "",
		Product:      "HHHHHHHH",
		Source:       "I",
		Subject:      "JJ",
		Comments:     "K",
		Technician:   "LLLL",
		Keywords:     "",
		Medium:       "",
	}

	var b bytes.Buffer
	_, err := ic.WriteTo(&b)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	b32 := make([]byte, 4)
	_, err = r.Read(b32)

	if string(b32) != INFO_CHUNK_ID {
		t.Errorf("first 4 bytes are not %s", INFO_CHUNK_ID)
	}

	dic, err := Decoder.ReadINFOChunk(r)
	if err != nil {
		t.Error(err.Error())
	}

	if dic.Location != "AAA" {
		t.Error("info location is incorrect")
	}
}
