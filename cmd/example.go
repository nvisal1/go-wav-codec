package main

import (
	"fmt"
	"os"
	wav "wav-concat/pkg"
)

/**
Library functions
1. Concat two wav files --
2. Store all wav headers in memory
3. Optional pull all data into memory
4. Easily iterate over pcm data - and allow user to specify batch size
5. Mix two wav files together - overlap them
6. Do the best we can to update headers when they do not match
7. Allow for multiple files to concat and mix

Wav file specification
http://soundfile.sapp.org/doc/WaveFormat/
*/

func main() {
	path1 := "./recording-1.wav"
	//path2 := "./recording-2.wav"

	//w1, w2, w3, err := wav.Concat(path1, path2)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(w1.Subchunk2Size)
	//fmt.Println(w2.Subchunk2Size)
	//fmt.Println(w3.Subchunk2Size)
	//
	//fmt.Println(len(w1.Data))
	//fmt.Println(len(w2.Data))
	//fmt.Println(len(w3.Data))
	//
	//err = w3.WriteTo("./recording-3.wav")
	//if err != nil {
	//	panic(err)
	//}

	f, _ := os.Open(path1)
	w := wav.NewWavReader(f)
	w1, _ := w.ReadHeaders()
	fmt.Println("====")
	fmt.Println(w1.BitsPerSample)
	fmt.Println(w1.AudioFormat)
	fmt.Println(w1.FileFormat)
}
