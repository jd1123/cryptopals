package main

import (
	"encoding/hex"
	"fmt"

	"github.com/jd1123/cryptopals/xor"
)

func main() {
	h1, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	h2, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	result, _ := xor.XOR(h1, h2)
	fmt.Printf("%x", result)
}
