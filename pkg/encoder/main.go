package encoder

import (
	"fmt"
	"io"
	"os"
)

const (
	riffChunkID    = "RIFF"
	waveFileFormat = "WAVE"
	fmtChunkID     = "fmt "
	listChunkID    = "LIST"
	infoChunkID    = "INFO"
	dataChunkID    = "data"
)

// Encoder is used for writing new wav files
type Encoder struct {
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	BitsPerSample uint16
	F             io.WriteSeeker
	dataStart     int64
	dataOffset    int64
	framesWritten int
	bytesWritten  int
	infoChunk     *InfoChunk
}

// NewEncoder returns an Encoder
func NewEncoder(audioFormat uint16, numChannels uint16, sampleRate uint32, bitsPerSample uint16, f io.WriteSeeker) (*Encoder, error) {
	e := &Encoder{
		AudioFormat:   audioFormat,
		NumChannels:   numChannels,
		SampleRate:    sampleRate,
		BitsPerSample: bitsPerSample,
		F:             f,
	}

	err := e.writeHeadersAndFMT()
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e *Encoder) writeHeadersAndFMT() error {
	n, err := writeWavFileHeader(e.F)
	if err != nil {
		return err
	}
	e.bytesWritten += n

	f := &fmtChunk{
		AudioFormat:   e.AudioFormat,
		NumChannels:   e.NumChannels,
		SampleRate:    e.SampleRate,
		BitsPerSample: e.BitsPerSample,
	}

	n, err = writeFMTChunk(e.F, f)
	if err != nil {
		return err
	}
	e.bytesWritten += n

	switch e.F.(type) {
	case *os.File:
		return e.F.(*os.File).Sync()
	}

	return nil
}

// WriteMetadata accepts a *ListChunk and writes it to the
// end of the file. The LIST INFO chunk will always
// appear after the audio data.
// This will not immediately write the LIST chunk to the wav file
// Instead, it will store the *ListChunk and write it
// at the end of the file when Encoder.Close is called.
// If this function is called multiple times, only
// the most recent metadata will be written to the file
// when Encoder.Close is called. This encoder does not
// support multiple LIST chunks in a single wav file.
func (e *Encoder) WriteMetadata(i *InfoChunk) {
	e.infoChunk = i
}

func (e *Encoder) writeAudioDataHeader() error {
	pos, err := e.F.Seek(0, 1)
	if err != nil {
		return err
	}

	e.dataStart = pos
	e.dataOffset = pos

	n, err := writeDataChunkID(e.F)
	if err != nil {
		return err
	}

	e.bytesWritten += n

	switch e.F.(type) {
	case *os.File:
		return e.F.(*os.File).Sync()
	}

	return nil
}

// WriteAudioData accepts a buffer of audio data and writes it to the
// wav file. This can be called multiple times. The first call to this function
// will write the data chunk ID to the file. Every subsequent call will
// write to the end of the existing audio data.
// whence can be either 0 or 1
// whence 0: write the provided audio data at the start of the
// data chunk. This will overwrite any existing audio data that already
// exists in that position
// whence 1: write the provided audio data after the existing audio data
func (e *Encoder) WriteAudioData(p []int, whence int) error {
	if e.dataStart == 0 {
		e.writeAudioDataHeader()
	}

	if whence == 0 {
		e.F.Seek(e.dataStart+8, 0)
	} else if whence == 1 {
		e.F.Seek(e.dataOffset, 0)
	}

	n, f, err := writeDataChunkBuffer(e.F, p, e.NumChannels, e.BitsPerSample)
	e.bytesWritten += n
	e.framesWritten += f
	e.dataOffset, err = e.F.Seek(0, 1)
	if err != nil {
		return err
	}

	return nil
}

// Close should always be called before using the created wav file
// it is responsible for writing the LIST INFO chunk (if provided)
// and updating the headers with the correct byte sizes
func (e *Encoder) Close() error {
	if e.infoChunk != nil {
		// go to the end of the file.
		if _, err := e.F.Seek(0, 2); err != nil {
			return err
		}

		n, err := writeLISTChunk(e.F, e.infoChunk)
		if err != nil {
			return err
		}
		e.bytesWritten += n
	}

	// go back and write total size in header
	if _, err := e.F.Seek(4, 0); err != nil {
		return err
	}
	if _, err := e.F.Write(bytesFromUINT32(uint32(e.bytesWritten - 8 + 4))); err != nil {
		return fmt.Errorf("%v when writing the total written bytes", err)
	}

	// rewrite the audio chunk length header
	if e.dataStart > 0 {
		if _, err := e.F.Seek(e.dataStart+4, 0); err != nil {
			return err
		}
		chunkSize := uint32(calculateDataChunkSize(e.NumChannels, e.BitsPerSample, e.framesWritten))
		if _, err := e.F.Write(bytesFromUINT32(chunkSize)); err != nil {
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
