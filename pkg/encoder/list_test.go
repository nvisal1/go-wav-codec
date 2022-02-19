package Encoder

import (
	"bytes"
	"encoding/binary"
	"testing"
	"wav-concat/pkg/Decoder"
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
	lc := &ListChunk{info: ic}

	var b bytes.Buffer
	_, err := writeLISTChunk(&b, lc)
	if err != nil {
		t.Error(err.Error())
	}

	r := bytes.NewReader(b.Bytes())

	b32 := make([]byte, 4)
	_, err = r.Read(b32)
	if string(b32[:]) != LIST_CHUNK_ID {
		t.Errorf("first 4 bytes are not %s", LIST_CHUNK_ID)
	}

	_, err = r.Read(b32)
	if binary.LittleEndian.Uint32(b32[:]) != 24 {
		t.Errorf("second 4 bytes are not %d", binary.LittleEndian.Uint32(b32[:]))
	}

	_, err = r.Read(b32)
	if string(b32[:]) != INFO_CHUNK_ID {
		t.Errorf("third 4 bytes are not %s", INFO_CHUNK_ID)
	}

	dc, err := Decoder.ReadINFOChunk(r)
	if err != nil {
		t.Error(err.Error())
	}

	if dc.Location != "a cool place" {
		t.Error("decoded info location is incorrect")
	}

}
