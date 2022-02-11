package Encoder

import (
	"encoding/binary"
	"io"
)

type InstChunk struct {
	UnshiftedNote uint8
	FineTuneDB    uint8
	Gain          uint8
	LowNote       uint8
	HighNote      uint8
	LowVelocity   uint8
	HighVelocity  uint8
}

func (i InstChunk) WriteTo(w io.Writer) (int, error) {
	bytesWritten := 0

	b := Align(bytesFromString(INST_CHUNK_ID))
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = Align(bytesFromUINT32(uint32(7)))
	err = binary.Write(w, binary.LittleEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	if err := binary.Write(w, binary.LittleEndian, &i.UnshiftedNote); err != nil {
		return bytesWritten, err
	}
	bytesWritten += 1
	if err := binary.Write(w, binary.LittleEndian, &i.FineTuneDB); err != nil {
		return bytesWritten, err
	}
	bytesWritten += 1
	if err := binary.Write(w, binary.LittleEndian, &i.Gain); err != nil {
		return bytesWritten, err
	}
	bytesWritten += 1
	if err := binary.Write(w, binary.LittleEndian, &i.LowNote); err != nil {
		return bytesWritten, err
	}
	bytesWritten += 1
	if err := binary.Write(w, binary.LittleEndian, &i.HighNote); err != nil {
		return bytesWritten, err
	}
	bytesWritten += 1
	if err := binary.Write(w, binary.LittleEndian, &i.HighVelocity); err != nil {
		return bytesWritten, err
	}
	if err := binary.Write(w, binary.LittleEndian, &i.LowVelocity); err != nil {
		return bytesWritten, err
	}
	bytesWritten += 1

	return bytesWritten, nil
}
