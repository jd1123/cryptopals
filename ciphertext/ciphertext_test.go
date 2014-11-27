package ciphertext

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/jd1123/cryptopals/freq"
	"github.com/jd1123/cryptopals/xor"
)

func TestCipher(t *testing.T) {
	plaintext := []byte("This is a super secret message")
	key := []byte("AbC")
	fmt.Println(plaintext)
	ct := xor.XORRepeatingKey(plaintext, key)
	fmt.Println(string(ct))
	fmt.Println(string(xor.XORRepeatingKey(ct, key)))
}

func TestRepeatingKey(t *testing.T) {
	f, _ := os.Open("../repeatingkey.txt")
	b, _ := ioutil.ReadAll(f)
	c := NewCiphertextFromBase64(b)
	c.DetermineKeyLength()
	fmt.Println(string(c.BreakVigenere()))
}

func TestAlphabet(t *testing.T) {
	sum := 0.0
	for _, v := range freq.FREQ {
		sum += v
	}
	fmt.Println(sum)
}
