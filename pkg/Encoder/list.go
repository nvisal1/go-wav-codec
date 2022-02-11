package Encoder

import (
	"encoding/binary"
	"errors"
	"io"
)

type ListChunk struct {
	info *InfoChunk
	adtl *ADTLChunk
}

func (l ListChunk) WriteTo(w io.Writer) (int, error) {
	bytesWritten := 0

	b := bytesFromString(LIST_CHUNK_ID)
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	if l.info != nil && l.adtl != nil {
		return bytesWritten, errors.New("Can only add one of INFO or ADTL in the LIST chunk")
	}

	if l.info != nil {
		s := l.getINFOChunkSize()
		b = bytesFromUINT32(uint32(s))
		err = binary.Write(w, binary.LittleEndian, &b)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)

		n, err := l.info.WriteTo(w)
		bytesWritten += n
		if err != nil {
			return bytesWritten, err
		}
	}

	if l.adtl != nil {
		s := l.getADTLChunkSize()
		b = bytesFromUINT32(uint32(s))
		err = binary.Write(w, binary.LittleEndian, &b)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)

		n, err := l.adtl.WriteTo(w)
		bytesWritten += n
		if err != nil {
			return bytesWritten, err
		}
	}

	return bytesWritten, nil
}

func (l ListChunk) getINFOChunkSize() int {
	total := 0

	if l.info.Location != "" {
		b := bytesFromString(l.info.Location)
		total += len(b)
	}

	if l.info.Artist != "" {
		b := bytesFromString(l.info.Artist)
		total += len(b)
	}

	if l.info.Software != "" {
		b := bytesFromString(l.info.Software)
		total += len(b)
	}

	if l.info.CreationDate != "" {
		b := bytesFromString(l.info.CreationDate)
		total += len(b)
	}

	if l.info.Copyright != "" {
		b := bytesFromString(l.info.Copyright)
		total += len(b)
	}

	if l.info.Title != "" {
		b := bytesFromString(l.info.Title)
		total += len(b)
	}

	if l.info.Engineer != "" {
		b := bytesFromString(l.info.Engineer)
		total += len(b)
	}

	if l.info.Genre != "" {
		b := bytesFromString(l.info.Genre)
		total += len(b)
	}

	if l.info.Product != "" {
		b := bytesFromString(l.info.Product)
		total += len(b)
	}

	if l.info.Source != "" {
		b := bytesFromString(l.info.Source)
		total += len(b)
	}

	if l.info.Subject != "" {
		b := bytesFromString(l.info.Subject)
		total += len(b)
	}

	if l.info.Comments != "" {
		b := bytesFromString(l.info.Comments)
		total += len(b)
	}

	if l.info.Technician != "" {
		b := bytesFromString(l.info.Technician)
		total += len(b)
	}

	if l.info.Keywords != "" {
		b := bytesFromString(l.info.Keywords)
		total += len(b)
	}

	if l.info.Medium != "" {
		b := bytesFromString(l.info.Medium)
		total += len(b)
	}

	// add 4 to account for the type ID
	return total + 4
}

func (l ListChunk) getADTLChunkSize() int {
	total := 0

	// Get size of all Label chunks
	for _, lb := range l.adtl.Labels {
		total += lb.getChunkSize()
	}
	// Get size of all Note chunks
	for _, n := range l.adtl.Notes {
		total += n.getChunkSize()
	}
	// Get size of all LTXT chunks
	for _, lt := range l.adtl.LabeledTexts {
		total += lt.getChunkSize()
	}

	// Add 4 to account for the type ID
	return total + 4
	// I love you - uwu (c) 2/5/2022
}
