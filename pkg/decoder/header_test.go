package decoder

import (
	"bytes"
	"io"
	"testing"
)

func TestReadWavHeader_Success(t *testing.T) {
	b := []byte{0x52, 0x49, 0x46, 0x46, 0x18, 0x6f, 0x28, 0x00, 0x57, 0x41, 0x56, 0x45}

	r := bytes.NewReader(b)

	wfh, err := readWavFileHeader(r)
	if err != nil {
		t.Error(err.Error())
	}

	if wfh.FileSize != 2649880 {
		t.Errorf("expected file size %d. received file size %d", 2649880, wfh.FileSize)
	}
}

func TestReadWavHeader_Fail_No_ID(t *testing.T) {
	b := []byte{}

	r := bytes.NewReader(b)

	wfh, err := readWavFileHeader(r)

	if wfh != nil {
		t.Error("returned wav file header is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("returned error is not \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavHeader_Fail_Short_ID(t *testing.T) {
	b := []byte{0x52}

	r := bytes.NewReader(b)

	wfh, err := readWavFileHeader(r)

	if wfh != nil {
		t.Error("returned wav file header is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("returned error is not \"Unexpected EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavHeader_Fail_ID_Not_RIFF(t *testing.T) {
	b := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	wfh, err := readWavFileHeader(r)

	if wfh != nil {
		t.Error("returned wav file header is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err.Error() != "File descriptor is not RIFF" {
		t.Errorf("returned error is not \"File descriptor is not RIFF\". received \"%s\"", err.Error())
	}
}

func TestReadWavHeader_Fail_No_Size(t *testing.T) {
	b := []byte{0x52, 0x49, 0x46, 0x46}

	r := bytes.NewReader(b)

	wfh, err := readWavFileHeader(r)

	if wfh != nil {
		t.Error("returned wav file header is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("returned error is not \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadWavHeader_Fail_Short_Size(t *testing.T) {
	b := []byte{0x52, 0x49, 0x46, 0x46, 0x18}

	r := bytes.NewReader(b)

	wfh, err := readWavFileHeader(r)

	if wfh != nil {
		t.Error("returned wav file header is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("returned error is not \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadWavHeader_Fail_Size_Is_0(t *testing.T) {
	b := []byte{0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	wfh, err := readWavFileHeader(r)

	if wfh != nil {
		t.Error("returned wav file header is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err.Error() != "file size is less than or equal to 0" {
		t.Errorf("returned error is not \"file size is less than or equal to 0\". received \"%s\"", err.Error())
	}
}

func TestReadWavHeader_Fail_No_File_Format(t *testing.T) {
	b := []byte{0x52, 0x49, 0x46, 0x46, 0x18, 0x6f, 0x28, 0x00}

	r := bytes.NewReader(b)

	wfh, err := readWavFileHeader(r)

	if wfh != nil {
		t.Error("returned wav file header is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("returned error is not \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadWavHeader_Fail_Short_File_Format(t *testing.T) {
	b := []byte{0x52, 0x49, 0x46, 0x46, 0x18, 0x6f, 0x28, 0x00, 0x57}

	r := bytes.NewReader(b)

	wfh, err := readWavFileHeader(r)

	if wfh != nil {
		t.Error("returned wav file header is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("returned error is not \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadWavHeader_Fail_File_Format_Not_WAVE(t *testing.T) {
	b := []byte{0x52, 0x49, 0x46, 0x46, 0x18, 0x6f, 0x28, 0x00, 0x00, 0x00, 0x00, 0x00}

	r := bytes.NewReader(b)

	wfh, err := readWavFileHeader(r)

	if wfh != nil {
		t.Error("returned wav file header is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err.Error() != "File format is not WAVE. Actual \u0000\u0000\u0000\u0000" {
		t.Errorf("returned error is not \"File format is not WAVE. Actual \u0000\u0000\u0000\u0000\". received \"%s\"", err.Error())
	}
}
