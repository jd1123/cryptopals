package aes

import (
	"crypto/aes"

	"github.com/jd1123/cryptopals/padding"
)

func ECBEncrypt(pt, key []byte) []byte {
	c, err := aes.NewCipher(key)
	pt = padding.PKCS7(pt, 16)
	if err != nil {
		panic(err)
	}
	ct := make([]byte, 16)
	c.Encrypt(ct, pt)
	return ct
}

func ECBDecrypt(ct, key []byte) []byte {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	pt := make([]byte, 16)
	c.Decrypt(pt, ct)
	return pt
}
