package exercises

import (
	"encoding/hex"
	"fmt"

	"github.com/jd1123/cryptopals/freq"
	"github.com/jd1123/cryptopals/xor"
)

var alphebet = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

var FREQ = map[byte]float64{
	'E': .1202,
	'T': .0910,
	'A': .0812,
	'O': .0768,
	'I': .0731,
	'N': .0695,
	'S': .0628,
	'R': .0602,
	'H': .0592,
	'D': .0432,
	'L': .0398,
	'U': .0288,
	'C': .0271,
	'M': .0261,
	'F': .0230,
	'Y': .0211,
	'W': .0209,
	'G': .0203,
	'P': .0182,
	'B': .0149,
	'V': .0111,
	'K': .0069,
	'X': .0017,
	'Q': .0011,
	'J': .0010,
	'Z': .0007,
}

func Ex1_3() {
	h1, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	stuff := make([][]byte, 256)
	key_byte := uint8(0)
	for i := 0; i < 256; i++ {
		key := make([]byte, len(h1))
		for j := 0; j < len(h1); j++ {
			key[j] = key_byte + uint8(i)
		}
		result := xor.XORSingleChar(h1, key[0])
		stuff[i] = result
	}
	fmt.Printf("%s\n", stuff[freq.ScoreList(stuff)])
}
