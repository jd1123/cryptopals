package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jd1123/cryptopals/ciphertext"
)

func main() {
	f, _ := os.Open("7.txt")
	c, _ := ioutil.ReadAll(f)
	ct := ciphertext.NewCiphertextFromBase64(c)
	pt := ct.DecryptECB([]byte("YELLOW SUBMARINE"))
	fmt.Println(string(pt))

}
