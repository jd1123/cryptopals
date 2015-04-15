package text

import (
	"encoding/base64"
	"fmt"

	"github.com/jd1123/cryptopals/aes"
	"github.com/jd1123/cryptopals/padding"
	"github.com/jd1123/cryptopals/xor"
)

type Ciphertext struct {
	ciphertext []byte
	blockSize  int
	blocks     [][]byte
}

func (c *Ciphertext) GetCt() []byte {
	return c.ciphertext
}

func NewCiphertextFromBase64(data []byte) (Ciphertext, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return Ciphertext{}, err
	}
	return Ciphertext{ciphertext: decoded, blockSize: 0, blocks: nil}, nil
}

func NewCiphertext(data []byte, blockSize int) Ciphertext {
	if blockSize <= 0 {
		return Ciphertext{ciphertext: data, blockSize: 0, blocks: nil}
	} else {
		blocks := BreakBlocks(data, blockSize)
		return Ciphertext{ciphertext: data, blockSize: blockSize, blocks: blocks}
	}
}

func (c Ciphertext) GetCiphertext() []byte {
	return c.ciphertext
}

func (c *Ciphertext) ChangeBlockSize(blockSize int) {
	c.blocks = BreakBlocks(c.ciphertext, blockSize)
	c.blockSize = blockSize
}

func (c *Ciphertext) DecodeWithRepeatingKey(key []byte) string {
	return string(xor.XORRepeatingKey(c.ciphertext, key))
}

func (c *Ciphertext) DetermineKeyLength() {
	blockSize := fullKeyLengthTest(*c, 2, 40)
	fmt.Println("score 29", keyLengthTest(*c, 29))
	fmt.Println("score 2", keyLengthTest(*c, 2))
	c.ChangeBlockSize(blockSize)
}

func (c *Ciphertext) DecryptCBC(key, iv []byte) []byte {
	c.ChangeBlockSize(16)
	iv = c.blocks[0]
	c.ciphertext = c.ciphertext[16:]
	c.ChangeBlockSize(16)
	numBlocks := len(c.blocks)
	pt := make([][]byte, numBlocks)
	for i := range c.blocks {
		if i == 0 {
			pt[i] = xor.XOR1(aes.ECBDecrypt(c.blocks[i], key), iv)
		} else {
			pt[i] = xor.XOR1(aes.ECBDecrypt(c.blocks[i], key), c.blocks[i-1])
		}
	}
	result, _ := padding.ValidatePKCS7(AssembleBlocks(pt), 16)
	return result
}

func (c *Ciphertext) BreakVigenere() []byte {
	c.DetermineKeyLength()
	key := make([]byte, c.blockSize)
	tBlocks := TransposeBlocks(c.blocks)
	for i := range tBlocks {
		key[i] = breakSingleKey(tBlocks[i])
	}
	return key
}

func (c *Ciphertext) DecryptECB(key []byte) []byte {
	c.ChangeBlockSize(16)
	pt := make([]byte, 0)
	for i := range c.blocks {
		pt = append(pt, aes.ECBDecrypt(c.blocks[i], key)...)
	}
	return pt
}

func (c *Ciphertext) CheckRepeatedBlocks(blockSize int) int {
	c.ChangeBlockSize(blockSize)
	count := 0
	for i := 0; i < len(c.blocks)-1; i++ {
		for j := i + 1; j < len(c.blocks); j++ {
			if BlocksEqual(c.blocks[i], c.blocks[j]) {
				count++
			}
		}
	}
	return count
}
