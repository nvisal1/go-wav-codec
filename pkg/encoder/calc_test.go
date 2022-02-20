package encoder

import "testing"

func TestCalculateByteRate(t *testing.T) {
	val := calculateByteRate(48000, 2, 16)
	if val != 192000 {
		t.Errorf("calculated byte rate was incorrect. Actual: %d", val)
	}
}

func TestCalculateBlockAlign(t *testing.T) {
	val := calculateBlockAlign(2, 16)
	if val != 4 {
		t.Errorf("calculated block align was incorrect. Actual: %d", val)
	}
}

func TestCalculateDataChunkSize(t *testing.T) {
	val := calculateDataChunkSize(2, 16, 100)
	if val != 400 {
		t.Errorf("calculated block align was incorrect. Actual: %d", val)
	}
}

func TestCalculateFrameCount(t *testing.T) {
	b := []int{0, 0, 0, 0}
	val := calculateFrameCount(b, 2)
	if val != 2 {
		t.Errorf("calculated block align was incorrect. Actual: %d", val)
	}
}
