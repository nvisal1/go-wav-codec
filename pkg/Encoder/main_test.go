package Encoder

import (
	"io"
	"os"
	"testing"
	"wav-concat/pkg/Decoder"
)

func TestEncoder_WriteAudioData_Close(t *testing.T) {

	f, err := os.Create("./TestEncoder_WriteAudioData_Close.wav")

	e, err := NewEncoder(1, 2, 48000, 16, f)
	if err != nil {
		t.Error(err.Error())
	}

	f2, err := os.Open("../../assets/recording-2.wav")
	if err != nil {
		t.Error(err.Error())
	}

	defer f.Close()
	defer f2.Close()

	d := Decoder.NewDecoder(f2)
	err = d.ReadMetadata()
	if err != nil {
		t.Error(err.Error())
	}

	a := make([]int, 0)
	c := 0
	ad, err := d.ReadAudioData(100, 0)
	if err != nil {
		t.Error(err.Error())
	}
	c += 1
	a = append(a, ad...)

	for {
		ad, err = d.ReadAudioData(100, 1)
		if err != nil {
			if err == io.EOF {
				a = append(a, ad...)
				break
			}
			t.Error(err.Error())
		}
		a = append(a, ad...)
	}

	err = e.WriteAudioData(a, 0)
	if err != nil {
		t.Error(err.Error())
	}

	err = e.Close()
	if err != nil {
		t.Error(err.Error())
	}

}
