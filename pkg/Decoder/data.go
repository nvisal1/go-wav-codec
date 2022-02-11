package Decoder

import (
	"encoding/binary"
	"fmt"
	"io"
)

func calculateBytesPerSample(bitsps uint16) uint16 {
	return (bitsps-1)/8 + 1
}

func ReadDataChunk(s int, bitsps uint16, r io.Reader) ([]int, error) {
	bytesPerSample := calculateBytesPerSample(bitsps)
	sampleBuf := make([]byte, bytesPerSample)
	outBuf := make([]int, s)

	f, err := sampleDecodeFunc(int(bitsps))
	if err != nil {
		return nil, err
	}

	i := 0
	for {
		outBuf[i], err = f(r, sampleBuf)
		if err != nil {
			return nil, err
		}
		i++
		if i == s {
			break
		}
	}

	return outBuf, nil
}

// sampleDecodeFunc returns a function that can be used to convert
// a byte range into an int value based on the amount of bits used per sample.
// Note that 8bit samples are unsigned, all other values are signed.
// Note: Credit - I found this elegant implementation here https://github.com/go-audio/wav
func sampleDecodeFunc(bitsPerSample int) (func(io.Reader, []byte) (int, error), error) {
	// NOTE: WAV PCM data is stored using little-endian
	switch bitsPerSample {
	case 8:
		// 8bit values are unsigned
		return func(r io.Reader, buf []byte) (int, error) {
			_, err := r.Read(buf[:1])
			return int(buf[0]), err
		}, nil
	case 16:
		return func(r io.Reader, buf []byte) (int, error) {
			_, err := r.Read(buf[:2])
			return int(int16(binary.LittleEndian.Uint16(buf[:2]))), err
		}, nil
	case 32:
		return func(r io.Reader, buf []byte) (int, error) {
			_, err := r.Read(buf[:4])
			return int(int32(binary.LittleEndian.Uint32(buf[:4]))), err
		}, nil
	default:
		return nil, fmt.Errorf("unhandled byte depth:%d", bitsPerSample)
	}
}
