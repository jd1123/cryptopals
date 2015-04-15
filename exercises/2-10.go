package exercises

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jd1123/cryptopals/text"
)

func Ex2_10() {
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)
	f, _ := os.Open("data/10.txt")
	buff, _ := ioutil.ReadAll(f)
	ct, _ := text.NewCiphertextFromBase64(buff)
	pt := ct.DecryptCBC(key, iv)
	fmt.Println(string(pt))
}
