# Go WAV Codec

[![Build](https://github.com/nvisal1/Wav-Codec-POC/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/nvisal1/Wav-Codec-POC/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/nvisal1/Go-Wav-Codec/branch/master/graph/badge.svg?token=4DRC08BB3G)](https://codecov.io/gh/nvisal1/Go-Wav-Codec)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)]()
[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)]()
[![Maintained?](https://img.shields.io/badge/Maintained%3F-yes-green.svg)]()
[![Open Source? Yes!](https://badgen.net/badge/Open%20Source%20%3F/Yes%21/blue?icon=github)]()
[![saythanks](https://img.shields.io/badge/say-thanks-ff69b4.svg)]()

<img src="./assets/gopher.png" alt="drawing" width="500"/>

## Table of Contents
* [Summary](#Summary)
* [Installation](#Installation)
* [Encoder](#Encoder)
* [Decoder](#Decoder)
* [Practical Examples](#Practical-Examples)
* [Expected Chunk Formats](#Expected-Chunk-Formats)
* [Contributing](#Contributing)
* [License](#License)

## Summary

This audio codec includes a simple API that makes it easy to read and write WAV files in Go!

It includes support for **reading** WAV files that include


| Chunk ID |                                                                                                                                            Description                                                                                                                                             |
|:--------:|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|
|   LIST   | Includes support for types ADTL and INFO. ADTL chunks can include <br/> LTXT, LABL and NOTE chunks<br/>Learn about [INFO](https://www.recordingblogs.com/wiki/list-chunk-of-a-wave-file) <br/>Learn about [ADTL](https://sites.google.com/site/musicgapi/technical-documents/wav-file-format#list) |
|   SMPL   |                                                                                                                                                301                                                                                                                                                 |
|   FACT   |                                                                                                                                                301                                                                                                                                                 |
|   PLST   |                                                                                                                                                301                                                                                                                                                 |
|   CUE    |                                                                                                                                                301                                                                                                                                                 |
|   INST   |                                                                                                                                                301                                                                                                                                                 |


and support for **writing** Wav files with

| Chunk ID |           Description           |
|:--------:|:-------------------------------:|
|   LIST   | Includes support for type INFO. |


_**(as most metadata chunks are frowned upon - they are not supported by many applications)**_

## Installation

Using the go wav codec is easy. First, use go get to install the latest version of the library. 

```bash
go get github.com/nvisal1/go-wav-codec
```

Next, include go-wav-codec in your application
```go
import "github.com/nvisal1/go-wav-codec"
```

## Encoder

`WriteMetadata`

`WriteAudioData`

`Close`

### Examples

#### Create a new Encoder
```go
f, err := os.Create("./path/to/your/file")
if err != nil {
    panic(err)
}

defer f.Close()

e, err := NewEncoder(1, 2, 48000, 16, f)
if err != nil {
    panic(err)
}
```

#### Write audio data to a new file
```go
a := []int{0,0,0,0} // This is your audio data

err = e.WriteAudioData(a, 0)
if err != nil {
    t.Error(err.Error())
}

err = e.Close()
if err != nil {
    t.Error(err.Error())
}
```

#### Write metadata to a new file
```go
a := []int{0,0,0,0} // This is your audio data

err = e.WriteAudioData(a, 0)
if err != nil {
    t.Error(err.Error())
}

ic := &InfoChunk{
		Location:     "",
		Artist:       "artist",
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

err = e.WriteMetadata(ic)
if err != nil {
    t.Error(err

err = e.Close()
if err != nil {
    t.Error(err.Error())
}
```


## Decoder

`ReadMetadata`

`ReadAudioData`

### Examples

#### Create a new Decoder
```go
f, err := os.Open("./path/to/your/file")
if err != nil {
    panic(err)
}

defer f.Close()

d := Decoder.NewDecoder(f)
```

#### Read a portion of audio data
```go
a := make([]int, 0)
ad, err := d.ReadAudioData(100, 0)
if err != nil {
    panic(err)
}
a = append(a, ad...)
```

#### Read all the audio data in chunks
```go
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
        panic(err)
    }
    a = append(a, ad...)
}
```

## License

Go Wav Codec is released under the MIT license. See [LICENSE](https://github.com/nvisal1/Go-Wav-Codec/blob/master/LICENSE)