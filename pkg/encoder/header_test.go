package Encoder

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestWriteWavFileHeader(t *testing.T) {
	var b bytes.Buffer
	_, err := WriteWavFileHeader(&b)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	b32 := make([]byte, 4)
	_, err = r.Read(b32)

	if string(b32[:]) != RIFF_CHUNK_ID {
		t.Errorf("first 4 bytes are not %s", RIFF_CHUNK_ID)
	}

	_, err = r.Read(b32)

	if binary.LittleEndian.Uint32(b32) != 0 {
		t.Error("bytes 4-8 are not 0")
	}

	_, err = r.Read(b32)

	if string(b32[:]) != WAVE_FILE_FORMAT {
		t.Errorf("bytes 9-12 are not %s", WAVE_FILE_FORMAT)
	}
}
