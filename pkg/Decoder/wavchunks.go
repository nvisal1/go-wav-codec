package Decoder

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

type WavChunks struct {
	CuePoints    []*CuePoint
	Fact         *FactChunk
	FMT          *FMTChunk
	Info         *InfoChunk
	ADTL         *ADTLChunk
	Inst         *InstChunk
	PlstSegments []*PlstSegment
	Sample       *SmplChunk
	DataPosition int64
}

func ReadWavChunks(r *bytes.Reader) (*WavChunks, error) {
	wc := &WavChunks{}

	for {
		c, err := NewChunk(r)
		if err == io.EOF {
			return wc, nil
		}

		if err != nil {
			return nil, err
		}

		// FIXME: Move to Chunk method and return error if the provided file is not word-aligned?
		//if size%2 == 1 {
		//	size++
		//}

		switch c.ID {
		case FMT_CHUNK_ID:
			err := handleFMTChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case LIST_CHUNK_ID:
			err := handleLISTChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case FACT_CHUNK_ID:
			err := handleFACTChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case PLST_CHUNK_ID:
			err := handlePLSTChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case SMPL_CHUNK_ID:
			err := handleSMPLChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case INST_CHUNK_ID:
			err := handleINSTChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case CUE_CHUNK_ID:
			err := handleCUEChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
		case DATA_CHUNK_ID:
			err := handleDATAChunk(r, c, wc)
			if err != nil {
				return nil, err
			}
			//return nil
		default:
			_, err = r.Seek(int64(c.Size), 1)
			if err != nil {
				return nil, err
			}
		}
	}
}

func handleFMTChunk(r *bytes.Reader, c *Chunk, wc *WavChunks) error {
	fmtr, err := RecordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}
	fmtc, err := ReadFMTChunk(fmtr)
	if err != nil {
		return err
	}
	wc.FMT = fmtc
	return nil
}

func handleLISTChunk(r *bytes.Reader, c *Chunk, wc *WavChunks) error {
	listr, err := RecordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}

	i := make([]byte, 4)
	err = binary.Read(listr, binary.BigEndian, &i)
	if err != nil {
		return err
	}
	if string(i[:]) == INFO_CHUNK_ID {
		ic, err := ReadINFOChunk(listr)
		if err != nil {
			return err
		}
		wc.Info = ic
	}
	if string(i[:]) == ASSOCIATED_DATA_LIST_CHUNK_ID {
		ac, err := ReadADTLChunk(listr)
		if err != nil {
			return err
		}
		wc.ADTL = ac
	}
	return nil
}

func handleFACTChunk(r *bytes.Reader, c *Chunk, wc *WavChunks) error {
	factr, err := RecordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}

	fc, err := ReadFactChunk(factr)
	if err != nil {
		return err
	}
	wc.Fact = fc
	return nil
}

func handlePLSTChunk(r *bytes.Reader, c *Chunk, wc *WavChunks) error {
	plstr, err := RecordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}

	ps, err := ReadPlstChunk(plstr)
	if err != nil {
		return err
	}
	wc.PlstSegments = ps
	return nil
}

func handleSMPLChunk(r *bytes.Reader, c *Chunk, wc *WavChunks) error {
	smplr, err := RecordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}

	sc, err := ReadSmplChunk(smplr)
	if err != nil {
		return nil
	}
	wc.Sample = sc
	return nil
}

func handleINSTChunk(r *bytes.Reader, c *Chunk, wc *WavChunks) error {
	instr, err := RecordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}

	ic, err := ReadInstChunk(instr)
	if err != nil {
		return err
	}
	wc.Inst = ic
	return nil
}

func handleCUEChunk(r *bytes.Reader, c *Chunk, wc *WavChunks) error {
	cuer, err := RecordAndForward(r, int(c.Size))
	if err != nil {
		return err
	}

	cp, err := ReadCueChunk(cuer)
	if err != nil {
		return err
	}
	wc.CuePoints = cp
	return nil
}

func handleDATAChunk(r *bytes.Reader, c *Chunk, wc *WavChunks) error {
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
	_, err = r.Seek(int64(c.Size), 1)
	if err != nil {
		return err
	}
	return nil
}
