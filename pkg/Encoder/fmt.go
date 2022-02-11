package Encoder

import (
	"encoding/binary"
	"io"
)

type FMTChunk struct {
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	BitsPerSample uint16
}

func WriteFMTChunk(w io.Writer, f *FMTChunk) (int, error) {
	bytesWritten := 0
	b := bytesFromString(FMT_CHUNK_ID)
	err := binary.Write(w, binary.BigEndian, b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromUINT32(uint32(16))
	err = binary.Write(w, binary.LittleEndian, b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromUINT16(uint16(f.AudioFormat))
	err = binary.Write(w, binary.LittleEndian, b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromUINT16(uint16(f.NumChannels))
	err = binary.Write(w, binary.LittleEndian, b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromUINT32(f.SampleRate)
	err = binary.Write(w, binary.LittleEndian, b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	byteRate := calculateByteRate(f.SampleRate, f.NumChannels, f.BitsPerSample)
	b = bytesFromUINT32(byteRate)
	err = binary.Write(w, binary.LittleEndian, b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	blockAlign := calculateBlockAlign(f.NumChannels, f.BitsPerSample)
	b = bytesFromUINT16(blockAlign)
	err = binary.Write(w, binary.LittleEndian, b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromUINT16(f.BitsPerSample)
	err = binary.Write(w, binary.LittleEndian, b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	return bytesWritten, nil
}
