package decoder

import (
	"bytes"
	"errors"
	"io"
)

const (
	riffChunkID               = "RIFF"
	waveFileFormat            = "WAVE"
	fmtChunkID                = "fmt "
	listChunkID               = "LIST"
	associatedDataListChunkID = "adtl"
	infoChunkID               = "INFO"
	lablChunkID               = "labl"
	noteChunkID               = "note"
	textChunkID               = "ltxt"
	smplChunkID               = "smpl"
	factChunkID               = "fact"
	plstChunkID               = "plst"
	cueChunkID                = "cue "
	instChunkID               = "inst"
	dataChunkID               = "data"
)

// Decoder is used for reading a wav file
type Decoder struct {
	r  io.ReadSeeker
	WC *wavChunks
}

// NewDecoder returns a new Decoder for the provided reader
func NewDecoder(r io.ReadSeeker) *Decoder {
	d := &Decoder{r: r}
	return d
}

func recordAndForward(r io.Reader, s int) (*bytes.Reader, error) {
	b := make([]byte, s)

	if _, err := r.Read(b); err != nil {
		return nil, err
	}

	nr := bytes.NewReader(b)

	return nr, nil
}

// ReadMetadata reads all the file headers and skips the audio data
// The header information will be stored in Decoder.WC
func (d *Decoder) ReadMetadata() error {
	if d == nil {
		return nil
	}

	wfhr, err := recordAndForward(d.r, 12)
	if err != nil {
		return err
	}

	wfh, err := readWavFileHeader(wfhr)
	if err != nil {
		return err
	}

	wcr, err := recordAndForward(d.r, int(wfh.FileSize-4))
	if err != nil {
		return err
	}

	wc, err := readWavChunks(wcr)
	if err != nil {
		return err
	}

	d.WC = wc

	return nil
}

// ReadAudioData fills the given buffer with audio data
func (d *Decoder) ReadAudioData(s int, whence int) ([]int, error) {
	if d == nil {
		return nil, errors.New("The Decoder is not set")
	}

	if whence == 0 {
		err := d.toDataStart()
		if err != nil {
			return nil, err
		}
	}

	// This should be called after `toDataStart` because
	// it assumes that the FMT chunk is set on the `Decoder`
	b, err := readDataChunk(s, d.WC.FMT.BitsPerSample, d.r)
	if err != nil {
		if err == io.EOF {
			return b, err
		}
		return nil, err
	}

	return b, nil
}

func (d *Decoder) toDataStart() error {
	if d.WC.DataPosition != 0 {
		_, err := d.r.Seek(d.WC.DataPosition+12, 0)
		if err != nil {
			return err
		}
		return nil
	}

	_, err := d.r.Seek(0, 0)
	if err != nil {
		return err
	}
	err = d.ReadMetadata()
	if err != nil {
		return err
	}

	_, err = d.r.Seek(d.WC.DataPosition+12, 0)
	if err != nil {
		return err
	}

	return nil
}
