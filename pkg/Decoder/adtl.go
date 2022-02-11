package Decoder

import (
	"bytes"
	"io"
)

type ADTLChunk struct {
	Labels       []*Label
	Notes        []*Note
	LabeledTexts []*LabeledText
}

func ReadADTLChunk(r *bytes.Reader) (*ADTLChunk, error) {
	ac := &ADTLChunk{}
	for {
		c, err := NewChunk(r)
		if err == io.EOF {
			return ac, nil
		}
		if err != nil {
			return nil, err
		}

		switch c.ID {
		case LABL_CHUNK_ID:
			l, err := ReadLABLChunk(r, c.Size-4)
			if err != nil {
				return nil, err
			}
			ac.Labels = append(ac.Labels, l)
		case NOTE_CHUNK_ID:
			n, err := ReadNoteChunk(r, c.Size-4)
			if err != nil {
				return nil, err
			}
			ac.Notes = append(ac.Notes, n)
		case TEXT_CHUNK_ID:
			lt, err := ReadLTXTChunk(r, c.Size-4)
			if err != nil {
				return nil, err
			}
			ac.LabeledTexts = append(ac.LabeledTexts, lt)
		}
	}
}
