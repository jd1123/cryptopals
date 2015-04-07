package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/jd1123/cryptopals/aes"
	"github.com/jd1123/cryptopals/text"
)

var unknownBuff []byte = []byte("this is an unknown string and I will decrypt it using ECB byte at a time attack")

func main() {
	key := aes.RandomKey()
	key = []byte("yellow submarine")
	pt, err := base64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ByteAtATimeECB(pt, key)
}

func ByteAtATimeECB(pt, key []byte) []byte {
	blockLength := DetermineBlockLength(pt, key)
	chkBytes := make([]byte, 0)
	for i := 0; i < blockLength*2; i++ {
		chkBytes = append(chkBytes, byte('A'))
	}
	plainT := text.NewPlaintext(append(chkBytes, pt...))
	fmt.Println(DetectECB(plainT.EncryptECB(key)))
	decrypted := make([]byte, 0)
	for i := 0; i < len(pt); i++ {
		attackBytes := make([]byte, blockLength-1-(i%blockLength))

		for j := 0; j < len(attackBytes); j++ {
			attackBytes[j] = byte('A')
		}

		attackBytes = append(attackBytes)

		//fmt.Println("NewBytes:", string(newBytes))
		dict := make(map[string]byte)
		for j := 0; j < 256; j++ {
			plainT = text.NewPlaintext(append(append(attackBytes, decrypted...), byte(j)))
			dict[string(plainT.EncryptECB(key))] = byte(j)
		}

		//decrypt byte
		plainT = text.NewPlaintext(append(attackBytes, pt...))
		ct := plainT.EncryptECB(key)
		aLen := len(attackBytes)
		dByte := dict[string(ct[0:aLen+1+len(decrypted)])]
		decrypted = append(decrypted, dByte)
		//fmt.Println("i:", i, "dByte:", dByte, "aLen:", aLen, "dLen:", dLen, "len(newBytes):", len(newBytes))
	}
	fmt.Println(string(decrypted))

	return nil
}

func DetermineBlockLength(pt, key []byte) int {
	attackBytes := make([]byte, 0)
	flag := false
	plainT := text.NewPlaintext(pt)
	ct := plainT.EncryptECB(key)
	initialLength := len(ct)
	newBlockLength := initialLength
	ctr := 1

	for i := 0; i < 50; i++ {
		attackBytes = append(attackBytes, byte('A'))
		plainT = text.NewPlaintext(append(attackBytes, pt...))
		ct = plainT.EncryptECB(key)
		//fmt.Println("len(ct):", len(ct), "i:", i, "l2:", l2, "flag:", flag)
		if !flag {
			if len(ct) > initialLength {
				newBlockLength = len(ct)
				flag = true
			}
		} else {
			if len(ct) > newBlockLength {
				break
			} else {
				ctr++
			}
		}
	}
	return ctr
}

func EncryptionOracle(pt []byte) []byte {
	rn := aes.RandInt(2)
	numPre := aes.RandInt(5) + 5
	numPost := aes.RandInt(5) + 5
	pre := make([]byte, numPre)
	post := make([]byte, numPost)

	if _, err := io.ReadFull(rand.Reader, pre); err != nil {
		fmt.Println(err)
	}
	if _, err := io.ReadFull(rand.Reader, post); err != nil {
		fmt.Println(err)
	}
	pt = append(pt, post...)
	pt = append(pre, pt...)
	key := aes.RandomKey()
	plainText := text.NewPlaintext(pt)

	if rn == 0 {
		return plainText.EncryptCBC(key, nil)
	} else {
		return plainText.EncryptECB(key)
	}
}

func DetectECB(ct []byte) bool {
	for i := 0; i < 34; i++ {
		if bytes.Equal(ct[i:i+16], ct[i+16:i+32]) {
			return true
		}
	}
	return false
}
