package main

import (
	"fmt"

	"github.com/jd1123/cryptopals/padding"
)

func main() {
	pt := []byte("Yellow submarine")
	fmt.Println("Len of pt:", len(pt))
	fmt.Println(padding.PKCS7(pt, 16))
}
