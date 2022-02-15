package Encoder

import (
	"encoding/binary"
	"io"
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

func (i InfoChunk) WriteTo(w io.Writer) (int, error) {
	bytesWritten := 0

	b := bytesFromString(INFO_CHUNK_ID)
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	if i.Location != "" {
		b := Align(bytesFromString(i.Location))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(IARL))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	if i.Artist != "" {
		l := len(bytesFromString(i.Artist))
		ba := Align(bytesFromString(i.Artist))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(IART))); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(bytesFromString(IART))
		if err := binary.Write(w, binary.LittleEndian, uint32(len(ba))); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(ba)
		if err := binary.Write(w, binary.BigEndian, &ba); err != nil {
			return bytesWritten, err
		}
		bytesWritten += l
	}

	if i.Software != "" {
		b := Align(bytesFromString(i.Software))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(ISFT))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	if i.CreationDate != "" {
		b := Align(bytesFromString(i.CreationDate))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(ICRD))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	if i.Copyright != "" {
		b := Align(bytesFromString(i.Copyright))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(ICOP))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	if i.Title != "" {
		b := Align(bytesFromString(i.Title))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(INAM))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	if i.Engineer != "" {
		b := Align(bytesFromString(i.Engineer))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(IENG))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	if i.Genre != "" {
		b := Align(bytesFromString(i.Genre))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(IGNR))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	if i.Product != "" {
		b := Align(bytesFromString(i.Product))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(IPRD))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	if i.Source != "" {
		b := Align(bytesFromString(i.Source))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(ISRC))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	if i.Subject != "" {
		b := Align(bytesFromString(i.Subject))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(ISBJ))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	if i.Comments != "" {
		b := Align(bytesFromString(i.Comments))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(ICMT))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	if i.Technician != "" {
		b := Align(bytesFromString(i.Technician))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(ITCH))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	if i.Keywords != "" {
		b := Align(bytesFromString(i.Keywords))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(IKEY))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	if i.Medium != "" {
		b := Align(bytesFromString(i.Medium))
		if err := binary.Write(w, binary.BigEndian, Align(bytesFromString(IMED))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.LittleEndian, bytesFromUINT32(uint32(len(b)))); err != nil {
			return bytesWritten, err
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(b)
	}

	return bytesWritten, nil
}
