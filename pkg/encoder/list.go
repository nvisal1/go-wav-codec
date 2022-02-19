package Encoder

import (
	"encoding/binary"
	"errors"
	"io"
)

type ListChunk struct {
	info *InfoChunk
}

func writeLISTChunk(w io.Writer, l *ListChunk) (int, error) {
	bytesWritten := 0

	if l.info == nil {
		return bytesWritten, errors.New("Must include an INFO chunk in the LIST chunk")
	}

	b := bytesFromString(LIST_CHUNK_ID)
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	s := getINFOChunkSize(l)
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

	return bytesWritten, nil
}

func getINFOChunkSize(l *ListChunk) int {
	total := 0

	if l.info.Location != "" {
		b := bytesFromString(l.info.Location)
		total += len(b) + 8
	}

	if l.info.Artist != "" {
		b := bytesFromString(l.info.Artist)
		total += len(b) + 8
	}

	if l.info.Software != "" {
		b := bytesFromString(l.info.Software)
		total += len(b) + 8
	}

	if l.info.CreationDate != "" {
		b := bytesFromString(l.info.CreationDate)
		total += len(b) + 8
	}

	if l.info.Copyright != "" {
		b := bytesFromString(l.info.Copyright)
		total += len(b) + 8
	}

	if l.info.Title != "" {
		b := bytesFromString(l.info.Title)
		total += len(b) + 8
	}

	if l.info.Engineer != "" {
		b := bytesFromString(l.info.Engineer)
		total += len(b) + 8
	}

	if l.info.Genre != "" {
		b := bytesFromString(l.info.Genre)
		total += len(b) + 8
	}

	if l.info.Product != "" {
		b := bytesFromString(l.info.Product)
		total += len(b) + 8
	}

	if l.info.Source != "" {
		b := bytesFromString(l.info.Source)
		total += len(b) + 8
	}

	if l.info.Subject != "" {
		b := bytesFromString(l.info.Subject)
		total += len(b) + 8
	}

	if l.info.Comments != "" {
		b := bytesFromString(l.info.Comments)
		total += len(b) + 8
	}

	if l.info.Technician != "" {
		b := bytesFromString(l.info.Technician)
		total += len(b) + 8
	}

	if l.info.Keywords != "" {
		b := bytesFromString(l.info.Keywords)
		total += len(b) + 8
	}

	if l.info.Medium != "" {
		b := bytesFromString(l.info.Medium)
		total += len(b) + 8
	}

	// add 4 to account for the type ID
	return total + 4
}
