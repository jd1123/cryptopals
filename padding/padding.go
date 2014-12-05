package padding

func PKCS7(b []byte, blockSize int) []byte {
	l := len(b)
	blocks := l / blockSize
	padLength := blockSize - l%blockSize
	if padLength > 0 && padLength < blockSize {
		blocks++
	}
	newPT := make([]byte, blockSize*blocks)
	for i := 0; i < len(b); i++ {
		newPT[i] = b[i]
	}
	for i := len(b); i < len(newPT); i++ {
		newPT[i] = byte(padLength)
	}
	return newPT
}
