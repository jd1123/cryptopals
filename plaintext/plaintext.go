package plaintext

import (
	"fmt"

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

func assembleBlocks(b [][]byte) []byte {
	numBlocks := len(b)
	blockSize := len(b[0])
	assembled := make([]byte, numBlocks*blockSize)
	for i := range b {
		for j := range b[i] {
			assembled[i*blockSize+j] = b[i][j]
		}
	}
	return assembled
}

func breakBlocks(byteSlice []byte, blockSize int) [][]byte {
	excess := false
	add := 0
	mod := len(byteSlice) % blockSize
	if mod > 0 {
		excess = true
		add = 1
	}

	blockCount := len(byteSlice) / blockSize
	result := make([][]byte, blockCount+add)

	for i := 0; i < blockCount; i++ {
		result[i] = byteSlice[i*blockSize : (i+1)*blockSize]
		if excess && i == blockCount-1 {
			fmt.Println("Excess!!")
			result[i+1] = make([]byte, blockSize)
			//result[i+1][:mod] = byteSlice[blockSize*(i+1):]
			for j := 0; j < blockSize; j++ {
				if j < mod {
					result[i+1][j] = byteSlice[blockSize*(i+1)+j]
				} else {
					result[i+1][j] = byte(' ')
				}
			}
		}
	}

	return result
}
