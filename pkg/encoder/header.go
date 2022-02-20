package encoder

import (
	"encoding/binary"
	"io"
)

func writeWavFileHeader(w io.Writer) (int, error) {
	bytesWritten := 0

	b := bytesFromString(riffChunkID)
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

	b = bytesFromString(waveFileFormat)
	err = binary.Write(w, binary.BigEndian, b)
	if err != nil {
		return bytesWritten, err
	}

	bytesWritten += len(b)

	return bytesWritten, nil
}
