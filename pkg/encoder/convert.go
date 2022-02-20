package encoder

import "encoding/binary"

func bytesFromString(d string) []byte {
	return []byte(d)
}

func bytesFromUINT16(d uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, d)
	return b
}

func bytesFromUINT32(d uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, d)
	return b
}
