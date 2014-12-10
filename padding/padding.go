package padding

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
