package ciphertext

import (
	"fmt"

	"github.com/jd1123/cryptopals/freq"
	"github.com/jd1123/cryptopals/xor"
)

// I think this is horribly innefficient
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

func transposeBlocks(byteSlices [][]byte) [][]byte {
	numBlocks := len(byteSlices)
	blockSize := len(byteSlices[0])
	result := make([][]byte, blockSize)
	for i := 0; i < blockSize; i++ {
		result[i] = make([]byte, numBlocks)
	}
	for i := 0; i < numBlocks; i++ {
		for j := 0; j < blockSize; j++ {
			result[j][i] = byteSlices[i][j]
		}
	}
	return result
}

func breakSingleKey(b []byte) byte {
	bestScore := 10000.0
	bestByte := 0
	for i := 0; i < 256; i++ {
		s := freq.ScoreBytes(xor.XORSingleChar(b, byte(i)))
		if s < bestScore {
			bestScore = s
			bestByte = i
		}
	}
	return byte(bestByte)
}

func keyLengthTest(c Ciphertext, keyLength int) float64 {
	if keyLength > len(c.ciphertext) {
		return 0.0
	}
	blocks := breakBlocks(c.ciphertext, keyLength)
	scores := 0.0
	for i := 0; i < 10; i++ {
		scores += 0.1 * float64(xor.HammingDistance(blocks[i], blocks[i+1])) / float64(keyLength)
	}
	return scores
}

func fullKeyLengthTest(c Ciphertext, startLen, endLen int) int {
	scores := make(map[int]float64)
	bestScore := 10000000.0
	bestIx := 0
	for i := startLen; i <= endLen; i++ {
		scores[i] = keyLengthTest(c, i)
		if scores[i] < bestScore {
			bestScore = scores[i]
			bestIx = i
		}
	}
	return bestIx
}
