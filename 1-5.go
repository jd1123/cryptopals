package main

import (
	"fmt"

	"github.com/jd1123/cryptopals/xor"
)

var stanza = []byte("Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal")

func main() {
	key := []byte("ICE")
	fmt.Printf("%x\n", xor.XORRepeatingKey(stanza, key))
}
