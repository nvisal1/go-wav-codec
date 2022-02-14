package Encoder

import (
	"bytes"
	"testing"
	"wav-concat/pkg/Decoder"
)

func TestSmplChunk_WriteTo(t *testing.T) {
	sl := &SampleLoop{
		CuePointID: "AAAA",
		Type:       0,
		Start:      0,
		End:        0,
		Fraction:   0,
		PlayCount:  0,
	}
	sll := make([]*SampleLoop, 0)
	sll = append(sll, sl)

	sc := SmplChunk{
		Manufacturer:      "BBBB",
		Product:           "CCCC",
		SamplePeriod:      0,
		MIDIUnityNote:     0,
		MIDIPitchFraction: 0,
		SMPTEFormat:       0,
		SMPTEOffset:       0,
		Loops:             sll,
	}

	var b bytes.Buffer
	_, err := sc.WriteTo(&b)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	b32 := make([]byte, 4)
	_, err = r.Read(b32)

	if string(b32[:]) != SMPL_CHUNK_ID {
		t.Errorf("the first 4 bytes are not %s", SMPL_CHUNK_ID)
	}

	_, err = r.Read(b32)

	dsc, err := Decoder.ReadSmplChunk(r)
	if err != nil {
		t.Error(err.Error())
	}

	if len(dsc.Loops) != 1 {
		t.Error("smpl loop count is incorrect")
	}

	if dsc.Manufacturer != "BBBB" {
		t.Error("smpl manufacturer is incorrect")
	}

	if dsc.Product != "CCCC" {
		t.Error("smpl product is incorrect")
	}

	if dsc.SamplePeriod != 0 {
		t.Error("smpl sample period is incorrect")
	}

	if dsc.MIDIUnityNote != 0 {
		t.Error("smpl midi unity note is incorrect")
	}

	if dsc.MIDIPitchFraction != 0 {
		t.Error("smpl midi pitch fraction is incorrect")
	}

	if dsc.SMPTEFormat != 0 {
		t.Error("smpl smpte format is incorrect")
	}

	if dsc.SMPTEOffset != 0 {
		t.Error("smpl smpte offset is incorrect")
	}

	if dsc.Loops[0].CuePointID != "AAAA" {
		t.Error("smpl loop cue point ID is incorrect")
	}

	if dsc.Loops[0].Type != 0 {
		t.Error("smpl loop type is incorrect")
	}

	if dsc.Loops[0].Start != 0 {
		t.Error("smpl loop start is incorrect")
	}

	if dsc.Loops[0].End != 0 {
		t.Error("smpl loop end is incorrect")
	}

	if dsc.Loops[0].Fraction != 0 {
		t.Error("smpl loop fraction is incorrect")
	}

	if dsc.Loops[0].PlayCount != 0 {
		t.Error("smpl loop play count is incorrect")
	}
}
