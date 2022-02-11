package Encoder

import (
	"bytes"
	"testing"
	"wav-concat/pkg/Decoder"
)

func TestInstChunk_WriteTo(t *testing.T) {
	ic := InstChunk{
		UnshiftedNote: 0,
		FineTuneDB:    0,
		Gain:          1,
		LowNote:       0,
		HighNote:      1,
		LowVelocity:   0,
		HighVelocity:  1,
	}

	var b bytes.Buffer
	_, err := ic.WriteTo(&b)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	b32 := make([]byte, 4)
	_, err = r.Read(b32)

	if string(b32[:]) != INST_CHUNK_ID {
		t.Errorf("first four bytes are not %s", INST_CHUNK_ID)
	}

	_, err = r.Read(b32)

	dic, err := Decoder.ReadInstChunk(r)
	if err != nil {
		t.Error(err.Error())
	}

	if dic.UnshiftedNote != 0 {
		t.Error("inst unshifted note is incorrect")
	}

	if dic.FineTuneDB != 0 {
		t.Error("inst fine tune db is incorrect")
	}

	if dic.Gain != 1 {
		t.Error("inst gain is incorrect")
	}

	if dic.LowNote != 0 {
		t.Error("inst low note is incorrect")
	}

	if dic.HighNote != 1 {
		t.Error("inst high note is incorrect")
	}

	if dic.LowVelocity != 0 {
		t.Error("inst low velocity is incorrect")
	}

	if dic.HighVelocity != 1 {
		t.Error("inst high velocity is incorrect")
	}
}
