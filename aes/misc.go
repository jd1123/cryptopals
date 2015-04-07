package aes

import (
	"crypto/rand"
	"io"
	"math/big"
)

func RandomKey() []byte {
	key := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil
	}
	return key
}

func RandInt(n int) int {
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(n)))
	rn := int(r.Int64())
	return rn
}
