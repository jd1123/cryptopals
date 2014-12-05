package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jd1123/cryptopals/ciphertext"
)

const AES128_BLOCK_SIZE = 16

func main() {
	count := 0
	best := 0
	f, _ := os.Open("data/8.txt")
	b, _ := ioutil.ReadAll(f)
	lines := strings.Split(string(b), "\n")
	for i := 0; i < len(lines); i++ {
		c := ciphertext.NewCiphertextFromBase64([]byte(lines[i]))
		n := c.CheckRepeatedBlocks(AES128_BLOCK_SIZE)
		if count < n {
			count = n
			best = i
		}
	}
	fmt.Println("There are", len(lines), "ciphertexts")
	fmt.Println("Most likely:", best, "with", count, "same blocks")
}
