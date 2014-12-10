package main

import (
	"crypto/rand"
	"encoding/binary"
)

func main() {
	var n int32
	binary.Read(rand.Reader, binary.LittleEndian, &n)
}
