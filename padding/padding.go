package padding

import "fmt"

func PKCS7(b []byte, blockSize int) []byte {
	l := len(b)
	blocks := l / blockSize
	padLength := blockSize - l%blockSize

	if padLength > 0 && padLength < blockSize {
		blocks++
	}

	for i := len(b); i < blockSize*blocks; i++ {
		b = append(b, byte(padLength))
	}
	return b
}

func ValidatePadding(b []byte, blockSize int) []byte {
	bLen := len(b)
	padByte := b[bLen]
	if (int(padByte) > blockSize) && (bLen%blockSize != 0) {
		fmt.Println("Invalid padding")
		return b
	}
	return nil
}
