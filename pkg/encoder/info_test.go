package encoder

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestInfoChunk_WriteTo(t *testing.T) {
	ic := &InfoChunk{
		Location:     "AAA",
		Artist:       "BBBB",
		Software:     "C",
		CreationDate: "DD",
		Copyright:    "EEEEE",
		Title:        "FFFFFF",
		Engineer:     "GGGGGGG",
		Genre:        "",
		Product:      "HHHHHHHH",
		Source:       "I",
		Subject:      "JJ",
		Comments:     "K",
		Technician:   "LLLL",
		Keywords:     "",
		Medium:       "",
	}

	var b bytes.Buffer
	_, err := writeInfoChunk(&b, ic)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())
	b32 := make([]byte, 4)
	_, err = r.Read(b32)

	if string(b32) != infoChunkID {
		t.Errorf("first 4 bytes are not %s", infoChunkID)
	}

	_, err = r.Read(b32)

	if string(b32) != iARL {
		t.Errorf("next 4 bytes are not %s", iARL)
	}

	_, err = r.Read(b32)

	if binary.LittleEndian.Uint32(b32) != 4 {
		t.Errorf("next 4 bytes are not %d", 4)
	}

	_, err = r.Read(b32)

	if string(b32) != "AAA\u0000" {
		t.Errorf("next 4 bytes are not %s", "AAA\u0000")
	}

}
