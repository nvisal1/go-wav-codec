package pkg

//
//import "testing"
//
//func TestConcat(t *testing.T) {
//	path1 := "../recording-1.wav"
//	path2 := "../recording-2.wav"
//
//	w1, w2, w3, err := Concat(path1, path2)
//	if err != nil {
//		panic(err)
//	}
//
//	expectedSubchunk2Size := w1.Subchunk2Size + w2.Subchunk2Size
//	if w3.Subchunk2Size != expectedSubchunk2Size {
//		t.Errorf("Subchunk 2 Size is not expected value [%d], Actual Subchunk 2 Size is [%d]", expectedSubchunk2Size, w3.Subchunk2Size)
//	}
//
//	expectedChunkSize := expectedSubchunk2Size + 32
//	if w3.ChunkSize != expectedChunkSize {
//		t.Errorf("Chunk Size is not expected value [%d], Actual Chunk Size is [%d]", expectedChunkSize, w3.ChunkSize)
//	}
//
//	if len(w3.Data) != int(expectedSubchunk2Size) {
//		t.Errorf("Length of Data is not expected value [%d], Actual length of Data is [%d]", expectedSubchunk2Size, len(w3.Data))
//	}
//}
