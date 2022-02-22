# WAV Codec

[![Go](https://github.com/nvisal1/Wav-Codec-POC/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/nvisal1/Wav-Codec-POC/actions/workflows/go.yml)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)
[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)
[![Open Source? Yes!](https://badgen.net/badge/Open%20Source%20%3F/Yes%21/blue?icon=github)](https://github.com/Naereen/badges/)
[![saythanks](https://img.shields.io/badge/say-thanks-ff69b4.svg)](https://saythanks.io/to/kennethreitz)

<img src="./assets/gopher.png" alt="drawing" width="500"/>

## Table of Contents
* Summary
* Installation
* Encoder
* Decoder
* Practical Examples
* Contributing

## Summary

This audio codec includes a simple API that makes it easy to read and write WAV files in Go!

It includes support for **reading** WAV files that include
* LIST
* adtl
* INFO
* labl
* note
* ltxt
* smpl
* fact
* plst
* cue
* inst

and support for **writing** Wav files with

* LIST
* INFO

_**(as most metadata chunks are frowned upon - they are not supported by many applications)**_

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

lic := &ListChunk{
    info: ic,
}

err = e.WriteMetadata(lic)
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

## Practical Examples



