package encoder

import (
	"encoding/binary"
	"fmt"
	"io"
)

func writeDataChunkBuffer(w io.Writer, p []int, numChannels uint16, bitsPerSample uint16) (int, int, error) {
	bytesWritten := 0
	framesWritten := 0
	frameCount := calculateFrameCount(p, numChannels)

	encodeSample, err := getSampleEncodeFunc(bitsPerSample)
	if err != nil {
		return bytesWritten, framesWritten, err
	}

	for i := 0; i < frameCount; i++ {
		for j := 0; j < int(numChannels); j++ {

			v := p[i*int(numChannels)+j]
			n, e := encodeSample(w, v)
			if e != nil {
				return bytesWritten, framesWritten, e
			}
			bytesWritten += n
		}

		framesWritten++
	}

	return bytesWritten, framesWritten, nil
}

func getSampleEncodeFunc(bitsPerSample uint16) (func(io.Writer, int) (int, error), error) {
	switch bitsPerSample {
	case 8:
		return func(w io.Writer, v int) (int, error) {
			if err := binary.Write(w, binary.LittleEndian, uint8(v)); err != nil {
				return 0, err
			}
			return 1, nil
		}, nil
	case 16:
		return func(w io.Writer, v int) (int, error) {
			if err := binary.Write(w, binary.LittleEndian, int16(v)); err != nil {
				return 0, err
			}
			return 2, nil
		}, nil
	case 32:
		return func(w io.Writer, v int) (int, error) {
			if err := binary.Write(w, binary.LittleEndian, int32(v)); err != nil {
				return 0, err
			}
			return 4, nil
		}, nil
	default:
		return nil, fmt.Errorf("unsupported bits per sample. expected one of [8, 16, 32]. received %d", bitsPerSample)
	}
}

func writeDataChunkID(w io.Writer) (int, error) {
	bytesWritten := 0
	b := bytesFromString(dataChunkID)
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
