package Encoder

import (
	"encoding/binary"
	"io"
)

type SampleLoop struct {
	CuePointID string
	Type       uint32
	Start      uint32
	End        uint32
	Fraction   uint32
	PlayCount  uint32
}

type SmplChunk struct {
	Manufacturer      string
	Product           string
	SamplePeriod      uint32
	MIDIUnityNote     uint32
	MIDIPitchFraction uint32
	SMPTEFormat       uint32
	SMPTEOffset       uint32
	Loops             []*SampleLoop
}

func (s SmplChunk) WriteTo(w io.Writer) (int, error) {
	bytesWritten := 0

	b := bytesFromString(s.Manufacturer)
	if err := binary.Write(w, binary.BigEndian, b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromString(s.Product)
	if err := binary.Write(w, binary.BigEndian, b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromUINT32(s.SamplePeriod)
	if err := binary.Write(w, binary.LittleEndian, b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromUINT32(s.MIDIUnityNote)
	if err := binary.Write(w, binary.LittleEndian, b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromUINT32(s.MIDIPitchFraction)
	if err := binary.Write(w, binary.LittleEndian, b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromUINT32(s.SMPTEFormat)
	if err := binary.Write(w, binary.LittleEndian, b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	b = bytesFromUINT32(s.SMPTEOffset)
	if err := binary.Write(w, binary.LittleEndian, b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	numSmplLoops := uint32(len(s.Loops))
	b = bytesFromUINT32(numSmplLoops)
	if err := binary.Write(w, binary.LittleEndian, b); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	if numSmplLoops > 0 {
		for i := uint32(0); i < numSmplLoops; i++ {
			b := bytesFromString(s.Loops[i].CuePointID)
			if err := binary.Write(w, binary.BigEndian, b); err != nil {
				return bytesWritten, err
			}
			bytesWritten += len(b)

			b = bytesFromUINT32(s.Loops[i].Type)
			if err := binary.Write(w, binary.LittleEndian, b); err != nil {
				return bytesWritten, err
			}
			bytesWritten += len(b)

			b = bytesFromUINT32(s.Loops[i].Start)
			if err := binary.Write(w, binary.LittleEndian, b); err != nil {
				return bytesWritten, err
			}
			bytesWritten += len(b)

			b = bytesFromUINT32(s.Loops[i].End)
			if err := binary.Write(w, binary.LittleEndian, b); err != nil {
				return bytesWritten, err
			}
			bytesWritten += len(b)

			b = bytesFromUINT32(s.Loops[i].Fraction)
			if err := binary.Write(w, binary.LittleEndian, b); err != nil {
				return bytesWritten, err
			}
			bytesWritten += len(b)

			b = bytesFromUINT32(s.Loops[i].PlayCount)
			if err := binary.Write(w, binary.LittleEndian, b); err != nil {
				return bytesWritten, err
			}
			bytesWritten += len(b)
		}
	}
	return bytesWritten, nil
}
