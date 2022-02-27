package decoder

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strings"
)

type wavChunks struct {
	CuePoints    []*cuePoint
	Fact         *factChunk
	FMT          *fmtChunk
	Info         *infoChunk
	ADTL         *adtlChunk
	Inst         *instChunk
	PlstSegments []*plstSegment
	Sample       *smplChunk
	DataPosition int64
	DataLength   uint32
	NumSamples   uint32
}

func readWavChunks(r *bytes.Reader) (*wavChunks, error) {
	wc := &wavChunks{}

	for {
		c, err := newChunk(r)
		if err == io.EOF {
			return wc, nil
		}

		if err != nil {
			return nil, err
		}

		// If the size is an odd number, read an extra byte
		// in order to stay in a word-aligned position
		if c.Size%2 == 1 {
			c.Size++
		}

		c.ID = strings.ToUpper(c.ID)

		switch c.ID {
		case fmtChunkID:
			err := handleFMTChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case listChunkID:
			err := handleLISTChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case factChunkID:
			err := handleFACTChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case plstChunkID:
			err := handlePLSTChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case smplChunkID:
			err := handleSMPLChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case instChunkID:
			err := handleINSTChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case cueChunkID:
			err := handleCUEChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case dataChunkID:
			err := handleDATAChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		default:
			_, err = r.Seek(int64(c.Size), 1)
			if err != nil {
				return nil, err
			}
		}
	}
}

func handleFMTChunk(r *bytes.Reader, c *chunk, wc *wavChunks) error {
	fmtr, err := recordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}
	fmtc, err := readFMTChunk(fmtr)
	if err != nil {
		return err
	}
	wc.FMT = fmtc
	return nil
}

func handleLISTChunk(r *bytes.Reader, c *chunk, wc *wavChunks) error {
	listr, err := recordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}

	i := make([]byte, 4)
	err = binary.Read(listr, binary.BigEndian, &i)
	if err != nil {
		return err
	}
	if strings.ToUpper(string(i[:])) == infoChunkID {
		ic, err := readINFOChunk(listr)
		if err != nil {
			return err
		}
		wc.Info = ic
	}
	if strings.ToUpper(string(i[:])) == associatedDataListChunkID {
		ac, err := readADTLChunk(listr)
		if err != nil {
			return err
		}
		wc.ADTL = ac
	}
	return nil
}

func handleFACTChunk(r *bytes.Reader, c *chunk, wc *wavChunks) error {
	factr, err := recordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}

	fc, err := readFactChunk(factr)
	if err != nil {
		return err
	}
	wc.Fact = fc
	return nil
}

func handlePLSTChunk(r *bytes.Reader, c *chunk, wc *wavChunks) error {
	plstr, err := recordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}

	ps, err := readPlstChunk(plstr)
	if err != nil {
		return err
	}
	wc.PlstSegments = ps
	return nil
}

func handleSMPLChunk(r *bytes.Reader, c *chunk, wc *wavChunks) error {
	smplr, err := recordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}

	sc, err := readSmplChunk(smplr)
	if err != nil {
		return err
	}
	wc.Sample = sc
	return nil
}

func handleINSTChunk(r *bytes.Reader, c *chunk, wc *wavChunks) error {
	instr, err := recordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}

	ic, err := readInstChunk(instr)
	if err != nil {
		return err
	}
	wc.Inst = ic
	return nil
}

func handleCUEChunk(r *bytes.Reader, c *chunk, wc *wavChunks) error {
	cuer, err := recordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}

	cp, err := readCueChunk(cuer)
	if err != nil {
		return err
	}
	wc.CuePoints = cp
	return nil
}

func handleDATAChunk(r *bytes.Reader, c *chunk, wc *wavChunks) error {
	if wc.FMT == nil {
		return errors.New("Data chunk was found before fmt chunk")
	}

	// Record the offset of the PCM data and skip to the end
	// in order to check for metadata chunks
	p, err := r.Seek(0, 1)
	if err != nil {
		return err
	}

	wc.DataPosition = p
	wc.DataLength = c.Size
	wc.NumSamples = uint32(int(wc.DataLength) / int(calculateBytesPerSample(wc.FMT.BitsPerSample)))
	_, err = r.Seek(int64(c.Size), 1)
	if err != nil {
		return err
	}
	return nil
}
