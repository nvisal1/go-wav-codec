package Decoder

import (
	"bytes"
	"encoding/binary"
	"io"
	"strings"
)

const (
	IARL = "IARL"
	IART = "IART"
	ICMS = "ICMS"
	ICMT = "ICMT"
	ICOP = "ICOP"
	ICRD = "ICRD"
	ICRP = "ICRP"
	IDIM = "IDIM"
	IDPI = "IDPI"
	IENG = "IENG"
	IGNR = "IGNR"
	IKEY = "IKEY"
	ILGT = "ILGT"
	IMED = "IMED"
	INAM = "INAM"
	IPLT = "INAM"
	IPRD = "IPRD"
	ISBJ = "ISBJ"
	ISFT = "ISFT"
	ISRC = "ISRC"
	ISRF = "ISRF"
	ITCH = "ITCH"
	ITRK = "ITRK"
)

type InfoChunk struct {
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

func ReadINFOChunk(r *bytes.Reader) (*InfoChunk, error) {
	i := &InfoChunk{}

	for {
		c, err := NewChunk(r)
		if err != nil {
			if err == io.EOF {
				return i, nil
			}
			return nil, err
		}

		scratch := make([]byte, c.Size)

		if err = binary.Read(r, binary.BigEndian, &scratch); err != nil {
			return nil, err
		}

		switch c.ID {
		case IARL:
			i.Location = removeNullCharacters(string(scratch[:]))
		case IART:
			i.Artist = removeNullCharacters(string(scratch[:]))
		case ISFT:
			i.Software = removeNullCharacters(string(scratch[:]))
		case ICRD:
			i.CreationDate = removeNullCharacters(string(scratch[:]))
		case ICOP:
			i.Copyright = removeNullCharacters(string(scratch[:]))
		case INAM:
			i.Title = removeNullCharacters(string(scratch[:]))
		case IENG:
			i.Engineer = removeNullCharacters(string(scratch[:]))
		case IGNR:
			i.Genre = removeNullCharacters(string(scratch[:]))
		case IPRD:
			i.Product = removeNullCharacters(string(scratch[:]))
		case ISRC:
			i.Source = removeNullCharacters(string(scratch[:]))
		case ISBJ:
			i.Subject = removeNullCharacters(string(scratch[:]))
		case ICMT:
			i.Comments = removeNullCharacters(string(scratch[:]))
		case ITCH:
			i.Technician = removeNullCharacters(string(scratch[:]))
		case IKEY:
			i.Keywords = removeNullCharacters(string(scratch[:]))
		case IMED:
			i.Medium = removeNullCharacters(string(scratch[:]))
		}
	}
	return i, nil
}
