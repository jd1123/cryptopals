package exercises

import (
	"crypto/rand"
	"fmt"
	"io"

	"github.com/jd1123/cryptopals/aes"
	"github.com/jd1123/cryptopals/text"
)

func EncryptionOracle(pt []byte) []byte {
	rn := aes.RandInt(2)
	numPre := aes.RandInt(5) + 5
	numPost := aes.RandInt(5) + 5
	pre := make([]byte, numPre)
	post := make([]byte, numPost)

	if _, err := io.ReadFull(rand.Reader, pre); err != nil {
		fmt.Println(err)
	}
	if _, err := io.ReadFull(rand.Reader, post); err != nil {
		fmt.Println(err)
	}
	pt = append(pt, post...)
	pt = append(pre, pt...)
	key := aes.RandomKey()
	plainText := text.NewPlaintext(pt)

	if rn == 0 {
		ct, _ := plainText.EncryptCBC(key, nil)
		return ct
	} else {
		return plainText.EncryptECB(key)
	}
}

func Ex2_11() {
	secret := []byte("yellow submarineyellow submarineyellow submarineyellow submarineyellow submarineyellow submarineyellow submarine")
	ct := EncryptionOracle(secret)
	fmt.Println(aes.DetectECB(ct))
}
