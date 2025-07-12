package core

import "github.com/withzeus/mugi-identity/core/lib"

type Helper struct {
	lib.Encoder
}

func New() string {
	h := Helper{}
	s := h.GenerateEncodedString(&lib.Hex{}, 10)
	return s
}
