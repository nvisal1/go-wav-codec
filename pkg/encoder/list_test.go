package encoder

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func Test_writeListChunk(t *testing.T) {
	ic := &InfoChunk{
		Location:     "a cool place",
		Artist:       "",
		Software:     "",
		CreationDate: "",
		Copyright:    "",
		Title:        "",
		Engineer:     "",
		Genre:        "",
		Product:      "",
		Source:       "",
		Subject:      "",
		Comments:     "",
		Technician:   "",
		Keywords:     "",
		Medium:       "",
	}

	var b bytes.Buffer
	_, err := writeLISTChunk(&b, ic)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())

	b32 := make([]byte, 4)
	_, err = r.Read(b32)
	if string(b32[:]) != listChunkID {
		t.Errorf("first 4 bytes are not %s", listChunkID)
	}

	_, err = r.Read(b32)
	if binary.LittleEndian.Uint32(b32[:]) != 24 {
		t.Errorf("second 4 bytes are not %d", binary.LittleEndian.Uint32(b32[:]))
	}

	_, err = r.Read(b32)
	if string(b32[:]) != infoChunkID {
		t.Errorf("third 4 bytes are not %s", infoChunkID)
	}

}
