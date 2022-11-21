package main

import (
	"crypto/rand"
	"fmt"
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

func main() {
	r := MakeRamdomID(16)
	fmt.Println(r)
}
