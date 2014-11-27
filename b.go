package main

import (
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"code.google.com/p/go.crypto/twofish"
)

func main() {
	key := []byte("YELLOW SUBMARINE")
	c, _ := twofish.NewCipher(key)
	a, _ := cipher.NewGCM(c)
	pt := []byte("This is a super secret message. Let's see if it works!")
	//	ct := make([]byte, 0)
	nonce := make([]byte, a.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}
	t := a.Seal(nil, nonce, pt, nil)
	fmt.Println(t)
	//	t[3] = 0
	n, err := a.Open(nil, nonce, t, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(n))
	fmt.Println(nonce)
}
