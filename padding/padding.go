package padding

import "errors"

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

func ValidatePKCS7(b []byte, blockSize int) ([]byte, error) {
	bLen := len(b)
	padByte := b[bLen-1]
	for i := 1; i < int(padByte); i++ {
		if b[bLen-i] != padByte {
			return nil, errors.New("Invalid padding")
		}
	}
	return b[:bLen-int(padByte)], nil
}
