package ciphertext

import (
	"encoding/base64"

	"github.com/jd1123/cryptopals/ecb"
	"github.com/jd1123/cryptopals/xor"
)

type Ciphertext struct {
	ciphertext []byte
	blockSize  int
	blocks     [][]byte
}

func NewCiphertextFromBase64(data []byte) Ciphertext {
	decoded, _ := base64.StdEncoding.DecodeString(string(data))
	return Ciphertext{ciphertext: decoded, blockSize: 0, blocks: nil}
}

func NewCiphertext(data []byte, blockSize int) Ciphertext {
	if blockSize <= 0 {
		return Ciphertext{ciphertext: data, blockSize: 0, blocks: nil}
	} else {
		blocks := breakBlocks(data, blockSize)
		return Ciphertext{ciphertext: data, blockSize: blockSize, blocks: blocks}
	}
}

func (c Ciphertext) GetCiphertext() []byte {
	return c.ciphertext
}

func (c *Ciphertext) ChangeBlockSize(blockSize int) {
	c.blocks = breakBlocks(c.ciphertext, blockSize)
	c.blockSize = blockSize
}

func (c *Ciphertext) DecodeWithRepeatingKey(key []byte) string {
	return string(xor.XORRepeatingKey(c.ciphertext, key))
}

func (c *Ciphertext) DetermineKeyLength() {
	blockSize := fullKeyLengthTest(*c, 2, 40)
	c.ChangeBlockSize(blockSize)
}

func (c *Ciphertext) BreakVigenere() []byte {
	c.DetermineKeyLength()
	key := make([]byte, c.blockSize)
	tBlocks := transposeBlocks(c.blocks)
	for i := range tBlocks {
		key[i] = breakSingleKey(tBlocks[i])
	}
	return key
}

func (c *Ciphertext) DecryptECB(key []byte) []byte {
	c.ChangeBlockSize(16)
	pt := make([]byte, 0)
	for i := range c.blocks {
		pt = append(pt, ecb.ECBDecrypt(c.blocks[i], key)...)
	}
	return pt
}
