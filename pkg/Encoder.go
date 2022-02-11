package pkg

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type Encoder struct {
	AudioFormat   int
	NumChannels   int
	SampleRate    int
	BitsPerSample int
	F             io.WriteSeeker
	pcmStart      int64
	framesWritten int
	bytesWritten  int
}

/**
Encoder |  pcm + metadata -> wav file
Decoder |  wav file -> pcm + metadata
*/

func NewEncoder(audioFormat int, numChannels int, sampleRate int, bitsPerSample int, f io.WriteSeeker) (*Encoder, error) {
	e := &Encoder{
		AudioFormat:   audioFormat,
		NumChannels:   numChannels,
		SampleRate:    sampleRate,
		BitsPerSample: bitsPerSample,
		F:             f,
	}

	e.writeHeaders()
	return e, nil
}

func (e *Encoder) writeHeaders() error {
	i, err := e.F.Write(e.bytesFromString(RIFF_CHUNK_ID))
	if err != nil {
		return err
	}

	e.bytesWritten += i

	chunkSize := e.calculateChunkSize()
	i, err = e.F.Write(e.bytesFromUINT32(uint32(chunkSize)))
	if err != nil {
		return err
	}

	e.bytesWritten += i

	i, err = e.F.Write(e.bytesFromString(WAVE_FILE_FORMAT))
	if err != nil {
		return err
	}

	e.bytesWritten += i

	i, err = e.F.Write(e.bytesFromString(FMT_SUB_CHUNK_ID))
	if err != nil {
		return err
	}

	e.bytesWritten += i

	subChunk1Size := uint32(16)
	i, err = e.F.Write(e.bytesFromUINT32(subChunk1Size))
	if err != nil {
		return err
	}

	e.bytesWritten += i

	i, err = e.F.Write(e.bytesFromUINT16(uint16(e.AudioFormat)))
	if err != nil {
		return err
	}

	e.bytesWritten += i

	i, err = e.F.Write(e.bytesFromUINT16(uint16(e.NumChannels)))
	if err != nil {
		return err
	}

	e.bytesWritten += i

	i, err = e.F.Write(e.bytesFromUINT32(uint32(e.SampleRate)))
	if err != nil {
		return err
	}

	e.bytesWritten += i

	byteRate := e.calculateByteRate()
	i, err = e.F.Write(e.bytesFromUINT32(uint32(byteRate)))
	if err != nil {
		return err
	}

	e.bytesWritten += i

	blockAlign := e.calculateBlockAlign()
	i, err = e.F.Write(e.bytesFromUINT16(uint16(blockAlign)))
	if err != nil {
		return err
	}

	e.bytesWritten += i

	i, err = e.F.Write(e.bytesFromUINT16(uint16(e.BitsPerSample)))
	if err != nil {
		return err
	}

	e.bytesWritten += i

	e.pcmStart, err = e.F.Seek(0, 1)
	if err != nil {
		return err
	}

	i, err = e.F.Write(e.bytesFromString(DATA_SUB_CHUNK_ID))
	if err != nil {
		return err
	}

	e.bytesWritten += i

	i, err = e.F.Write(e.bytesFromUINT32(uint32(0)))
	if err != nil {
		return err
	}

	e.bytesWritten += i

	switch e.F.(type) {
	case *os.File:
		return e.F.(*os.File).Sync()
	}

	return nil
}

func (e *Encoder) writeBuffer(p []int) error {
	frameCount := e.getFrameCount(p)
	for i := 0; i < frameCount; i++ {
		for j := 0; j < e.NumChannels; j++ {
			v := p[i*e.NumChannels+j]
			switch e.BitsPerSample {
			case 8:
				if err := binary.Write(e.F, binary.LittleEndian, uint8(v)); err != nil {
					return err
				}
				e.bytesWritten += 1
			case 16:
				if err := binary.Write(e.F, binary.LittleEndian, int16(v)); err != nil {
					return err
				}
				e.bytesWritten += 2
			case 24:
				if err := binary.Write(e.F, binary.LittleEndian, Int32toInt24LEBytes(int32(v))); err != nil {
					return err
				}
				e.bytesWritten += 3
			case 32:
				if err := binary.Write(e.F, binary.LittleEndian, int32(v)); err != nil {
					return err
				}
				e.bytesWritten += 4
			default:
				return fmt.Errorf("can't add frames of bit size %d", e.BitsPerSample)
			}
		}

		e.framesWritten += 1
	}
	return nil
}

