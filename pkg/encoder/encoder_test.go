package encoder

import (
	"io"
	"os"
	"testing"

	"github.com/nvisal1/go-wav-codec/pkg/decoder"
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

	d := decoder.NewDecoder(f2)
	err = d.ReadMetadata()
	if err != nil {
		t.Error(err.Error())
	}

	a := make([]int, 0)
	ad, err := d.ReadAudioData(100, 0)
	if err != nil {
		t.Error(err.Error())
	}
	a = append(a, ad...)

	for {
		ad, err = d.ReadAudioData(100, 1)
		if err != nil {
			if err == io.EOF {
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

func TestEncoder_WriteMetadata_WriteAudioData_Close(t *testing.T) {
	f, err := os.Create("./TestEncoder_WriteAudioData_Close-2.wav")

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

	d := decoder.NewDecoder(f2)
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
	c++
	a = append(a, ad...)

	for {
		ad, err = d.ReadAudioData(100, 1)
		if err != nil {
			if err == io.EOF {
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

	ic := &InfoChunk{
		Location:     "",
		Artist:       "Shinedow",
		Software:     "",
		CreationDate: "",
		Copyright:    "",
		Title:        "",
		Engineer:     "",
		Genre:        "",
		Product:      "",
		Source:       "",
		Subject:      "",
		Comments:     "",
		Technician:   "",
		Keywords:     "",
		Medium:       "",
	}

	e.WriteMetadata(ic)

	err = e.Close()
	if err != nil {
		t.Error(err.Error())
	}
}
