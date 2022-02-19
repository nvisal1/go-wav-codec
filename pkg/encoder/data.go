package Encoder

import (
	"encoding/binary"
	"fmt"
	"io"
)

func WriteDataChunkBuffer(w io.Writer, p []int, numChannels uint16, bitsPerSample uint16) (int, int, error) {
	bytesWritten := 0
	framesWritten := 0
	frameCount := calculateFrameCount(p, numChannels)
	for i := 0; i < frameCount; i++ {
		for j := 0; j < int(numChannels); j++ {
			v := p[i*int(numChannels)+j]
			switch bitsPerSample {
			case 8:
				if err := binary.Write(w, binary.LittleEndian, uint8(v)); err != nil {
					return bytesWritten, framesWritten, err
				}
				bytesWritten += 1
			case 16:
				if err := binary.Write(w, binary.LittleEndian, int16(v)); err != nil {
					return bytesWritten, framesWritten, err
				}
				bytesWritten += 2
			case 32:
				if err := binary.Write(w, binary.LittleEndian, int32(v)); err != nil {
					return bytesWritten, framesWritten, err
				}
				bytesWritten += 4
			default:
				return bytesWritten, framesWritten, fmt.Errorf("can't add frames of bit size %d", bitsPerSample)
			}
		}
		framesWritten += 1
	}
	return bytesWritten, framesWritten, nil
}

func WriteDataChunkID(w io.Writer) (int, error) {
	bytesWritten := 0
	b := bytesFromString(DATA_CHUNK_ID)
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten = len(b)

	b = bytesFromUINT32(uint32(0))
	err = binary.Write(w, binary.LittleEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten = len(b)

	return bytesWritten, nil
}