func (e *Encoder) Write(p []int) error {
	if e.pcmStart == 0 {
		e.writeHeaders()
	}

	if err := e.writeBuffer(p); err != nil {
		return err
	}

	return nil
}

func (e *Encoder) calculateChunkSize() int {
	return 0
}

//  == SampleRate * NumChannels * BitsPerSample/8
func (e *Encoder) calculateByteRate() int {
	return e.SampleRate * e.NumChannels * (e.BitsPerSample / 8)
}

// == NumChannels * BitsPerSample/8
func (e *Encoder) calculateBlockAlign() int {
	return e.NumChannels * (e.BitsPerSample / 8)
}

func (e *Encoder) calculateDataSubChunkSize() int {
	fmt.Println(e.BitsPerSample)
	fmt.Println(e.NumChannels)
	fmt.Println(e.framesWritten)
	return (int(e.BitsPerSample) / 8) * e.NumChannels * e.framesWritten
}

func (e *Encoder) getFrameCount(p []int) int {
	return len(p) / e.NumChannels // length of data / number of channels
}

func (e *Encoder) bytesFromUINT16(d uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, d)
	return b
}

func Int32toInt24LEBytes(n int32) []byte {
	bytes := make([]byte, 3)
	if (n & 0x800000) > 0 {
		n |= ^0xffffff
	}
	bytes[2] = byte(n >> 16)
	bytes[1] = byte(n >> 8)
	bytes[0] = byte(n >> 0)
	return bytes
}

func (e *Encoder) bytesFromUINT32(d uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, d)
	return b
}

func (e *Encoder) bytesFromString(d string) []byte {
	return []byte(d)
}

func (e *Encoder) ToDistortion(p []int) []int {
	//upperThresh := 32768
	upperThresh := 100
	//lowerThresh := -32768
	lowerThresh := -100
	for i := 0; i < len(p); i++ {
		if p[i] > upperThresh {
			p[i] = upperThresh
		}

		if p[i] < lowerThresh {
			p[i] = lowerThresh
		}
	}

	return p
}

func (e *Encoder) ToDelay(p []int) []int {
	delayMax := 5000
	c := 0
	buf := make([]int, delayMax)
	for i := 0; i < len(p); i++ {
		buf[c] = p[i]
		c++
		if c >= delayMax {
			c = 0
		}

		p[i] = buf[c] + p[i]

	}
	return p
}

func (e *Encoder) Close() error {
	// go back and write total size in header
	if _, err := e.F.Seek(4, 0); err != nil {
		return err
	}
	if _, err := e.F.Write(e.bytesFromUINT32(uint32(e.bytesWritten - 8))); err != nil {
		return fmt.Errorf("%v when writing the total written bytes", err)
	}

	// rewrite the audio chunk length header
	if e.pcmStart > 0 {
		if _, err := e.F.Seek(e.pcmStart+4, 0); err != nil {
			return err
		}
		chunkSize := uint32(e.calculateDataSubChunkSize())
		if _, err := e.F.Write(e.bytesFromUINT32(chunkSize)); err != nil {
			return fmt.Errorf("%v when writing wav data chunk size header", err)
		}
	}

	// jump back to the end of the file.
	if _, err := e.F.Seek(0, 2); err != nil {
		return err
	}
	switch e.F.(type) {
	case *os.File:
		return e.F.(*os.File).Sync()
	}

	return nil
}
