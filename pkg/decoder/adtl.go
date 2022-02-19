package decoder

import (
	"bytes"
	"io"
)

type adtlChunk struct {
	Labels       []*label
	Notes        []*note
	LabeledTexts []*labeledText
}

func readADTLChunk(r *bytes.Reader) (*adtlChunk, error) {
	ac := &adtlChunk{}
	for {
		c, err := newChunk(r)
		if err != nil {
			if err == io.EOF {
				return ac, nil
			}
			return nil, err
		}

		// If the size in the header is not an even number,
		// we need to read an extra byte in order to stay
		// word aligned
		if c.Size%2 != 0 {
			c.Size++
		}

		switch c.ID {
		case lablChunkID:
			l, err := readLABLChunk(r, c.Size)
			if err != nil {
				return nil, err
			}
			ac.Labels = append(ac.Labels, l)
		case noteChunkID:
			n, err := readNoteChunk(r, c.Size)
			if err != nil {
				return nil, err
			}
			ac.Notes = append(ac.Notes, n)
		case textChunkID:
			lt, err := readLTXTChunk(r, c.Size)
			if err != nil {
				return nil, err
			}
			ac.LabeledTexts = append(ac.LabeledTexts, lt)
		}
	}
}
