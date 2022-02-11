package pkg

import (
	"io"
	"os"
	"testing"
)

func TestNewEncoder(t *testing.T) {
	f, err := os.Create("./test1.wav")
	if err != nil {
		panic(err)
	}
	e, err := NewEncoder(1, 2, 48000, 16, f)
	if err != nil {
		panic(err)
	}

	if e.NumChannels != 2 {
		t.Errorf("ERROR!")
	}

	if e.bytesWritten != 44 {
		t.Errorf("ERROR")
	}

	if e.framesWritten != 0 {
		t.Errorf("ERROR")
	}

	f.Close()
}

func TestEncoder_Write(t *testing.T) {

	f, err := os.Open("../recording-3.wav")
	if err != nil {
		panic(err)
	}

	d := NewDecoder(f)

	h, err := d.ToPCMStart()
	if err != nil {
		panic(err)
	}

	l := 10000
	var pcm []int

	for {
		b := make([]byte, l)
		i, err := d.Read(b, h)
		if err == io.EOF {
			break
		}
		pcm = append(pcm, i...)
	}

	f.Close()

	f, err = os.Create("./test2.wav")
	if err != nil {
		panic(err)
	}
	e, err := NewEncoder(1, 2, 48000, 16, f)
	if err != nil {
		panic(err)
	}

	pcmD := e.ToDelay(pcm)

	err = e.Write(pcmD)
	if err != nil {
		panic(err)
	}

	err = e.Close()
	if err != nil {
		panic(err)
	}

	f.Close()

	//t.Error(len(pcm))

}
