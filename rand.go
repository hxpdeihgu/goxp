package goxp

import (
	"math/rand"
)

func (this *Model) Rand(len int8) string {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return string(b)
}

