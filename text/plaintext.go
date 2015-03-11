package text

import (
	"crypto/rand"
	"io"

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
	p.blocks = BreakBlocks(padding.PKCS7(p.plaintext, blockSize), blockSize)
	p.blockSize = blockSize
}

func (p *Plaintext) EncryptCBC(key, iv []byte) []byte {
	p.SetBlockSize(16)
	if len(iv) != 16 {
		iv = make([]byte, 16)
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		}
	}
	ct := make([][]byte, len(p.blocks))
	for i := range p.blocks {
		if i == 0 {
			ct[i] = aes.ECBEncrypt(xor.XOR1(p.blocks[i], iv), key)
		} else {
			ct[i] = aes.ECBEncrypt(xor.XOR1(p.blocks[i], ct[i-1]), key)
		}
	}
	return AssembleBlocks(ct)
}

func (p *Plaintext) EncryptECB(key []byte) []byte {
	p.ChangeBlockSize(16)
	ct := make([]byte, 0)
	for i := range p.blocks {
		ct = append(ct, aes.ECBEncrypt(p.blocks[i], key)...)
	}
	return ct
}

func (p *Plaintext) ChangeBlockSize(blockSize int) {
	p.blocks = BreakBlocks(p.plaintext, blockSize)
	p.blockSize = blockSize
}
