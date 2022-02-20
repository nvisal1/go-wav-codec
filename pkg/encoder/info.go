package encoder

import (
	"encoding/binary"
	"io"
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

// InfoChunk should be used for adding metadata to a wav file
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

func writeInfoChunk(w io.Writer, i *InfoChunk) (int, error) {
	bytesWritten := 0

	b := bytesFromString(infoChunkID)
	err := binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(b)

	if i.Location != "" {
		b := align(bytesFromString(i.Location))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iARL))); err != nil {
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
		ba := align(bytesFromString(i.Artist))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iART))); err != nil {
			return bytesWritten, err
		}
		bytesWritten += len(bytesFromString(iART))
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
		b := align(bytesFromString(i.Software))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iSFT))); err != nil {
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
		b := align(bytesFromString(i.CreationDate))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iCRD))); err != nil {
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
		b := align(bytesFromString(i.Copyright))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iCOP))); err != nil {
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
		b := align(bytesFromString(i.Title))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iNAM))); err != nil {
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
		b := align(bytesFromString(i.Engineer))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iENG))); err != nil {
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
		b := align(bytesFromString(i.Genre))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iGNR))); err != nil {
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
		b := align(bytesFromString(i.Product))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iPRD))); err != nil {
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
		b := align(bytesFromString(i.Source))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iSRC))); err != nil {
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
		b := align(bytesFromString(i.Subject))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iSBJ))); err != nil {
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
		b := align(bytesFromString(i.Comments))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iCMT))); err != nil {
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
		b := align(bytesFromString(i.Technician))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iTCH))); err != nil {
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
		b := align(bytesFromString(i.Keywords))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iKEY))); err != nil {
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
		b := align(bytesFromString(i.Medium))
		if err := binary.Write(w, binary.BigEndian, align(bytesFromString(iMED))); err != nil {
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
