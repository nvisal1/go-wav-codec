package Decoder

import (
	"bytes"
	"errors"
	"io"
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

type Decoder struct {
	r  io.ReadSeeker
	WC *WavChunks
}

func NewDecoder(r io.ReadSeeker) *Decoder {
	d := &Decoder{r: r}
	return d
}

//func (d *Decoder) Seek(offset int64, whence int) (int64, error) {
//	return d.r.Seek(offset, whence)
//}
//

func RecordAndForward(r io.Reader, s int) (*bytes.Reader, error) {
	b := make([]byte, s)

	if _, err := r.Read(b); err != nil {
		return nil, err
	}

	nr := bytes.NewReader(b)

	return nr, nil
}

func (d *Decoder) ReadMetadata() error {
	if d == nil {
		return nil
	}

	wfhr, err := RecordAndForward(d.r, 12)
	if err != nil {
		return err
	}

	wfh, err := ReadWavFileHeader(wfhr)
	if err != nil {
		return err
	}

	wcr, err := RecordAndForward(d.r, int(wfh.FileSize-4))
	if err != nil {
		return err
	}

	wc, err := ReadWavChunks(wcr)
	if err != nil {
		return err
	}

	d.WC = wc

	return nil
}

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
	b, err := ReadDataChunk(s, d.WC.FMT.BitsPerSample, d.r)
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
