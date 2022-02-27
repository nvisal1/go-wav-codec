# Go WAV Codec

[![Test](https://github.com/nvisal1/Go-Wav-Codec/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/nvisal1/Go-Wav-Codec/actions/workflows/test.yml)
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
* [Resources](#Resources)
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

## Expected Chunk Formats

#### [fmt](https://sites.google.com/site/musicgapi/technical-documents/wav-file-format#fmt)

| Size (bytes) |                Description                 |                           Value                           |
|:------------:|:------------------------------------------:|:---------------------------------------------------------:|
|      4       |                  Chunk ID                  |          "fmt " **(this library case insensitive)**           |
|      4       |                 Chunk Size                 | 16 **(this library does not support extra format bytes)** |
|      2       |      Audio Format (Compression Code)       |          1 **(this library only supports PCM)**           |
|      2       |             Number of Channels             |                        1 - 65,535                         |
|      4       |                Sample Rate                 |                      1 - 0xFFFFFFFF                       |
|      4       |              Bytes Per Second              |                      1 - 0xFFFFFFFF                       |
|      2       |                Block Align                 |                        1 - 65,535                         |
|      2       |        Bits Per Sample (Bit Depth)         |                        2 - 65,535                         |
|     N/A      |_**Extra Format Bytes are not supported**_  |                            N/A                            |

#### [fact](https://sites.google.com/site/musicgapi/technical-documents/wav-file-format#fact)

| Size (bytes) |    Description    |                                          Value                                          |
|:------------:|:-----------------:|:---------------------------------------------------------------------------------------:|
|      4       |     Chunk ID      |                       "fact" **(this library case insensitive)**                        |
|      4       |    Chunk Size     |                  4 **(this library only supports number of samples)**                   |
|      4       | Number of Samples |                                     1 - 0xFFFFFFFF                                      | 

#### [cue](https://sites.google.com/site/musicgapi/technical-documents/wav-file-format#cue)

| Size |       Description        |                           Value                            |
|:----:|:------------------------:|:----------------------------------------------------------:|
|  4   |         Chunk ID         |         "cue " **(this library case insensitive)**         |
|  4   |        Chunk Size        |                  4 + (NumCuePoints * 24)                   |
|  4   |   Number of Cue Points   |                    number of cue points                    |
| N/A  |**Cue Points Start Here** |                            N/A                             |

###### cue point

| Size |  Description  |                                        Value                                         |
|:----:|:-------------:|:------------------------------------------------------------------------------------:|
|  4   |      ID       |                             unique identification value                              |
|  4   |   Position    |                                 play order position                                  |
|  4   | Data Chunk ID |                         RIFF ID of corresponding data chunk                          |
|  4   |  Chunk Start  |                              	Byte Offset of Data Chunk                              |
|  4   |  Block Start  |                        Byte Offset to sample of First Channel                        |
|  4   | Sample Offset |                     Byte Offset to sample byte of First Channel                      |

#### [plst](https://sites.google.com/site/musicgapi/technical-documents/wav-file-format#plst)

| Size |    Description     |                   Value                    |
|:----:|:------------------:|:------------------------------------------:|
|  4   |      Chunk ID      | "plst" **(this library case insensitive)** |
|  4   |     Chunk Size     |          Number of Segments * 12           |
|  4   | Number of Segments |               1 - 0xFFFFFFFF               | 

###### segment

| Size |     Description     |     Value      |
|:----:|:-------------------:|:--------------:|
|  4   |    Cue Point ID     | 0 - 0xFFFFFFFF |
|  4   | Length (in samples) | 1 - 0xFFFFFFFF |
|  4   |  Number of Repeats  | 1 - 0xFFFFFFFF | 


#### list

| Size | Description |                        Value                         |
|:----:|:-----------:|:----------------------------------------------------:|
|  4   |  Chunk ID   |      "list" **(this library case insensitive)**      |
|  4   | Chunk Size  |                   depends on type                    |
|  4   |   Type ID   | "adtl" or "info" **(this library case insensitive)** | 

#### [labl](https://sites.google.com/site/musicgapi/technical-documents/wav-file-format#labl)

| Size | Description  |                   Value                    |
|:----:|:------------:|:------------------------------------------:|
|  4   |   Chunk ID   | "labl" **(this library case insensitive)** |
|  4   |  Chunk Size  |              depends on text               |
|  4   | Cue Point ID |               0 - 0xFFFFFFFF               | 
| N/A  |     Text     |                    N/A                     |

#### [ltxt](https://sites.google.com/site/musicgapi/technical-documents/wav-file-format#ltxt)

| Size |  Description  |                   Value                    |
|:----:|:-------------:|:------------------------------------------:|
|  4   |   Chunk ID    | "ltxt" **(this library case insensitive)** |
|  4   |  Chunk Size   |              depends on text               |
|  4   | Cue Point ID  |               0 - 0xFFFFFFFF               |
|  4   | Sample Length |               0 - 0xFFFFFFFF               | 
|  4   |  Purpose ID   |               0 - 0xFFFFFFFF               | 
|  2   |    Country    |                 0 - 0xFFFF                 |
|  2   |   Language    |                 0 - 0xFFFF                 | 
|  2   |    Dialect    |                 0 - 0xFFFF                 | 
|  2   |   Code Page   |                 0 - 0xFFFF                 | 
| N/A  |     Text      |                    N/A                     | 

#### [note](https://sites.google.com/site/musicgapi/technical-documents/wav-file-format#note)

| Size | Description  |                   Value                    |
|:----:|:------------:|:------------------------------------------:|
|  4   |   Chunk ID   | "note" **(this library case insensitive)** |
|  4   |  Chunk Size  |              depends on text               |
|  4   | Cue Point ID |               0 - 0xFFFFFFFF               | 
| N/A  |     Text     |                    N/A                     |

#### [smpl](https://sites.google.com/site/musicgapi/technical-documents/wav-file-format#smpl)

| Size |         Description          |                       Value                       |
|:----:|:----------------------------:|:-------------------------------------------------:|
|  4   |           Chunk ID           |    "smpl" **(this library case insensitive)**     |
|  4   |          Chunk Size          | 36 + (Number of Sample Loops * 24) + Sampler Data |
|  4   |         Manufacturer         |                  0 - 0xFFFFFFFF                   | 
|  4   |           Product            |                  0 - 0xFFFFFFFF                   | 
|  4   |        Sample Period         |                  0 - 0xFFFFFFFF                   | 
|  4   |       MIDI Unity Note        |                      0 - 127                      | 
|  4   |     MIDI Pitch Fraction      |                  0 - 0xFFFFFFFF                   | 
|  4   |         SMPTE Format         |                 0, 24, 25, 29, 30                 | 
|  4   |         SMPTE Offset         |                  0 - 0xFFFFFFFF                   | 
|  4   |    Number of Sample Loops    |                  0 - 0xFFFFFFFF                   | 
|  4   |         Sampler Data         |                  0 - 0xFFFFFFFF                   |
| N/A  | **Sampler Loops Start Here** |                        N/A                        | 

###### sampler loop

| Size |         Description          |      Value      |
|:----:|:----------------------------:|:---------------:|
|  4   |         Cue Point ID         | 0 - 0xFFFFFFFF  |
|  4   |             Type             | 0 - 0xFFFFFFFF  |
|  4   |            Start             | 0 - 0xFFFFFFFF  | 
|  4   |             End              | 0 - 0xFFFFFFFF  | 
|  4   |           Fraction           | 0 - 0xFFFFFFFF  | 
|  4   |          Play Count          | 0 - 0xFFFFFFFF  | 

#### [inst](https://sites.google.com/site/musicgapi/technical-documents/wav-file-format#inst)

| Size |  Description   |                                     Value                                      |
|:----:|:--------------:|:------------------------------------------------------------------------------:|
|  4   |    Chunk ID    |                   "inst" **(this library case insensitive)**                   |
|  4   |   Chunk Size   |                                       7                                        |
|  1   | Unshifted Note |                                    0 - 127                                     |
|  1   | Fine Tune (dB) |                                   -50 - +50                                    |
|  1   |      Gain      |                                   -64 - +64                                    |
|  1   |    Low Note    |                                    0 - 127                                     |
|  1   |   High Note    |                                    0 - 127                                     |
|  1   |  Low Velocity  |                                    1 - 127                                     |
|  1   | High Velocity  |                                    1 - 127                                     |

## Resources

|    Website     |         Description         |                                         Link                                          |
|:--------------:|:---------------------------:|:-------------------------------------------------------------------------------------:|
|   musicg-api   |       Wav File Format       |  [Here](https://sites.google.com/site/musicgapi/technical-documents/wav-file-format)  |
| recordingblogs | List Chunk (of a RIFF file) |         [Here](https://www.recordingblogs.com/wiki/list-chunk-of-a-wave-file)         |
|   soundfile    |  WAVE PCM soundfile format  |                   [Here](http://soundfile.sapp.org/doc/WaveFormat/)                   |


## License

Go Wav Codec is released under the MIT license. See [LICENSE](https://github.com/nvisal1/Go-Wav-Codec/blob/master/LICENSE)
