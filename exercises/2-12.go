package exercises

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/jd1123/cryptopals/aes"
	"github.com/jd1123/cryptopals/text"
)

var key = aes.RandomKey()

func Ex2_12() {
	ByteAtATimeECB()
}

func EncryptionOracle2(pt []byte) []byte {
	ap, err := base64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pt = append(pt, ap...)
	plainText := text.NewPlaintext(pt)

	return plainText.EncryptECB(key)
}

func ByteAtATimeECB() []byte {
	blockLength := DetermineBlockLength()
	chkBytes := make([]byte, 0)
	for i := 0; i < blockLength*2; i++ {
		chkBytes = append(chkBytes, byte('A'))
	}
	empty := make([]byte, 0)
	totalLen := len(EncryptionOracle2(empty))
	fmt.Println(aes.DetectECB(EncryptionOracle2(chkBytes)))
	decrypted := make([]byte, 0)
	block := 0
	for i := 0; i < totalLen; i++ {
		if (i != 0) && (i%16 == 0) {
			block++
		}
		attackBytes := make([]byte, blockLength-1-(i%blockLength))

		for j := 0; j < len(attackBytes); j++ {
			attackBytes[j] = byte('A')
		}

		dict := make(map[string]byte)
		for j := 0; j < 256; j++ {
			feedStock := append(append(attackBytes, decrypted...), byte(j))
			result := EncryptionOracle2(feedStock)[block*blockLength : (block+1)*blockLength]
			dict[string(result)] = byte(j)
		}

		//decrypt byte
		ct := EncryptionOracle2(attackBytes)
		dByte := dict[string(ct[block*blockLength:(block+1)*blockLength])]
		decrypted = append(decrypted, dByte)
	}
	fmt.Println(string(decrypted))

	return decrypted
}

func DetermineBlockLength() int {
	attackBytes := make([]byte, 0)
	flag := false
	ct := EncryptionOracle2(attackBytes)
	initialLength := len(ct)
	newBlockLength := initialLength
	ctr := 1

	for i := 0; i < 50; i++ {
		attackBytes = append(attackBytes, byte('A'))
		ct = EncryptionOracle2(attackBytes)
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
