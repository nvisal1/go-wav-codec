package Encoder

import (
	"encoding/binary"
	"io"
)

func WriteWavFileHeader(w io.Writer) (int, error) {
	bytesWritten := 0

	b := bytesFromString(RIFF_CHUNK_ID)
	err := binary.Write(w, binary.BigEndian, b)
	if err != nil {
		return bytesWritten, err
	}

	bytesWritten += len(b)

	// We put a placeholder 0 here to init the file
	// This value is updated right before the file
	// is closed
	b = bytesFromUINT32(uint32(0))
	err = binary.Write(w, binary.LittleEndian, b)
	if err != nil {
		return bytesWritten, err
	}

	bytesWritten += len(b)

	b = bytesFromString(WAVE_FILE_FORMAT)
	err = binary.Write(w, binary.BigEndian, b)
	if err != nil {
		return bytesWritten, err
	}

	bytesWritten += len(b)

	return bytesWritten, nil
}
