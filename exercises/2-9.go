package exercises

import (
	"fmt"

	"github.com/jd1123/cryptopals/padding"
)

func Ex2_9() {
	pt := []byte("Yellow submarine")
	fmt.Println("Len of pt:", len(pt))
	fmt.Println(padding.PKCS7(pt, 16))
}
