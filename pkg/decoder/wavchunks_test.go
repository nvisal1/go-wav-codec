package decoder

import (
	"bytes"
	"io"
	"testing"
)

func TestReadWavChunks_Success_No_Chunk_ID(t *testing.T) {
	b := []byte{}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc == nil {
		t.Error("returned wav chunks is nil")
	}

	if wc.FMT != nil {
		t.Error("returned fmt is not nil")
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func TestReadWavChunks_Fail_Short_Chunk_ID(t *testing.T) {
	b := []byte{
		// FMT
		0x66, 0x6d, 0x74, // Chunk ID
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadWavChunks_Success_No_Chunk_Size(t *testing.T) {
	b := []byte{
		// FMT
		0x66, 0x6d, 0x74, 0x20, // Chunk ID
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc == nil {
		t.Error("returned wav chunks is nil")
	}

	if wc.FMT != nil {
		t.Error("returned fmt is not nil")
	}

	if err != nil {
		t.Error(err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_Short_Size_FMT(t *testing.T) {
	b := []byte{
		// FMT
		0x66, 0x6d, 0x74, 0x20, // Chunk ID
		0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected \"Unexpected EOF\". received \"%s\"", err.Error())
	}
}

func TestReadWavChunks_Fail_Data_Before_FMT(t *testing.T) {
	b := []byte{
		// DATA
		0x64, 0x61, 0x74, 0x61,
		0x08, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// FMT
		0x66, 0x6d, 0x74, 0x20, // Chunk ID
		0x10, 0x00, 0x00, 0x00, // Chunk Size
		0x01, 0x00, // AudioFormat
		0x01, 0x00, // NumChannels
		0x40, 0x1f, 0x00, 0x00, // Sample Rate
		0x80, 0x3e, 0x00, 0x00, // ByteRate
		0x02, 0x00, // Block Align
		0x10, 0x00, // BitsPerSample
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned chunk is not nil")
	}

	if err == nil {
		t.Error("returned error is not nil")
	}

	if err.Error() != "Data chunk was found before fmt chunk" {
		t.Errorf("expected \"Data chunk was found before fmt chunk\". receieved\"%s\"", err.Error())
	}
}

func TestReadWavChunks_Success_Unsupported_Chunk_ID_Odd_Chunk_Size(t *testing.T) {
	b := []byte{
		0x66, 0x0c, 0x0c, 0x20, // Chunk ID
		0x07, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// FMT
		0x66, 0x6d, 0x74, 0x20, // Chunk ID
		0x10, 0x00, 0x00, 0x00, // Chunk Size
		0x01, 0x00, // AudioFormat
		0x01, 0x00, // NumChannels
		0x44, 0xac, 0x00, 0x00, // Sample Rate
		0x80, 0x3e, 0x00, 0x00, // ByteRate
		0x02, 0x00, // Block Align
		0x10, 0x00, // BitsPerSample
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if err != nil {
		t.Error(err.Error())
	}

	if wc == nil {
		t.Error("returned chunk is nil")
	}

	if wc.FMT == nil {
		t.Error("returned fmt chunk is nil")
	}

	if wc.FMT.AudioFormat != 1 {
		t.Errorf("expected audio format 1. received audio format %d", wc.FMT.AudioFormat)
	}

	if wc.FMT.NumChannels != 1 {
		t.Errorf("expected num channels 1. received num channels %d", wc.FMT.NumChannels)
	}

	if wc.FMT.SampleRate != 44100 {
		t.Errorf("expected sample rate 44100. received sample rate %d", wc.FMT.SampleRate)
	}

	if wc.FMT.ByteRate != 16000 {
		t.Errorf("expected byte rate 16000. received byte rate %d", wc.FMT.ByteRate)
	}

	if wc.FMT.BlockAlign != 2 {
		t.Errorf("expected block align 2. received block align %d", wc.FMT.BlockAlign)
	}

	if wc.FMT.BitsPerSample != 16 {
		t.Errorf("expected bits per sample 16. received bits per sample %d", wc.FMT.BitsPerSample)
	}

}

func TestReadWavChunks_Success_Before_Data(t *testing.T) {
	b := []byte{
		// FMT
		0x66, 0x6d, 0x74, 0x20, // Chunk ID
		0x10, 0x00, 0x00, 0x00, // Chunk Size
		0x01, 0x00, // AudioFormat
		0x01, 0x00, // NumChannels
		0x40, 0x1f, 0x00, 0x00, // Sample Rate
		0x80, 0x3e, 0x00, 0x00, // ByteRate
		0x02, 0x00, // Block Align
		0x10, 0x00, // BitsPerSample

		// LIST
		0x4c, 0x49, 0x53, 0x54, 0x4e, 0x00, 0x00, 0x00, 0x61,
		0x64, 0x74, 0x6c, 0x6c, 0x74, 0x78, 0x74, 0x14, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02,
		0x00, 0x72, 0x67, 0x6e, 0x20, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x6c, 0x61, 0x62, 0x6c, 0x0b,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x4c, 0x6f,
		0x6f, 0x70, 0x20, 0x31, 0x00, 0x00, 0x6e, 0x6f, 0x74,
		0x65, 0x11, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x20, 0x6c,
		0x6f, 0x6f, 0x70, 0x00, 0x00,

		// FACT
		0x46, 0x41, 0x43, 0x54,
		0x04, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00,

		// PLST
		0x50, 0x4c, 0x53, 0x54,
		0x10, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00,
		0x49, 0x44, 0x20, 0x20,
		0x01, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00,

		// SMPL
		0x53, 0x4d, 0x50, 0x4c,
		0x3c, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x93, 0x58, 0x00, 0x00, 0x2d, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x10, 0xb1, 0x02, 0x00, 0x1f, 0x62, 0x05, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		// INST
		0x49, 0x4e, 0x53, 0x54,
		0x08, 0x00, 0x00, 0x00,
		0x3c, 0x00, 0x00, 0x00, 0x7f, 0x01, 0x7f, 0x00,

		// CUE
		0x43, 0x55, 0x45, 0x20,
		0x1c, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0xa0, 0x0f, 0x00, 0x00, 0x64, 0x61, 0x74, 0x61,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xa0, 0x0f, 0x00, 0x00,

		// DATA
		0x64, 0x61, 0x74, 0x61,
		0x08, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)
	if err != nil {
		t.Error(err.Error())
	}

	if wc.FMT == nil {
		t.Error("fmt chunk is nil")
	}

	if wc.ADTL.LabeledTexts == nil {
		t.Error("adtl labeled text is nil")
	}

	if len(wc.ADTL.LabeledTexts) != 1 {
		t.Errorf("Expected 1 labeled text. received %d labeled text(s)", len(wc.ADTL.LabeledTexts))
	}

	if wc.ADTL.Labels == nil {
		t.Error("adtl labels is nil")
	}

	if len(wc.ADTL.Labels) != 1 {
		t.Errorf("Expected 1 label. received %d label(s)", len(wc.ADTL.Labels))
	}

	if wc.ADTL.Notes == nil {
		t.Error("adtl notes is nil")
	}

	if len(wc.ADTL.Notes) != 1 {
		t.Errorf("Expected 1 note. received %d labeled text(s)", len(wc.ADTL.Notes))
	}

	if len(wc.CuePoints) != 1 {
		t.Errorf("Expected 1 cue point. received %d cue point(s)", len(wc.CuePoints))
	}

	if wc.Fact == nil {
		t.Error("fact is nil")
	}

	if wc.Info != nil {
		t.Error("info is not nil")
	}

	if wc.Inst == nil {
		t.Error("inst is nil")
	}

	if len(wc.PlstSegments) != 1 {
		t.Errorf("Expected 1 plst segment. received %d plst segment(s)", len(wc.PlstSegments))
	}

	if wc.Sample == nil {
		t.Error("sample is nil")
	}

	if wc.DataPosition != 274 {
		t.Errorf("Expected data position 274. received data position %d", wc.DataPosition)
	}

	if wc.DataLength != 8 {
		t.Errorf("Expected data length of 8. received data length %d", wc.DataLength)
	}

}

func TestReadWavChunks_Fail_Before_Data_Short_FMT(t *testing.T) {
	b := []byte{
		// FMT
		0x66, 0x6d, 0x74, 0x20, // Chunk ID
		0x08, 0x00, 0x00, 0x00, // Chunk Size
		0x01, 0x00, // AudioFormat
		0x01, 0x00, // NumChannels
		0x40, 0x1f, 0x00, 0x00, // Sample Rate
		0x80, 0x3e, 0x00, 0x00, // ByteRate
		0x02, 0x00, // Block Align
		0x10, 0x00, // BitsPerSample
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_No_FMT(t *testing.T) {
	b := []byte{
		// FMT
		0x66, 0x6d, 0x74, 0x20, // Chunk ID
		0x08, 0x00, 0x00, 0x00, // Chunk Size
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_Short_LIST(t *testing.T) {
	b := []byte{
		// LIST
		0x4c, 0x49, 0x53, 0x54, 0x0c, 0x00, 0x00, 0x00, 0x61,
		0x64, 0x74, 0x6c, 0x6c, 0x74, 0x78, 0x74, 0x14, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02,
		0x00, 0x72, 0x67, 0x6e, 0x20, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x6c, 0x61, 0x62, 0x6c, 0x0b,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x4c, 0x6f,
		0x6f, 0x70, 0x20, 0x31, 0x00, 0x00, 0x6e, 0x6f, 0x74,
		0x65, 0x11, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x20, 0x6c,
		0x6f, 0x6f, 0x70, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_No_LIST(t *testing.T) {
	b := []byte{
		// LIST
		0x4c, 0x49, 0x53, 0x54, 0x0c, 0x00, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_Short_FACT(t *testing.T) {
	b := []byte{
		// FACT
		0x46, 0x41, 0x43, 0x54,
		0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_No_FACT(t *testing.T) {
	b := []byte{
		// FACT
		0x46, 0x41, 0x43, 0x54,
		0x00, 0x00, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_Short_PLST(t *testing.T) {
	b := []byte{
		// PLST
		0x50, 0x4c, 0x53, 0x54,
		0x08, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00,
		0x49, 0x44, 0x20, 0x20,
		0x01, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_No_PLST(t *testing.T) {
	b := []byte{
		// PLST
		0x50, 0x4c, 0x53, 0x54,
		0x08, 0x00, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_Short_SMPL(t *testing.T) {
	b := []byte{
		// SMPL
		0x53, 0x4d, 0x50, 0x4c,
		0x08, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x93, 0x58, 0x00, 0x00, 0x2d, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x10, 0xb1, 0x02, 0x00, 0x1f, 0x62, 0x05, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_No_SMPL(t *testing.T) {
	b := []byte{
		// SMPL
		0x53, 0x4d, 0x50, 0x4c,
		0x08, 0x00, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_Short_INST(t *testing.T) {
	b := []byte{
		// INST
		0x49, 0x4e, 0x53, 0x54,
		0x00, 0x00, 0x00, 0x00,
		0x3c, 0x00, 0x00, 0x00,
		0x7f, 0x01, 0x7f, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_No_INST(t *testing.T) {
	b := []byte{
		// INST
		0x49, 0x4e, 0x53, 0x54,
		0x00, 0x00, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_Short_CUE(t *testing.T) {
	b := []byte{
		// CUE
		0x43, 0x55, 0x45, 0x20,
		0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00,
		0xa0, 0x0f, 0x00, 0x00,
		0x64, 0x61, 0x74, 0x61,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xa0, 0x0f, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Fail_Before_Data_No_CUE(t *testing.T) {
	b := []byte{
		// CUE
		0x43, 0x55, 0x45, 0x20,
		0x00, 0x00, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)

	if wc != nil {
		t.Error("returned wav chunks is not nil")
	}

	if err == nil {
		t.Error("returned error is nil")
	}

	if err != io.EOF {
		t.Errorf("expected \"EOF\". received \"%s\"", err.Error())
	}

}

func TestReadWavChunks_Success_After_Data(t *testing.T) {
	b := []byte{
		// FMT
		0x66, 0x6d, 0x74, 0x20, 0x10, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x40, 0x1f, 0x00, 0x00, 0x80, 0x3e,
		0x00, 0x00, 0x02, 0x00, 0x10, 0x00,

		// DATA
		0x64, 0x61, 0x74, 0x61,
		0x08, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// LIST
		0x4c, 0x49, 0x53, 0x54, 0x4e, 0x00, 0x00, 0x00, 0x61,
		0x64, 0x74, 0x6c, 0x6c, 0x74, 0x78, 0x74, 0x14, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x10, 0xb1, 0x02,
		0x00, 0x72, 0x67, 0x6e, 0x20, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x6c, 0x61, 0x62, 0x6c, 0x0b,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x4c, 0x6f,
		0x6f, 0x70, 0x20, 0x31, 0x00, 0x00, 0x6e, 0x6f, 0x74,
		0x65, 0x11, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x20, 0x6c,
		0x6f, 0x6f, 0x70, 0x00, 0x00,

		// FACT
		0x46, 0x41, 0x43, 0x54,
		0x04, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00,

		// PLST
		0x50, 0x4c, 0x53, 0x54,
		0x10, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00,
		0x49, 0x44, 0x20, 0x20,
		0x01, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00,

		// SMPL
		0x53, 0x4d, 0x50, 0x4c,
		0x3c, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x93, 0x58, 0x00, 0x00, 0x2d, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x10, 0xb1, 0x02, 0x00, 0x1f, 0x62, 0x05, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

		// INST
		0x49, 0x4e, 0x53, 0x54,
		0x08, 0x00, 0x00, 0x00,
		0x3c, 0x00, 0x00, 0x00, 0x7f, 0x01, 0x7f, 0x00,

		// CUE
		0x43, 0x55, 0x45, 0x20,
		0x1c, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0xa0, 0x0f, 0x00, 0x00, 0x64, 0x61, 0x74, 0x61,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xa0, 0x0f, 0x00, 0x00,
	}

	r := bytes.NewReader(b)

	wc, err := readWavChunks(r)
	if err != nil {
		t.Error(err.Error())
	}

	if wc.FMT == nil {
		t.Error("fmt chunk is nil")
	}

	if wc.ADTL.LabeledTexts == nil {
		t.Error("adtl labeled text is nil")
	}

	if len(wc.ADTL.LabeledTexts) != 1 {
		t.Errorf("Expected 1 labeled text. received %d labeled text(s)", len(wc.ADTL.LabeledTexts))
	}

	if wc.ADTL.Labels == nil {
		t.Error("adtl labels is nil")
	}

	if len(wc.ADTL.Labels) != 1 {
		t.Errorf("Expected 1 label. received %d label(s)", len(wc.ADTL.Labels))
	}

	if wc.ADTL.Notes == nil {
		t.Error("adtl notes is nil")
	}

	if len(wc.ADTL.Notes) != 1 {
		t.Errorf("Expected 1 note. received %d labeled text(s)", len(wc.ADTL.Notes))
	}

	if len(wc.CuePoints) != 1 {
		t.Errorf("Expected 1 cue point. received %d cue point(s)", len(wc.CuePoints))
	}

	if wc.Fact == nil {
		t.Error("fact is nil")
	}

	if wc.Info != nil {
		t.Error("info is not nil")
	}

	if wc.Inst == nil {
		t.Error("inst is nil")
	}

	if len(wc.PlstSegments) != 1 {
		t.Errorf("Expected 1 plst segment. received %d plst segment(s)", len(wc.PlstSegments))
	}

	if wc.Sample == nil {
		t.Error("sample is nil")
	}

	if wc.DataPosition != 32 {
		t.Errorf("Expected data position 32. received data position %d", wc.DataPosition)
	}

	if wc.DataLength != 8 {
		t.Errorf("Expected data length of 8. received data length %d", wc.DataLength)
	}

}
