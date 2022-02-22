package decoder

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

func calculateBytesPerSample(bitsPerSample uint16) uint16 {
	return bitsPerSample / 8
}

func readDataChunk(numSamples int, bitsPerSample uint16, r io.Reader) ([]int, error) {
	if numSamples <= 0 {
		return nil, errors.New("numSamples must be 1 or greater")
	}
	bytesPerSample := calculateBytesPerSample(bitsPerSample)
	sampleBuf := make([]byte, bytesPerSample)
	outBuf := make([]int, numSamples)

	f, err := getSampleDecodeFunc(int(bitsPerSample), int(bytesPerSample))
	if err != nil {
		return nil, err
	}

	i := 0
	for {
		outBuf[i], err = f(r, sampleBuf)
		if err != nil {
			if err == io.EOF {
				return outBuf, err
			}
			return nil, err
		}
		i++
		if i == numSamples {
			break
		}
	}

	return outBuf, nil
}

// getSampleDecodeFunc returns a function that can be used to convert
// a byte range into an int value based on the amount of bits used per sample.
// note that 8bit samples are unsigned, all other values are signed.
// note that pcm data is stored using little-endian
func getSampleDecodeFunc(bitsPerSample int, bytesPerSample int) (func(io.Reader, []byte) (int, error), error) {

	switch bitsPerSample {
	case 8:
		return func(r io.Reader, b []byte) (int, error) {
			err := binary.Read(r, binary.LittleEndian, b[:bytesPerSample])
			out := int(b[bytesPerSample-1])
			return out, err
		}, nil
	case 16:
		return func(r io.Reader, b []byte) (int, error) {
			err := binary.Read(r, binary.LittleEndian, b[:bytesPerSample])
			out := int(int16(binary.LittleEndian.Uint16(b[:bytesPerSample])))
			return out, err
		}, nil
	case 32:
		return func(r io.Reader, b []byte) (int, error) {
			err := binary.Read(r, binary.LittleEndian, b[:bytesPerSample])
			out := int(int32(binary.LittleEndian.Uint32(b[:bytesPerSample])))
			return out, err
		}, nil
	default:
		return nil, fmt.Errorf("received unsupported bits per sample: %d | supported bits per sample: [8, 16, 32]", bitsPerSample)
	}
}
