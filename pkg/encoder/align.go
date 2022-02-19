package Encoder

func Align(b []byte) []byte {
	if len(b)%2 != 0 {
		b = append(b, 0)
	}
	return b
}
