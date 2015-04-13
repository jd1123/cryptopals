package exercises

import (
	"fmt"

	"github.com/jd1123/cryptopals/padding"
)

func Ex2_15() {
	b := []byte("This is a test of the PKCS7 padding")
	padded := padding.PKCS7(b, 16)
	validated, err := padding.ValidatePKCS7(padded, 16)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("Passed", string(validated))
	}
	padded[len(padded)-2] = byte(19)
	validated, err = padding.ValidatePKCS7(padded, 16)
	if err != nil {
		fmt.Println("error:", err)
	}

}
