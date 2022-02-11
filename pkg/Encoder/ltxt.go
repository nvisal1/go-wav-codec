package Encoder

import (
	"encoding/binary"
	"io"
)

type LabeledText struct {
	CuePointID   string
	SampleLength uint32
	PurposeID    string
	Country      string
	Language     string
	Dialect      string
	CodePage     string
	Data         string
}

func (lt LabeledText) WriteTo(w io.Writer) (int, error) {
	bytesWritten := 0

	b := bytesFromString(TEXT_CHUNK_ID)
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = Align(bytesFromUINT32(uint32(lt.getChunkSize())))
	err = binary.Write(w, binary.LittleEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromString(lt.CuePointID)
	if err := binary.Write(w, binary.BigEndian, &b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)
	b = bytesFromUINT32(lt.SampleLength)
	if err := binary.Write(w, binary.BigEndian, &b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)
	b = bytesFromString(lt.PurposeID)
	if err := binary.Write(w, binary.BigEndian, &b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)
	b = bytesFromString(lt.Country)
	if err := binary.Write(w, binary.BigEndian, &b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)
	b = bytesFromString(lt.Language)
	if err := binary.Write(w, binary.BigEndian, &b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)
	b = bytesFromString(lt.Dialect)
	if err := binary.Write(w, binary.BigEndian, &b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)
	b = bytesFromString(lt.CodePage)
	if err := binary.Write(w, binary.BigEndian, &b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)
	b = bytesFromString(lt.Data)
	if err := binary.Write(w, binary.BigEndian, &b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	return bytesWritten, nil
}

func (lt LabeledText) getChunkSize() int {
	b := bytesFromString(lt.Data)
	return len(b) + 28
}
