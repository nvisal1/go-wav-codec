package pkg

import (
	"fmt"
	"os"
	"testing"
)

// ReadHeaders
// -- 8 bit (done)
// -- 16 bit (done)
// -- 24 bit
// -- 32 bit (done)
// -- misaligned inst chunk
// -- with list
// ---- with info (done)
// ---- with associated data list
// -------- with labl
// -------- with note
// -------- with ltxt
// -- with fact
// -- with plst
// -- with smpl
// -- with inst
// -- with cue (done)

func TestDecoder_ReadHeaders_8Bit(t *testing.T) {
	f, err := os.Open("../8bit.wav")
	if err != nil {
		t.Error("Test failed when reading ../8bit.wav")
	}

	d := NewDecoder(f)

	wh, err := d.ReadHeaders()
	if err != nil {
		fmt.Println(err.Error())
		t.Error("An error occurred when reading headers from ../8bit.wav")
	}

	if wh.BitsPerSample != 8 {
		t.Errorf("The bits per sample is not 8. Actual %d", wh.BitsPerSample)
	}
}

func TestDecoder_ReadHeaders_16Bit(t *testing.T) {
	f, err := os.Open("../kick-16b441k.wav")
	if err != nil {
		t.Error("Test failed when reading ../kick-16b441k.wav")
	}

	d := NewDecoder(f)

	wh, err := d.ReadHeaders()
	if err != nil {
		fmt.Println(err.Error())
		t.Error("An error occurred when reading headers from ../kick-16b441k.wav")
	}

	if wh.BitsPerSample != 16 {
		t.Errorf("The bits per sample is not 8. Actual %d", wh.BitsPerSample)
	}
}

func TestDecoder_ReadHeaders_32Bit(t *testing.T) {
	f, err := os.Open("../32bit.wav")
	if err != nil {
		t.Error("Test failed when reading ../32bit.wav")
	}

	d := NewDecoder(f)

	wh, err := d.ReadHeaders()
	if err != nil {
		fmt.Println(err.Error())
		t.Error("An error occurred when reading headers from ../32bit.wav")
	}

	if wh.BitsPerSample != 32 {
		t.Errorf("The bits per sample is not 8. Actual %d", wh.BitsPerSample)
	}
}

func TestDecoder_ReadHeaders_WithListInfo(t *testing.T) {
	f, err := os.Open("../listinfo.wav")
	if err != nil {
		t.Error("Test failed when reading ../listinfo.wav")
	}

	d := NewDecoder(f)

	_, err = d.ReadHeaders()
	if err != nil {
		fmt.Println(err.Error())
		t.Error("An error occurred when reading headers from ../listinfo.wav")
	}

	if d.metadata.Genre != "genre" {
		t.Errorf("Genre was not expected value. Acutal %s", d.metadata.Genre)
	}
}

func TestDecoder_ReadHeaders_WithCue(t *testing.T) {
	f, err := os.Open("../cue.wav")
	if err != nil {
		t.Error("Test failed when reading ../cue.wav")
	}

	d := NewDecoder(f)

	_, err = d.ReadHeaders()
	if err != nil {
		fmt.Println(err.Error())
		t.Error("An error occurred when reading headers from ../cue.wav")
	}

	if len(d.metadata.CuePoints) != 1 {
		t.Errorf("Expected 1 cue point. Acutal %d", len(d.metadata.CuePoints))
	}

	if d.metadata.CuePoints[0].ID != "\u0001" {
		t.Errorf("Expected cue point ID to be A. Actual %s", d.metadata.CuePoints[0].ID)
	}
}

//func TestDecoder_ReadHeaders(t *testing.T) {
//	f, err := os.Open("../listinfo.wav")
//	if err != nil {
//		panic(err)
//	}
//
//	d := NewDecoder(f)
//
//	h, err := d.ReadHeaders()
//
//	if err != nil {
//		panic(err)
//	}
//
//	if strings.Compare(d.metadata.Artist, "artist") == 0 {
//		t.Errorf("Failed - artist not found. Actual %s | %d", d.metadata.Artist, len(d.metadata.Artist))
//	}
//
//	if h.FileDescriptor != RIFF_CHUNK_ID {
//		t.Errorf("Failed")
//	}
//
//	if h.FileFormat != WAVE_FILE_FORMAT {
//		t.Errorf("Failed")
//	}
//}
//
//func TestDecoder_ToPCMStart(t *testing.T) {
//	f, err := os.Open("../recording-3.wav")
//	if err != nil {
//		panic(err)
//	}
//
//	d := NewDecoder(f)
//
//	_, err = d.ToPCMStart()
//	if err != nil {
//		panic(err)
//	}
//
//	c, err := d.Seek(0, 1)
//	if err != nil {
//		panic(err)
//	}
//
//	if c != 44 {
//		t.Errorf("nah dawg")
//	}
//}
//
//func TestDecoder_ToPCMStart_AfterReadHeaders(t *testing.T) {
//	f, err := os.Open("../recording-3.wav")
//	if err != nil {
//		panic(err)
//	}
//
//	d := NewDecoder(f)
//
//	_, err = d.ReadHeaders()
//	if err != nil {
//		panic(err)
//	}
//
//	_, err = d.ToPCMStart()
//	if err != nil {
//		panic(err)
//	}
//
//	c, err := d.Seek(0, 1)
//	if err != nil {
//		panic(err)
//	}
//
//	if c != 44 {
//		t.Errorf("nah dawg | %d", c)
//	}
//}
//
//func TestDecoder_Read(t *testing.T) {
//	f, err := os.Open("../recording-3.wav")
//	if err != nil {
//		panic(err)
//	}
//
//	d := NewDecoder(f)
//
//	h, err := d.ToPCMStart()
//	if err != nil {
//		panic(err)
//	}
//
//	l := 10000
//	b := make([]byte, l)
//	i, err := d.Read(b, h)
//	if err != nil {
//		panic(err)
//	}
//
//	if len(i) != l {
//		t.Error("i does not equal l")
//	}
//
//	//if len(b) != l {
//	//	t.Error("data len does not equal l")
//	//}
//}
