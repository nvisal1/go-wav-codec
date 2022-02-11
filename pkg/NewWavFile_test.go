package pkg

//
//import "testing"
//
//func TestNewWavFile(t *testing.T) {
//	w := NewWavFile()
//
//	if w.ChunkID != "" {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//
//	if w.ChunkSize != 0 {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//
//	if w.Format != "" {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//
//	if w.Subchunk1ID != "" {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//
//	if w.Subchunk1Size != 0 {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//
//	if w.AudioFormat != 0 {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//
//	if w.NumChannels != 0 {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//
//	if w.SampleRate != 0 {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//
//	if w.ByteRate != 0 {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//
//	if w.BlockAlign != 0 {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//
//	if w.BitsPerSample != 0 {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//
//	if w.Subchunk2ID != "" {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//
//	if w.Subchunk2Size != 0 {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//
//	if len(w.Data) != 0 {
//		t.Error("NewWavFile did not return an empty WavFile")
//	}
//}
//
//func TestNewWavFileFromPath(t *testing.T) {
//
//	w := NewWavFileFromPath("../recording-1.wav")
//
//	expectedChunkID := "RIFF"
//	if w.ChunkID != expectedChunkID {
//		t.Errorf("Chunk ID is not expected value [%s], Actual Chunk ID is [%s]", expectedChunkID, w.ChunkID)
//	}
//
//	expectedChunkSize := uint32(3840036)
//	if w.ChunkSize != expectedChunkSize {
//		t.Errorf("Chunk Size is not expected value [%d], Actual Chunk Size is [%d]", expectedChunkSize, w.ChunkSize)
//	}
//
//	expectedFormat := "WAVE"
//	if w.Format != expectedFormat {
//		t.Errorf("Format is not expected value [%s], Actual Format is [%s]", expectedFormat, w.Format)
//	}
//
//	expectedSubchunk1ID := "fmt "
//	if w.Subchunk1ID != expectedSubchunk1ID {
//		t.Errorf("Subchunk 1 ID is not expected value [%s], Actual Subchunk 1 ID is [%s]", expectedSubchunk1ID, w.Subchunk1ID)
//	}
//
//	expectedSubchunk1Size := uint32(16)
//	if w.Subchunk1Size != expectedSubchunk1Size {
//		t.Errorf("Subchunk 1 Size is not expected value [%d], Actual Subchunk 1 Size is [%d]", expectedSubchunk1Size, w.Subchunk1Size)
//	}
//
//	expectedAudioFormat := uint16(1)
//	if w.AudioFormat != expectedAudioFormat {
//		t.Errorf("Audio Format is not expected value [%d], Actual Audio Format is [%d]", expectedAudioFormat, w.AudioFormat)
//	}
//
//	expectedNumChannels := uint16(2)
//	if w.NumChannels != expectedNumChannels {
//		t.Errorf("Num Channels is not expected value [%d], Actual Num Channels is [%d]", expectedNumChannels, w.NumChannels)
//	}
//
//	expectedSampleRate := uint32(48000)
//	if w.SampleRate != expectedSampleRate {
//		t.Errorf("Sample Rate is not expected value [%d], Actual Sample Rate is [%d]", expectedSampleRate, w.SampleRate)
//	}
//
//	expectedByteRate := uint32(192000)
//	if w.ByteRate != expectedByteRate {
//		t.Errorf("Byte Rate is not expected value [%d], Actual Byte Rate is [%d]", expectedByteRate, w.ByteRate)
//	}
//
//	expectedBlockAlign := uint16(4)
//	if w.BlockAlign != expectedBlockAlign {
//		t.Errorf("Block Align is not expected value [%d], Actual Block Align is [%d]", expectedBlockAlign, w.BlockAlign)
//	}
//
//	expectedBitsPerSample := uint16(16)
//	if w.BitsPerSample != expectedBitsPerSample {
//		t.Errorf("Bits Per Sample is not expected value [%d], Actual Bits Per Sample is [%d]", expectedBitsPerSample, w.BitsPerSample)
//	}
//
//	expectedSubchunk2ID := "data"
//	if w.Subchunk2ID != expectedSubchunk2ID {
//		t.Errorf("Subchunk 2 ID is not expected value [%s], Actual Subchunk 2 ID is [%s]", expectedSubchunk2ID, w.Subchunk2ID)
//	}
//
//	expectedSubchunk2Size := uint32(3840000)
//	if w.Subchunk2Size != expectedSubchunk2Size {
//		t.Errorf("Subchunk 2 Size is not expected value [%d], Actual Subchunk 2 Size is [%d]", expectedSubchunk2Size, w.Subchunk2Size)
//	}
//
//	if len(w.Data) != int(expectedSubchunk2Size) {
//		t.Errorf("Length of Data is not expected value [%d], Actual length of Data is [%d]", expectedSubchunk2Size, len(w.Data))
//	}
//}
