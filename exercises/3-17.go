package exercises

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/jd1123/cryptopals/aes"
	"github.com/jd1123/cryptopals/text"
)

func Ex3_17() {
	randKey := aes.RandomKey()
	randLine, err := RandString()
	if err != nil {
		die(err)
	}
	randLineB, err := base64.StdEncoding.DecodeString(randLine)
	if err != nil {
		die(err)
	}
	plaintext := text.NewPlaintext(randLineB)
	ct, _ := plaintext.EncryptCBC(randKey, nil)
	if err != nil {
		die(err)
	}
	fmt.Println(VerifyPadding(ct))
}

func VerifyPadding(ct []byte) bool {
	ciphertext := text.NewCiphertext(ct, 16)
	pt, err := ciphertext.DecryptCBC(key, nil)
	fmt.Println(pt)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

func die(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func RandString() (string, error) {
	fn := "data/17.txt"
	lines, err := countLines(fn)
	if err != nil {
		return "", nil
	}
	file, err := os.Open(fn)
	if err != nil {
		return "", err
	}
	lineNum := aes.RandIntR(1, lines)
	reader := bufio.NewReader(file)
	for i := 0; i <= lineNum; i++ {
		ln, _, err := reader.ReadLine()
		if err != nil {
			file.Close()
			return "", nil
		}
		if i == lineNum-1 {
			file.Close()
			return string(ln), nil
		}
	}
	return "", nil
}

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return -1, nil
	}
	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	file.Close()
	return lines, nil
}
