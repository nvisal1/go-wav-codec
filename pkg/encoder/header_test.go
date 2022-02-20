package encoder

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestWriteWavFileHeader(t *testing.T) {
	var b bytes.Buffer
	_, err := writeWavFileHeader(&b)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	b32 := make([]byte, 4)
	_, err = r.Read(b32)

	if string(b32[:]) != riffChunkID {
		t.Errorf("first 4 bytes are not %s", riffChunkID)
	}

	_, err = r.Read(b32)

	if binary.LittleEndian.Uint32(b32) != 0 {
		t.Error("bytes 4-8 are not 0")
	}

	_, err = r.Read(b32)

	if string(b32[:]) != waveFileFormat {
		t.Errorf("bytes 9-12 are not %s", waveFileFormat)
	}
}
