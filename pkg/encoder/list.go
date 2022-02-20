package encoder

import (
	"encoding/binary"
	"io"
)

func writeLISTChunk(w io.Writer, i *InfoChunk) (int, error) {
	bytesWritten := 0

	b := bytesFromString(listChunkID)
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	s := getINFOChunkSize(i)
	b = bytesFromUINT32(uint32(s))
	err = binary.Write(w, binary.LittleEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	n, err := writeInfoChunk(w, i)
	bytesWritten += n
	if err != nil {
		return bytesWritten, err
	}

	return bytesWritten, nil
}

func getINFOChunkSize(i *InfoChunk) int {
	total := 0

	if i.Location != "" {
		b := bytesFromString(i.Location)
		total += len(b) + 8
	}

	if i.Artist != "" {
		b := bytesFromString(i.Artist)
		total += len(b) + 8
	}

	if i.Software != "" {
		b := bytesFromString(i.Software)
		total += len(b) + 8
	}

	if i.CreationDate != "" {
		b := bytesFromString(i.CreationDate)
		total += len(b) + 8
	}

	if i.Copyright != "" {
		b := bytesFromString(i.Copyright)
		total += len(b) + 8
	}

	if i.Title != "" {
		b := bytesFromString(i.Title)
		total += len(b) + 8
	}

	if i.Engineer != "" {
		b := bytesFromString(i.Engineer)
		total += len(b) + 8
	}

	if i.Genre != "" {
		b := bytesFromString(i.Genre)
		total += len(b) + 8
	}

	if i.Product != "" {
		b := bytesFromString(i.Product)
		total += len(b) + 8
	}

	if i.Source != "" {
		b := bytesFromString(i.Source)
		total += len(b) + 8
	}

	if i.Subject != "" {
		b := bytesFromString(i.Subject)
		total += len(b) + 8
	}

	if i.Comments != "" {
		b := bytesFromString(i.Comments)
		total += len(b) + 8
	}

	if i.Technician != "" {
		b := bytesFromString(i.Technician)
		total += len(b) + 8
	}

	if i.Keywords != "" {
		b := bytesFromString(i.Keywords)
		total += len(b) + 8
	}

	if i.Medium != "" {
		b := bytesFromString(i.Medium)
		total += len(b) + 8
	}

	// add 4 to account for the type ID
	return total + 4
}
