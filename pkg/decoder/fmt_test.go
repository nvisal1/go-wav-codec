package decoder

import (
	"bytes"
	"io"
	"testing"
)

func TestReadFMTChunk_Success(t *testing.T) {
	b := []byte{
		0x01, 0x00, // AudioFormat
		0x02, 0x00, // NumChannels
		0x44, 0xac, 0x00, 0x00, // SampleRate
		0x98, 0x09, 0x04, 0x00, // ByteRate
		0x06, 0x00, // BlockAlign
		0x18, 0x00, // BitsPerSample
	}

	r := bytes.NewReader(b)

	f, err := readFMTChunk(r)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if f.AudioFormat != 1 {
		t.Errorf("Error: did not find the correct audio format. Found %d", f.AudioFormat)
	}

	if f.NumChannels != 2 {
		t.Errorf("Error: did not find the correct number of channels. Found %d", f.NumChannels)
	}

	if f.SampleRate != 44100 {
		t.Errorf("Error: did not find the correct sample rate. Found %d", f.SampleRate)
	}

	if f.ByteRate != 264600 {
		t.Errorf("Error: did not find the correct byte rate. Found %d", f.ByteRate)
	}

	if f.BlockAlign != 6 {
		t.Errorf("Error: did not find the correct block align. Found %d", f.BlockAlign)
	}

	if f.BitsPerSample != 24 {
		t.Errorf("Error: did not find the correct bits per sample. Found %d", f.BitsPerSample)
	}
}

func TestReadFMTChunk_Fail_No_AudioFormat(t *testing.T) {
	b := []byte{}

	r := bytes.NewReader(b)

	f, err := readFMTChunk(r)

	if f != nil {
		t.Error("returned fmt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("returned error is not \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadFMTChunk_Fail_Short_AudioFormat(t *testing.T) {
	b := []byte{0x01}

	r := bytes.NewReader(b)

	f, err := readFMTChunk(r)

	if f != nil {
		t.Error("returned fmt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("returned error is not \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadFMTChunk_Fail_No_NumChannels(t *testing.T) {
	b := []byte{0x01, 0x00}

	r := bytes.NewReader(b)

	f, err := readFMTChunk(r)

	if f != nil {
		t.Error("returned fmt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("returned error is not \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadFMTChunk_Fail_Short_NumChannels(t *testing.T) {
	b := []byte{0x01, 0x00, // AudioFormat
		0x02}

	r := bytes.NewReader(b)

	f, err := readFMTChunk(r)

	if f != nil {
		t.Error("returned fmt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("returned error is not \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadFMTChunk_Fail_No_SampleRate(t *testing.T) {
	b := []byte{0x01, 0x00, // AudioFormat
		0x02, 0x00, // NumChannels
	}

	r := bytes.NewReader(b)

	f, err := readFMTChunk(r)

	if f != nil {
		t.Error("returned fmt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("returned error is not \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadFMTChunk_Fail_Short_SampleRate(t *testing.T) {
	b := []byte{
		0x01, 0x00, // AudioFormat
		0x02, 0x00, // NumChannels
		0x44,
	}

	r := bytes.NewReader(b)

	f, err := readFMTChunk(r)

	if f != nil {
		t.Error("returned fmt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("returned error is not \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadFMTChunk_Fail_No_ByteRate(t *testing.T) {
	b := []byte{0x01, 0x00, // AudioFormat
		0x02, 0x00, // NumChannels
		0x44, 0xac, 0x00, 0x00, // SampleRate
	}

	r := bytes.NewReader(b)

	f, err := readFMTChunk(r)

	if f != nil {
		t.Error("returned fmt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("returned error is not \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadFMTChunk_Fail_Short_ByteRate(t *testing.T) {
	b := []byte{
		0x01, 0x00, // AudioFormat
		0x02, 0x00, // NumChannels
		0x44, 0xac, 0x00, 0x00, // SampleRate
		0x98,
	}

	r := bytes.NewReader(b)

	f, err := readFMTChunk(r)

	if f != nil {
		t.Error("returned fmt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("returned error is not \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadFMTChunk_Fail_No_BlockAlign(t *testing.T) {
	b := []byte{0x01, 0x00, // AudioFormat
		0x02, 0x00, // NumChannels
		0x44, 0xac, 0x00, 0x00, // SampleRate
		0x98, 0x09, 0x04, 0x00, // ByteRate
	}

	r := bytes.NewReader(b)

	f, err := readFMTChunk(r)

	if f != nil {
		t.Error("returned fmt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("returned error is not \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadFMTChunk_Fail_Short_BlockAlign(t *testing.T) {
	b := []byte{
		0x01, 0x00, // AudioFormat
		0x02, 0x00, // NumChannels
		0x44, 0xac, 0x00, 0x00, // SampleRate
		0x98, 0x09, 0x04, 0x00, // ByteRate
		0x06,
	}

	r := bytes.NewReader(b)

	f, err := readFMTChunk(r)

	if f != nil {
		t.Error("returned fmt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("returned error is not \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadFMTChunk_Fail_No_BitsPerSample(t *testing.T) {
	b := []byte{0x01, 0x00, // AudioFormat
		0x02, 0x00, // NumChannels
		0x44, 0xac, 0x00, 0x00, // SampleRate
		0x98, 0x09, 0x04, 0x00, // ByteRate
		0x06, 0x00,
	}

	r := bytes.NewReader(b)

	f, err := readFMTChunk(r)

	if f != nil {
		t.Error("returned fmt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("returned error is not \"EOF\". received \"%s\"", err.Error())
	}
}

func TestReadFMTChunk_Fail_Short_BitsPerSample(t *testing.T) {
	b := []byte{
		0x01, 0x00, // AudioFormat
		0x02, 0x00, // NumChannels
		0x44, 0xac, 0x00, 0x00, // SampleRate
		0x98, 0x09, 0x04, 0x00, // ByteRate
		0x06, 0x00, // BlockAlign
		0x18,
	}

	r := bytes.NewReader(b)

	f, err := readFMTChunk(r)

	if f != nil {
		t.Error("returned fmt chunk is not nil")
	}

	if err == nil {
		t.Errorf("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("returned error is not \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}
