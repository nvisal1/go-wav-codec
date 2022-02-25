package decoder

import (
	"bytes"
	"encoding/binary"
	"io"
	"strings"
)

const (
	iARL = "IARL"
	iART = "IART"
	iCMS = "ICMS"
	iCMT = "ICMT"
	iCOP = "ICOP"
	iCRD = "ICRD"
	iCRP = "ICRP"
	iDIM = "IDIM"
	iDPI = "IDPI"
	iENG = "IENG"
	iGNR = "IGNR"
	iKEY = "IKEY"
	iLGT = "ILGT"
	iMED = "IMED"
	iNAM = "INAM"
	iPLT = "INAM"
	iPRD = "IPRD"
	iSBJ = "ISBJ"
	iSFT = "ISFT"
	iSRC = "ISRC"
	iSRF = "ISRF"
	iTCH = "ITCH"
	iTRK = "ITRK"
)

type infoChunk struct {
	Location     string
	Artist       string
	Software     string
	CreationDate string
	Copyright    string
	Title        string
	Engineer     string
	Genre        string
	Product      string
	Source       string
	Subject      string
	Comments     string
	Technician   string
	Keywords     string
	Medium       string
}

func removeNullCharacters(s string) string {
	return strings.ReplaceAll(s, "\u0000", "")
}

func readINFOChunk(r *bytes.Reader) (*infoChunk, error) {
	i := &infoChunk{}

	for {
		c, err := newChunk(r)
		if err != nil {
			if err == io.EOF {
				return i, nil
			}
			return nil, err
		}

		if c.Size%2 != 0 {
			c.Size++
		}

		c.ID = strings.ToUpper(c.ID)

		scratch := make([]byte, c.Size)

		if err = binary.Read(r, binary.BigEndian, &scratch); err != nil {
			return nil, err
		}

		switch c.ID {
		case iARL:
			i.Location = removeNullCharacters(string(scratch[:]))
		case iART:
			i.Artist = removeNullCharacters(string(scratch[:]))
		case iSFT:
			i.Software = removeNullCharacters(string(scratch[:]))
		case iCRD:
			i.CreationDate = removeNullCharacters(string(scratch[:]))
		case iCOP:
			i.Copyright = removeNullCharacters(string(scratch[:]))
		case iNAM:
			i.Title = removeNullCharacters(string(scratch[:]))
		case iENG:
			i.Engineer = removeNullCharacters(string(scratch[:]))
		case iGNR:
			i.Genre = removeNullCharacters(string(scratch[:]))
		case iPRD:
			i.Product = removeNullCharacters(string(scratch[:]))
		case iSRC:
			i.Source = removeNullCharacters(string(scratch[:]))
		case iSBJ:
			i.Subject = removeNullCharacters(string(scratch[:]))
		case iCMT:
			i.Comments = removeNullCharacters(string(scratch[:]))
		case iTCH:
			i.Technician = removeNullCharacters(string(scratch[:]))
		case iKEY:
			i.Keywords = removeNullCharacters(string(scratch[:]))
		case iMED:
			i.Medium = removeNullCharacters(string(scratch[:]))
		}
	}
}
