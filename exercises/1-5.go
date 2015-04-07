package exercises

import (
	"fmt"

	"github.com/jd1123/cryptopals/xor"
)

var stanza = []byte("Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal")

func Ex1_5() {
	key := []byte("ICE")
	fmt.Printf("%x\n", xor.XORRepeatingKey(stanza, key))
}
