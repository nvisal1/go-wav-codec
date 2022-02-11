package Encoder

import (
	"encoding/binary"
	"io"
)

type ADTLChunk struct {
	Labels       []*Label
	Notes        []*Note
	LabeledTexts []*LabeledText
}

func (a *ADTLChunk) WriteTo(w io.Writer) (int, error) {
	bytesWritten := 0

	b := Align(bytesFromString(ASSOCIATED_DATA_LIST_CHUNK_ID))
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	if len(a.Labels) > 0 {
		for _, lb := range a.Labels {
			n, err := lb.WriteTo(w)
			bytesWritten += n
			if err != nil {
				return bytesWritten, err
			}
		}
	}

	if len(a.Notes) > 0 {
		for _, n := range a.Notes {
			n, err := n.WriteTo(w)
			bytesWritten += n
			if err != nil {
				return bytesWritten, err
			}
		}
	}

	if len(a.LabeledTexts) > 0 {
		for _, lt := range a.LabeledTexts {
			n, err := lt.WriteTo(w)
			bytesWritten += n
			if err != nil {
				return bytesWritten, err
			}
		}
	}

	return bytesWritten, nil
}
