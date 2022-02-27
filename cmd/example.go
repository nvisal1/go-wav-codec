package main

import (
	"fmt"
	"io"
	"os"
	"wav-concat/pkg/decoder"
	"wav-concat/pkg/encoder"
)

// This example will copy all the audio data from one wav file
// and write it to a new wav file
func main() {
	p := "./assets/recording-1.wav"
	f, err := os.Open(p)
	if err != nil {
		panic(err)
	}

	d := decoder.NewDecoder(f)

	err = d.ReadMetadata()
	if err != nil {
		panic(err)
	}

	numSamples := d.Metadata.NumSamples

	chunkSize := int(numSamples / 10)

	buf := make([]int, 0)

	b, err := d.ReadAudioData(chunkSize, 0)
	if err != nil {
		panic(err)
	}

	buf = append(buf, b...)

	for {
		fmt.Println("Reading another chunk...")
		b, err = d.ReadAudioData(chunkSize, 1)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		buf = append(buf, b...)
	}

	f.Close()

	f2, err := os.Create("./assets/a-new-file.wav")
	e, err := encoder.NewEncoder(1, d.Metadata.FMT.NumChannels, d.Metadata.FMT.SampleRate, d.Metadata.FMT.BitsPerSample, f2)
	if err != nil {
		panic(err)
	}

	err = e.WriteAudioData(buf, 0)
	if err != nil {
		panic(err)
	}

	err = e.Close()
	if err != nil {
		panic(err)
	}

	f2.Close()

}
