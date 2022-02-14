package Encoder

func (e *Encoder) calculateChunkSize() int {
	return 0
}

//  == SampleRate * NumChannels * BitsPerSample/8
func calculateByteRate(sampleRate uint32, numChannels uint16, bitsPerSample uint16) uint32 {
	return sampleRate * uint32(numChannels) * (uint32(bitsPerSample) / 8)
}

// == NumChannels * BitsPerSample/8
func calculateBlockAlign(numChannels uint16, bitsPerSample uint16) uint16 {
	return numChannels * (bitsPerSample / 8)
}

func calculateDataChunkSize(numChannels uint16, bitsPerSample uint16, framesWritten int) int {
	return (int(bitsPerSample) / 8) * int(numChannels) * framesWritten
}

func calculateFrameCount(p []int, numChannels uint16) int {
	return len(p) / int(numChannels) // length of data / number of channels
}
