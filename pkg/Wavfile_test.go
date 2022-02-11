package pkg

//func TestWavFile_WriteTo(t *testing.T) {
//	path := "../recording-3.wav"
//	os.Remove(path)
//
//	w := NewWavFileFromPath("../recording-1.wav")
//
//	err := w.WriteTo(path)
//
//	if err != nil {
//		panic(err)
//	}
//
//	w2 := NewWavFileFromPath(path)
//
//	if w2.Subchunk2Size != w.Subchunk2Size {
//		t.Errorf("Subchunk 2 Size is not expected value [%d], Actual Subchunk 2 Size is [%d]", w.Subchunk2Size, w2.Subchunk2Size)
//	}
//}
