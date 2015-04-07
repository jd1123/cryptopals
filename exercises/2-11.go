package exercises

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"

	"github.com/jd1123/cryptopals/text"
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

func EncryptionOracle(pt []byte) []byte {
	rn := RandInt(2)
	numPre := RandInt(5) + 5
	numPost := RandInt(5) + 5
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
	key := RandomKey()
	plainText := text.NewPlaintext(pt)

	if rn == 0 {
		return plainText.EncryptCBC(key, nil)
	} else {
		return plainText.EncryptECB(key)
	}
}

func DetectECB(ct []byte) bool {
	for i := 0; i < 34; i++ {
		if bytes.Equal(ct[i:i+16], ct[i+16:i+32]) {
			return true
		}
	}
	return false
}

func Ex2_11() {
	secret := []byte("yellow submarineyellow submarineyellow submarineyellow submarineyellow submarineyellow submarineyellow submarine")
	ct := EncryptionOracle(secret)
	fmt.Println(DetectECB(ct))
}
