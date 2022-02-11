package pkg

//
//func Concat(path1 string, path2 string) (*WavFile, *WavFile, *WavFile, error) {
//	w1, _ := NewWavFileFromPath(path1)
//	w2, _ := NewWavFileFromPath(path2)
//	w3 := NewWavFile()
//
//	// Add data from the second file to the end of the first file
//	w3.Data = append(w1.Data, w2.Data...)
//
//	// Modify Subchunk2Size
//	w3.Subchunk2Size = uint32(len(w3.Data))
//
//	// Modify ChunkSize (36 + SubChunk2Size)
//	w3.ChunkSize = w3.Subchunk2Size + 32
//
//	// Copy the rest of the fields from file1
//	w3.ChunkID = w1.ChunkID
//	w3.Format = w1.Format
//	w3.Subchunk1ID = w1.Subchunk1ID
//	w3.Subchunk1Size = w1.Subchunk1Size
//	w3.AudioFormat = w1.AudioFormat
//	w3.NumChannels = w1.NumChannels
//	w3.SampleRate = w1.SampleRate
//	w3.ByteRate = w1.ByteRate
//	w3.BlockAlign = w1.BlockAlign
//	w3.BitsPerSample = w1.BitsPerSample
//	w3.Subchunk2ID = w1.Subchunk2ID
//
//	return w1, w2, w3, nil
//}
