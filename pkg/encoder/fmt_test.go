package encoder

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestWriteFMTChunk(t *testing.T) {
	fc := &fmtChunk{
		AudioFormat:   0,
		NumChannels:   0,
		SampleRate:    0,
		BitsPerSample: 0,
	}

	var b bytes.Buffer
	_, err := writeFMTChunk(&b, fc)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	b32 := make([]byte, 4)
	_, err = r.Read(b32)

	if string(b32[:]) != fmtChunkID {
		t.Errorf("First 4 bytes is not %s", fmtChunkID)
	}

	_, err = r.Read(b32)

	b16 := make([]byte, 2)

	if err := binary.Read(r, binary.LittleEndian, &b16); err != nil || binary.LittleEndian.Uint16(b16) != 0 {
		t.Error("fmt audio format is incorrect")
	}

	if err := binary.Read(r, binary.LittleEndian, &b16); err != nil || binary.LittleEndian.Uint16(b16) != 0 {
		t.Error("fmt num channels is incorrect")
	}

	if err := binary.Read(r, binary.LittleEndian, &b32); err != nil || binary.LittleEndian.Uint32(b32) != 0 {
		t.Error("fmt sample rate is incorrect")
	}

	if err := binary.Read(r, binary.LittleEndian, &b32); err != nil || binary.LittleEndian.Uint32(b32) != 0 {
		t.Error("fmt byte rate is incorrect")
	}

	if err := binary.Read(r, binary.LittleEndian, &b16); err != nil || binary.LittleEndian.Uint16(b16) != 0 {
		t.Error("fmt block align is incorrect")
	}

	if err := binary.Read(r, binary.LittleEndian, &b16); err != nil || binary.LittleEndian.Uint16(b16) != 0 {
		t.Error("fmt bits per sample is incorrect")
	}

}
