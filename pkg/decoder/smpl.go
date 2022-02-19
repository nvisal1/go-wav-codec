package Decoder

import (
	"bytes"
	"encoding/binary"
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

func ReadSmplChunk(r *bytes.Reader) (*SmplChunk, error) {
	s := &SmplChunk{}

	B32 := make([]byte, 4)
	if err := binary.Read(r, binary.BigEndian, &B32); err != nil {
		return nil, err
	}
	s.Manufacturer = string(B32[:])

	if err := binary.Read(r, binary.BigEndian, &B32); err != nil {
		return nil, err
	}
	s.Product = string(B32[:])

	if err := binary.Read(r, binary.LittleEndian, &s.SamplePeriod); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &s.MIDIUnityNote); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &s.MIDIPitchFraction); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &s.SMPTEFormat); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &s.SMPTEOffset); err != nil {
		return nil, err
	}

	var numSmplLoops uint32
	if err := binary.Read(r, binary.LittleEndian, &numSmplLoops); err != nil {
		return nil, err
	}

	if numSmplLoops > 0 {
		for i := uint32(0); i < numSmplLoops; i++ {
			l := &SampleLoop{}

			if err := binary.Read(r, binary.BigEndian, &B32); err != nil {
				return nil, err
			}
			l.CuePointID = string(B32[:])

			if err := binary.Read(r, binary.LittleEndian, &l.Type); err != nil {
				return nil, err
			}
			if err := binary.Read(r, binary.LittleEndian, &l.Start); err != nil {
				return nil, err
			}
			if err := binary.Read(r, binary.LittleEndian, &l.End); err != nil {
				return nil, err
			}
			if err := binary.Read(r, binary.LittleEndian, &l.Fraction); err != nil {
				return nil, err
			}
			if err := binary.Read(r, binary.LittleEndian, &l.PlayCount); err != nil {
				return nil, err
			}

			s.Loops = append(s.Loops, l)
		}
	}
	return s, nil
}
