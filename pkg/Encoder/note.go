package Encoder

import (
	"encoding/binary"
	"io"
)

type Note struct {
	CuePointID string
	Data       string
}

func (n Note) WriteTo(w io.Writer) (int, error) {
	bytesWritten := 0

	b := bytesFromString(NOTE_CHUNK_ID)
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = Align(bytesFromUINT32(uint32(n.getChunkSize())))
	err = binary.Write(w, binary.LittleEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromString(n.CuePointID)
	if err := binary.Write(w, binary.BigEndian, &b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)
	b = bytesFromString(n.Data)
	if err := binary.Write(w, binary.BigEndian, &b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	return bytesWritten, nil
}

func (n Note) getChunkSize() int {
	b := bytesFromString(n.Data)
	return len(b) + 8
}
