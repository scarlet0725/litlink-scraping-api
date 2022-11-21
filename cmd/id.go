package cmd

import (
	"crypto/rand"
)

func MakeRamdomID(size uint32) string {
	const chars = "ABCDEFGHIJKLMNPQRSTWXYZ123456789"
	b := make([]byte, size)

	rand.Read(b)

	var result string
	for _, v := range b {
		result += string(chars[int(v)%len(chars)])
	}

	return result
}
