package exercises

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jd1123/cryptopals/text"
	"github.com/jd1123/cryptopals/xor"
)

func Ex1_6() {
	f, err := os.Open("data/repeatingkey.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c, _ := ioutil.ReadAll(f)
	ct := text.NewCiphertextFromBase64(c)
	key := ct.BreakVigenere()
	fmt.Println(string(xor.XORRepeatingKey(ct.GetCiphertext(), key)))
	fmt.Println("Key: ", key)
}
