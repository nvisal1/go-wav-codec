# WAV Codec

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

`WriteAudioDataHeader`

`WriteAudioData`

`Close`

### Examples

#### Create a new Encoder
```
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
```
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
```
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
```
f, err := os.Open("./path/to/your/file")
if err != nil {
    panic(err)
}

defer f.Close()

d := Decoder.NewDecoder(f)
```

#### Read a portion of audio data
```
a := make([]int, 0)
ad, err := d.ReadAudioData(100, 0)
if err != nil {
    panic(err)
}
a = append(a, ad...)
```

#### Read all the audio data in chunks
```
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


