package text

import (
	"github.com/jd1123/cryptopals/aes"
	"github.com/jd1123/cryptopals/padding"
	"github.com/jd1123/cryptopals/xor"
)

type Plaintext struct {
	plaintext []byte
	blocks    [][]byte
	blockSize int
}

func NewPlaintext(b []byte) Plaintext {
	return Plaintext{plaintext: b, blockSize: 0}
}

func (p *Plaintext) SetBlockSize(blockSize int) {
	p.blocks = breakBlocks(padding.PKCS7(p.plaintext, blockSize), blockSize)
	p.blockSize = blockSize
}

func (p *Plaintext) EncryptCBC(key []byte) []byte {
	p.SetBlockSize(16)
	iv := make([]byte, p.blockSize)
	ct := make([][]byte, len(p.blocks))
	for i := range p.blocks {
		if i == 0 {
			ct[i] = aes.ECBEncrypt(xor.XOR1(p.blocks[i], iv), key)
		} else {
			ct[i] = aes.ECBEncrypt(xor.XOR1(p.blocks[i], ct[i-1]), key)
		}
	}
	return assembleBlocks(ct)
}
