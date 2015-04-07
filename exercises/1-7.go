package exercises

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jd1123/cryptopals/text"
)

func Ex1_7() {
	f, _ := os.Open("data/7.txt")
	c, _ := ioutil.ReadAll(f)
	ct := text.NewCiphertextFromBase64(c)
	pt := ct.DecryptECB([]byte("YELLOW SUBMARINE"))
	fmt.Println(string(pt))
}
