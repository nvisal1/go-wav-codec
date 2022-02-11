package pkg

//
//import (
//	"os"
//)
//
//func NewWavFile() *WavFile {
//	return &WavFile{}
//}
//
//func NewWavFileFromHeaders(h *WavHeaders) (*WavFile, error) {
//	w := &WavFile{}
//	mapHeadersToWavFile(w, h)
//	return w, nil
//}
//
//func NewWavFileFromPath(path string) (*WavFile, error) {
//	F, err := os.Open(path)
//	if err != nil {
//		panic(err)
//	}
//	defer F.Close()
//	w := &WavFile{}
//
//	h, err := NewWavReader(F).ReadHeaders()
//	mapHeadersToWavFile(w, h)
//	return w, nil
//}
//
//func mapHeadersToWavFile(w *WavFile, h *WavHeaders) {
//	w.ChunkID = h.FileDescriptor
//	w.ChunkSize = h.FileSize
//	w.Format = h.FileFormat
//	w.Subchunk1ID = "fmt "
//	w.Subchunk1Size = 0
//	w.AudioFormat = h.AudioFormat
//	w.NumChannels = h.NumChannels
//	w.SampleRate = h.SampleRate
//	w.ByteRate = h.ByteRate
//	w.BlockAlign = h.BlockAlign
//	w.BitsPerSample = h.BitsPerSample
//	w.Subchunk2ID = "data"
//	w.Subchunk2Size = h.DataLen
//}
