package decoder

import (
	"bytes"
	"encoding/binary"
)

type sampleLoop struct {
	CuePointID string
	Type       uint32
	Start      uint32
	End        uint32
	Fraction   uint32
	PlayCount  uint32
}

type smplChunk struct {
	Manufacturer      string
	Product           string
	SamplePeriod      uint32
	MIDIUnityNote     uint32
	MIDIPitchFraction uint32
	SMPTEFormat       uint32
	SMPTEOffset       uint32
	Loops             []*sampleLoop
}

func readSmplChunk(r *bytes.Reader) (*smplChunk, error) {
	s := &smplChunk{}

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
			l := &sampleLoop{}

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
