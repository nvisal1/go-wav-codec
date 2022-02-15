package Encoder

import (
	"fmt"
	"io"
	"os"
)

const (
	RIFF_CHUNK_ID                 = "RIFF"
	WAVE_FILE_FORMAT              = "WAVE"
	FMT_CHUNK_ID                  = "fmt "
	LIST_CHUNK_ID                 = "LIST"
	ASSOCIATED_DATA_LIST_CHUNK_ID = "adtl"
	INFO_CHUNK_ID                 = "INFO"
	LABL_CHUNK_ID                 = "labl"
	NOTE_CHUNK_ID                 = "note"
	TEXT_CHUNK_ID                 = "ltxt"
	SMPL_CHUNK_ID                 = "smpl"
	FACT_CHUNK_ID                 = "fact"
	PLST_CHUNK_ID                 = "plst"
	CUE_CHUNK_ID                  = "cue "
	INST_CHUNK_ID                 = "inst"
	DATA_CHUNK_ID                 = "data"
)

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
}

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
	n, err := WriteWavFileHeader(e.F)
	if err != nil {
		return err
	}
	e.bytesWritten += n

	f := &FMTChunk{
		AudioFormat:   e.AudioFormat,
		NumChannels:   e.NumChannels,
		SampleRate:    e.SampleRate,
		BitsPerSample: e.BitsPerSample,
	}

	n, err = WriteFMTChunk(e.F, f)
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

func (e *Encoder) WriteMetadata(c Chunk) error {
	n, err := c.WriteTo(e.F)
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

func (e *Encoder) WriteAudioDataHeader() error {
	pos, err := e.F.Seek(0, 1)
	if err != nil {
		return err
	}

	e.dataStart = pos
	e.dataOffset = pos

	n, err := WriteDataChunkID(e.F)
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

func (e *Encoder) WriteAudioData(p []int, whence int) error {
	if e.dataStart == 0 {
		e.WriteAudioDataHeader()
	}

	if whence == 0 {
		e.F.Seek(e.dataStart+8, 0)
	} else if whence == 1 {
		e.F.Seek(e.dataOffset, 0)
	}

	n, f, err := WriteDataChunkBuffer(e.F, p, e.NumChannels, e.BitsPerSample)
	e.bytesWritten += n
	e.framesWritten += f
	e.dataOffset, err = e.F.Seek(0, 1)
	if err != nil {
		return err
	}

	return nil
}

func (e *Encoder) Close() error {
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
