package decoder

import (
	"bytes"
	"encoding/binary"
)

/**
0x08	1	Unshifted note	0 - 127
0x09	1	Fine Tune (dB)	-50 - +50
0x0A	1	Gain	-64 - +64
0x0B	1	Low note	0 - 127
0x0C	1	High note	0 - 127
0x0D	1	Low Velocity	1 - 127
0x0E	1	High Velocity	1 - 127
*/
type instChunk struct {
	UnshiftedNote uint8
	FineTuneDB    uint8
	Gain          uint8
	LowNote       uint8
	HighNote      uint8
	LowVelocity   uint8
	HighVelocity  uint8
}

func readInstChunk(r *bytes.Reader) (*instChunk, error) {

	i := &instChunk{}

	if err := binary.Read(r, binary.LittleEndian, &i.UnshiftedNote); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &i.FineTuneDB); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &i.Gain); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &i.LowNote); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &i.HighNote); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &i.HighVelocity); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &i.LowVelocity); err != nil {
		return nil, err
	}

	return i, nil
}
