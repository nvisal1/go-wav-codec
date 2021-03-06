package decoder

import (
	"os"
	"testing"
)

func TestDecoder_ReadMetadata_ReadAudioData(t *testing.T) {
	f, err := os.Open("../../assets/recording-1.wav")
	if err != nil {
		t.Error(err.Error())
	}
	d := NewDecoder(f)

	err = d.ReadMetadata()
	if err != nil {
		t.Error(err.Error())
	}

	if d.Metadata.FMT == nil {
		t.Error("fmt chunk is nil")
	}

	if d.Metadata.FMT.AudioFormat != 1 {
		t.Errorf("expected audio format \"1\". received audio format \"%d\"", d.Metadata.FMT.AudioFormat)
	}

	if d.Metadata.FMT.NumChannels != 2 {
		t.Errorf("expected num channels \"2\". received num channels \"%d\"", d.Metadata.FMT.NumChannels)
	}

	if d.Metadata.FMT.SampleRate != 48000 {
		t.Errorf("expected sample rate \"48000\". received sample rate \"%d\"", d.Metadata.FMT.SampleRate)
	}

	if d.Metadata.FMT.ByteRate != 192000 {
		t.Errorf("expected byte rate \"192000\". received byte rate \"%d\"", d.Metadata.FMT.ByteRate)
	}

	if d.Metadata.FMT.BlockAlign != 4 {
		t.Errorf("expected block align \"4\". received block align \"%d\"", d.Metadata.FMT.BlockAlign)
	}

	if d.Metadata.FMT.BitsPerSample != 16 {
		t.Errorf("expected bits per sample \"16\". received bits per sample \"%d\"", d.Metadata.FMT.BitsPerSample)
	}

	b, err := d.ReadAudioData(100, 0)
	if err != nil {
		t.Error(err.Error())
	}

	if len(b) != 100 {
		t.Errorf("expected buffer with length \"100\". received buffer with length \"%d\"", len(b))
	}

	p, err := f.Seek(0, 1)
	if err != nil {
		t.Error(err.Error())
	}

	// 100 * 2 == 100 samples * 2 bytes per sample
	if p != (d.Metadata.DataPosition + (100 * 2)) {
		t.Errorf("expected reader to be at position \"%d\". but reader is actually at position \"%d\"", d.Metadata.DataPosition+(100*2), p)
	}

}

func TestDecoder_ReadAudioData(t *testing.T) {
	f, err := os.Open("../../assets/recording-1.wav")
	if err != nil {
		t.Error(err.Error())
	}
	d := NewDecoder(f)

	b, err := d.ReadAudioData(100, 0)
	if err != nil {
		t.Error(err.Error())
	}

	if len(b) != 100 {
		t.Errorf("expected buffer with length \"100\". received buffer with length \"%d\"", len(b))
	}

	p, err := f.Seek(0, 1)
	if err != nil {
		t.Error(err.Error())
	}

	// 100 * 2 == 100 samples * 2 bytes per sample
	if p != (d.Metadata.DataPosition + (100 * 2)) {
		t.Errorf("expected reader to be at position \"%d\". but reader is actually at position \"%d\"", d.Metadata.DataPosition+(100*2), p)
	}
}
