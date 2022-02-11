package Encoder

import (
	"encoding/binary"
	"io"
)

type Label struct {
	CuePointID string
	Data       string
}

func (l Label) WriteTo(w io.Writer) (int, error) {
	bytesWritten := 0

	b := Align(bytesFromString(LABL_CHUNK_ID))
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = Align(bytesFromUINT32(uint32(l.getChunkSize())))
	err = binary.Write(w, binary.LittleEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = Align(bytesFromString(l.CuePointID))
	if err := binary.Write(w, binary.BigEndian, &b); err != nil {
		return bytesWritten, nil
	}
	bytesWritten += len(b)

	b = Align(bytesFromString(l.Data))
	if err := binary.Write(w, binary.BigEndian, &b); err != nil {
		return bytesWritten, nil
	}
	bytesWritten += len(b)

	return bytesWritten, nil
}

func (l Label) getChunkSize() int {
	b := bytesFromString(l.Data)
	return len(b) + 8
}
