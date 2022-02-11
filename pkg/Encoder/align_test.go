package Encoder

import "testing"

func TestAlign(t *testing.T) {
	b := make([]byte, 1)
	b2 := Align(b)

	if len(b2) != 2 {
		t.Error("B2: Length is not 2")
	}

	b3 := Align(b2)
	if len(b3) != 2 {
		t.Error("B3: Length is not 2")
	}

	b4 := make([]byte, 3)
	b5 := Align(b4)
	if len(b5) != 4 {
		t.Error("B5: Length is not 4")
	}

}
