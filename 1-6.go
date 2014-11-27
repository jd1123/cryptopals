package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jd1123/cryptopals/ciphertext"
	"github.com/jd1123/cryptopals/xor"
)

func main() {
	f, _ := os.Open("repeatingkey.txt")
	c, _ := ioutil.ReadAll(f)
	ct := ciphertext.NewCiphertextFromBase64(c)
	key := ct.BreakVigenere()
	fmt.Println(string(xor.XORRepeatingKey(ct.GetCiphertext(), key)))
	fmt.Println("Key: ", string(key))
}
