package aes

import (
	"bytes"
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

func RandIntR(b, e int) int {
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(e-b)))
	rn := int(r.Int64()) + b
	return rn
}

func DetectECB(ct []byte) bool {
	for i := 0; i < 34; i++ {
		if bytes.Equal(ct[i:i+16], ct[i+16:i+32]) {
			return true
		}
	}
	return false
}
