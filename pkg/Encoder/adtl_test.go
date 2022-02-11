package Encoder

import (
	"bytes"
	"testing"
	"wav-concat/pkg/Decoder"
)

func TestADTLChunk_WriteTo(t *testing.T) {
	l := &Label{
		CuePointID: "AAAA",
		Data:       "This is label data",
	}
	ls := make([]*Label, 0)
	ls = append(ls, l)
	n := &Note{
		CuePointID: "BBBB",
		Data:       "This is note data",
	}
	ns := make([]*Note, 0)
	ns = append(ns, n)
	lt := &LabeledText{
		CuePointID:   "CCCC",
		SampleLength: 0,
		PurposeID:    "DDDD",
		Country:      "US",
		Language:     "EN",
		Dialect:      "EN",
		CodePage:     "PO",
		Data:         "This is ltxt data",
	}
	lts := make([]*LabeledText, 0)
	lts = append(lts, lt)
	adtl := &ADTLChunk{
		Labels:       ls,
		Notes:        ns,
		LabeledTexts: lts,
	}

	var b bytes.Buffer
	_, err := adtl.WriteTo(&b)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	adtlID := make([]byte, 4)
	_, err = r.Read(adtlID)

	if string(adtlID[:]) != "adtl" {
		t.Error("First 4 bytes is not adtl")
	}

	ac, err := Decoder.ReadADTLChunk(r)

	if len(ac.Labels) != 1 {
		t.Error("Labels is not the correct length")
	}

	if ac.Labels[0].CuePointID != "AAAA" {
		t.Error("Label cue point ID is incorrect")
	}

	if ac.Labels[0].Data != "This is label data" {
		t.Error("Label data is incorrect")
	}

	if len(ac.Notes) != 1 {
		t.Error("Notes is not the correct length")
	}

	if ac.Notes[0].CuePointID != "BBBB" {
		t.Error("Note cue point ID is incorrect")
	}

	if ac.Notes[0].Data != "This is note data" {
		t.Error("Note data is incorrect")
	}

	if len(ac.LabeledTexts) != 1 {
		t.Error("LabeledTexts is not the correct length")
	}

	if ac.LabeledTexts[0].CuePointID != "CCCC" {
		t.Error("ltxt cue point ID is incorrect")
	}

	if ac.LabeledTexts[0].SampleLength != 0 {
		t.Error("ltxt sample length is incorrect")
	}

	if ac.LabeledTexts[0].PurposeID != "DDDD" {
		t.Error("ltxt purpose ID is incorrect")
	}

	if ac.LabeledTexts[0].Country != "US" {
		t.Error("ltxt country is incorrect")
	}

	if ac.LabeledTexts[0].Language != "EN" {
		t.Error("ltxt language is incorrect")
	}

	if ac.LabeledTexts[0].Dialect != "EN" {
		t.Error("ltxt dialect is incorrect")
	}

	if ac.LabeledTexts[0].CodePage != "PO" {
		t.Error("ltxt code page is incorrect")
	}

	if ac.LabeledTexts[0].Data != "This is ltxt data" {
		t.Errorf("ltxt data is incorrect. Actual: %s", ac.LabeledTexts[0].Data)
	}

}
