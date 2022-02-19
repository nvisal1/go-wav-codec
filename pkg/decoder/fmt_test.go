package Decoder

import (
	"bytes"
	"testing"
)

func TestReadFMTChunk(t *testing.T) {
	b := []byte{
		0x01, 0x00, // AudioFormat
		0x02, 0x00, // NumChannels
		0x44, 0xac, 0x00, 0x00, // SampleRate
		0x98, 0x09, 0x04, 0x00, // ByteRate
		0x06, 0x00, // BlockAlign
		0x18, 0x00, // BitsPerSample
	}

	r := bytes.NewReader(b)

	f, err := ReadFMTChunk(r)
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
