package decoder

import (
	"bytes"
	"io"
	"testing"
)

func TestReadADTLChunk_Success_With_Word_Align(t *testing.T) {

	b := []byte{0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x10,
		0xb1, 0x02, 0x00, 0x72, 0x67, 0x6e, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x6c, 0x61, 0x62, 0x6c, 0x0b, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x4c, 0x6f, 0x6f,
		0x70, 0x20, 0x31, 0x00, 0x00, 0x6e, 0x6f, 0x74, 0x65, 0x11, 0x00, 0x00, 0x00, 0x01, 0x00,
		0x00, 0x00, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x20, 0x6c, 0x6f, 0x6f, 0x70}

	r := bytes.NewReader(b)

	a, err := readADTLChunk(r)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if len(a.LabeledTexts) != 1 {
		t.Error("length of ltxt is incorrect")
	}

	if len(a.Labels) != 1 {
		t.Error("length of labl is incorrect")
	}

	if len(a.Notes) != 1 {
		t.Error("length of note is incorrect")
	}

}

func TestReadADTLChunk_Fail_With_No_Data_LTXT(t *testing.T) {
	b := []byte{0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	a, err := readADTLChunk(r)

	if a != nil {
		t.Error("returned adtl chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
		if err != io.EOF {
			t.Error("returned error is not EOF")
		}
	}
}

func TestReadADTLChunk_Fail_With_Short_Data_LTXT(t *testing.T) {
	b := []byte{0x6c, 0x74, 0x78, 0x74, 0x14, 0x00, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	a, err := readADTLChunk(r)

	if a != nil {
		t.Error("returned adtl chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Error("returned error is not unexpected EOF")
	}

}

func TestReadADTLChunk_Fail_With_No_Data_LABL(t *testing.T) {
	b := []byte{0x6c, 0x61, 0x62, 0x6c, 0x0b, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	a, err := readADTLChunk(r)

	if a != nil {
		t.Error("returned adtl chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
		if err != io.EOF {
			t.Error("returned error is not EOF")
		}
	}
}

func TestReadADTLChunk_Fail_With_Short_Data_LABL(t *testing.T) {
	b := []byte{0x6c, 0x61, 0x62, 0x6c, 0x0b, 0x00, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	a, err := readADTLChunk(r)

	if a != nil {
		t.Error("returned adtl chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Error("returned error is not unexpected EOF")
	}

}

func TestReadADTLChunk_Fail_With_No_Data_NOTE(t *testing.T) {
	b := []byte{0x6e, 0x6f, 0x74, 0x65, 0x11, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	a, err := readADTLChunk(r)

	if a != nil {
		t.Error("returned adtl chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
		if err != io.EOF {
			t.Error("returned error is not EOF")
		}
	}
}

func TestReadADTLChunk_Fail_With_Short_Data_NOTE(t *testing.T) {
	b := []byte{0x6e, 0x6f, 0x74, 0x65, 0x11, 0x00, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	a, err := readADTLChunk(r)

	if a != nil {
		t.Error("returned adtl chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Error("returned error is not unexpected EOF")
	}

}

func TestReadADTLChunk_Fail_With_No_Size_Chunk_Header(t *testing.T) {
	b := []byte{0x6e, 0x6f, 0x74, 0x65}

	r := bytes.NewReader(b)

	a, err := readADTLChunk(r)

	if a == nil {
		t.Error("returned adtl chunk is nil")
	}

	if err != nil {
		t.Error("returned error is not nil")
	}

	if len(a.LabeledTexts) != 0 {
		t.Error("length of ltxt is incorrect")
	}

	if len(a.Labels) != 0 {
		t.Error("length of labl is incorrect")
	}

	if len(a.Notes) != 0 {
		t.Error("length of note is incorrect")
	}

}

func TestReadADTLChunk_Fail_With_No_ID_Chunk_Header(t *testing.T) {
	b := []byte{}

	r := bytes.NewReader(b)

	a, err := readADTLChunk(r)

	if a == nil {
		t.Error("returned adtl chunk is nil")
	}

	if err != nil {
		t.Error("returned error is not nil")
	}

	if len(a.LabeledTexts) != 0 {
		t.Error("length of ltxt is incorrect")
	}

	if len(a.Labels) != 0 {
		t.Error("length of labl is incorrect")
	}

	if len(a.Notes) != 0 {
		t.Error("length of note is incorrect")
	}

}

func TestReadADTLChunk_Fail_With_Short_ID_Chunk_Header(t *testing.T) {
	b := []byte{0x00}

	r := bytes.NewReader(b)

	a, err := readADTLChunk(r)

	if a != nil {
		t.Error("returned adtl chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Error("returned error is not unexpected EOF")
	}

}

func TestReadADTLChunk_Fail_With_Short_Size_Chunk_Header(t *testing.T) {
	b := []byte{0x6e, 0x6f, 0x74, 0x65, 0x00}

	r := bytes.NewReader(b)

	a, err := readADTLChunk(r)

	if a != nil {
		t.Error("returned adtl chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Error("returned error is not unexpected EOF")
	}

}
