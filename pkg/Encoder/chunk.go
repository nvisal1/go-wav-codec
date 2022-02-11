package Encoder

import "io"

type Chunk interface {
	WriteTo(w io.Writer) (int, error)
}
