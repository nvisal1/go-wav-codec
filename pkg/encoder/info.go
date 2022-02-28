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

func writeInfoPart(w io.Writer, t string, v string) (int, error) {
	bytesWritten := 0
	l := len(bytesFromString(v))
	ba := align(bytesFromString(v))

	// Write the chunk ID
	if err := binary.Write(w, binary.BigEndian, align(bytesFromString(t))); err != nil {
		return bytesWritten, err
	}

	bytesWritten += len(bytesFromString(t))

	// Write the chunk size
	if err := binary.Write(w, binary.LittleEndian, uint32(len(ba))); err != nil {
		return bytesWritten, err
	}
	bytesWritten += len(ba)

	// Write the chunk data
	if err := binary.Write(w, binary.BigEndian, &ba); err != nil {
		return bytesWritten, err
	}
	bytesWritten += l
	return bytesWritten, nil
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
		n, err := writeInfoPart(w, iARL, i.Location)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	if i.Artist != "" {
		n, err := writeInfoPart(w, iART, i.Artist)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	if i.Software != "" {
		n, err := writeInfoPart(w, iSFT, i.Software)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	if i.CreationDate != "" {
		n, err := writeInfoPart(w, iCRD, i.CreationDate)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	if i.Copyright != "" {
		n, err := writeInfoPart(w, iCOP, i.Copyright)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	if i.Title != "" {
		n, err := writeInfoPart(w, iNAM, i.Title)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	if i.Engineer != "" {
		n, err := writeInfoPart(w, iENG, i.Engineer)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	if i.Genre != "" {
		n, err := writeInfoPart(w, iGNR, i.Genre)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	if i.Product != "" {
		n, err := writeInfoPart(w, iPRD, i.Product)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	if i.Source != "" {
		n, err := writeInfoPart(w, iSRC, i.Source)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	if i.Subject != "" {
		n, err := writeInfoPart(w, iSBJ, i.Subject)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
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

		n, err := writeInfoPart(w, iCMT, i.Comments)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	if i.Technician != "" {
		n, err := writeInfoPart(w, iTCH, i.Technician)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	if i.Keywords != "" {
		n, err := writeInfoPart(w, iKEY, i.Keywords)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	if i.Medium != "" {
		n, err := writeInfoPart(w, iMED, i.Medium)
		if err != nil {
			return bytesWritten, err
		}
		bytesWritten += n
	}

	return bytesWritten, nil
}
