package decoder

import (
	"bytes"
	"io"
	"testing"
)

func TestReadLTXTChunk(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72,
		0x67, 0x6e, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x54, 0x65, 0x78, 0x74, 0x20, 0x44, 0x61,
		0x74, 0x61}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if l.SampleLength != 176400 {
		t.Errorf("Error: did not find the correct sample length. Found %d", l.SampleLength)
	}

	if l.PurposeID != "rgn " {
		t.Errorf("Error: did not find the correct purpose ID. Found %s", l.PurposeID)
	}

	if l.Data != "Text Data" {
		t.Errorf("Error: did not find the correct data. Found %s", l.Data)
	}
}

func TestReadLTXTChunk_Fail_No_Cue_Point_ID(t *testing.T) {
	b := []byte{}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_Short_Cue_Point_ID(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_No_Sample_Length(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_Short_Sample_Length(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_No_Purpose_ID(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_Short_Purpose_ID(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_No_Country(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72,
		0x67, 0x6e, 0x20}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_Short_Country(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72,
		0x67, 0x6e, 0x20, 0x00}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_No_Language(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72,
		0x67, 0x6e, 0x20, 0x00, 0x00}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_Short_Language(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72,
		0x67, 0x6e, 0x20, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_No_Dialect(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72,
		0x67, 0x6e, 0x20, 0x00, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_Short_Dialect(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72,
		0x67, 0x6e, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_No_Code_Page(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72,
		0x67, 0x6e, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_Short_Code_Page(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72,
		0x67, 0x6e, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00}

	r := bytes.NewReader(b)

	l, err := readLTXTChunk(r, uint32(len(b)))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Success_No_Data(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72,
		0x67, 0x6e, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00}

	r := bytes.NewReader(b)

	// Length is 20
	l, err := readLTXTChunk(r, uint32(len(b)))

	if err != nil {
		t.Error(err.Error())
	}

	if l == nil {
		t.Error("returned ltxt chunk is nil")
	}

	if l.PurposeID != "rgn " {
		t.Errorf("expected \"rgn \". received \"%s\"", l.PurposeID)
	}

	if l.Data != "" {
		t.Errorf("expected \" \". received \"%s\"", l.Data)
	}
}

func TestReadLTXTChunk_Fail_No_Data(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72,
		0x67, 0x6e, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00}

	r := bytes.NewReader(b)

	// Length is 21
	// Function looks for another byte but cannot find it
	l, err := readLTXTChunk(r, uint32(len(b)+1))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadLTXTChunk_Fail_Short_Data(t *testing.T) {
	b := []byte{0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02, 0x00, 0x72,
		0x67, 0x6e, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	// Length is 22
	// Function looks for 2 more bytes but only finds 1
	l, err := readLTXTChunk(r, uint32(len(b)+2))

	if l != nil {
		t.Error("returned ltxt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}
