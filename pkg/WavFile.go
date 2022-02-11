package pkg

import (
	"encoding/binary"
	"os"
)

/**
Wav file specification
http://soundfile.sapp.org/doc/WaveFormat/
*/
type WavFile struct {
	ChunkID       string
	ChunkSize     uint32
	Format        string
	Subchunk1ID   string
	Subchunk1Size uint32
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
	Subchunk2ID   string
	Subchunk2Size uint32
	Data          []byte
}

func (w *WavFile) SaveTo(p string) error {
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(w.ChunkID))
	if err != nil {

	}

	chunkSizeBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(chunkSizeBytes, w.ChunkSize)
	_, err = f.Write(chunkSizeBytes)

	_, err = f.Write([]byte(w.Format))

	_, err = f.Write([]byte(w.Subchunk1ID))

	subchunk1SizeBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(subchunk1SizeBytes, w.Subchunk1Size)
	_, err = f.Write(subchunk1SizeBytes)

	audioFormatBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(audioFormatBytes, w.AudioFormat)
	_, err = f.Write(audioFormatBytes)

	numChannelsBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(numChannelsBytes, w.NumChannels)
	_, err = f.Write(numChannelsBytes)

	sampleRateBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(sampleRateBytes, w.SampleRate)
	_, err = f.Write(sampleRateBytes)

	byteRateBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(byteRateBytes, w.ByteRate)
	_, err = f.Write(byteRateBytes)

	blockAlignBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(blockAlignBytes, w.BlockAlign)
	_, err = f.Write(blockAlignBytes)

	bitsPerSampleBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bitsPerSampleBytes, w.BitsPerSample)
	_, err = f.Write(bitsPerSampleBytes)

	_, err = f.Write([]byte(w.Subchunk2ID))

	subchunk2SizeBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(subchunk2SizeBytes, w.Subchunk2Size)
	_, err = f.Write(subchunk2SizeBytes)

	_, err = f.Write(w.Data)

	f.Sync()

	return nil
}
