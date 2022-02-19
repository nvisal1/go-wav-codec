package Encoder

import (
	"bytes"
	"testing"
	"wav-concat/pkg/Decoder"
)

func TestWriteFMTChunk(t *testing.T) {
	fc := &FMTChunk{
		AudioFormat:   0,
		NumChannels:   0,
		SampleRate:    0,
		BitsPerSample: 0,
	}

	var b bytes.Buffer
	_, err := WriteFMTChunk(&b, fc)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	b32 := make([]byte, 4)
	_, err = r.Read(b32)

	if string(b32[:]) != FMT_CHUNK_ID {
		t.Errorf("First 4 bytes is not %s", FMT_CHUNK_ID)
	}

	_, err = r.Read(b32)

	dfc, err := Decoder.ReadFMTChunk(r)
	if err != nil {
		t.Error(err.Error())
	}

	if dfc.AudioFormat != 0 {
		t.Error("fmt audio format is incorrect")
	}

	if dfc.NumChannels != 0 {
		t.Error("fmt num channels is incorrect")
	}

	if dfc.SampleRate != 0 {
		t.Error("fmt sample rate is incorrect")
	}

	if dfc.ByteRate != 0 {
		t.Error("fmt byte rate is incorrect")
	}

	if dfc.BlockAlign != 0 {
		t.Error("fmt block align is incorrect")
	}

	if dfc.BitsPerSample != 0 {
		t.Error("fmt bits per sample is incorrect")
	}
}
