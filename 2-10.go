package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jd1123/cryptopals/ciphertext"
)

func main() {
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)
	f, _ := os.Open("data/10.txt")
	buff, _ := ioutil.ReadAll(f)
	ct := ciphertext.NewCiphertextFromBase64(buff)
	pt := ct.DecryptCBC(key, iv)
	fmt.Println(string(pt))
}
