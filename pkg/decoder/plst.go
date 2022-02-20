package decoder

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type plstSegment struct {
	CuePointID      string
	Length          uint32
	NumberOfRepeats uint32
}

func readPlstChunk(r *bytes.Reader) ([]*plstSegment, error) {
	var numSegments uint32

	if err := binary.Read(r, binary.LittleEndian, &numSegments); err != nil {
		return nil, errors.New("An error occurred when reading the number of PLST segments")
	}

	if numSegments > 0 {
		str := make([]byte, 4)
		s := make([]*plstSegment, 0, numSegments)

		for i := uint32(0); i < numSegments; i++ {
			p := &plstSegment{}

			if err := binary.Read(r, binary.BigEndian, &str); err != nil {
				return nil, errors.New("An error occurred when reading the PLST segment cue point ID")
			}

			p.CuePointID = string(str[:])

			if err := binary.Read(r, binary.LittleEndian, &p.Length); err != nil {
				return nil, errors.New("An error occurred when reading the PLST segment length")
			}

			if err := binary.Read(r, binary.LittleEndian, &p.NumberOfRepeats); err != nil {
				return nil, errors.New("An error occurred when reading the PLST segment position")
			}

			s = append(s, p)
		}
		return s, nil
	}
	return nil, nil
}
