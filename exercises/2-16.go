package exercises

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jd1123/cryptopals/aes"
	"github.com/jd1123/cryptopals/text"
)

var prependString = "comment1=cooking%20MCs;userdata="
var appendString = ";comment2=%20like%20a%20pound%20of%20bacon"
var randKey = aes.RandomKey()

func Ex2_16() {
	Repl()
}

func Repl() {
	runflag := true
	for runflag {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command>> ")
		inp, _ := reader.ReadString('\n')
		inp = strings.Trim(inp, "\n")

		tokens := strings.Split(inp, " ")

		switch tokens[0] {
		case "exit":
			{
				runflag = false
			}
		case "enc":
			{
				if len(tokens) > 1 {
					for i := range tokens[0:] {
						tokens[i] = strings.Replace(tokens[i], "%20", " ", -1)
					}
					input := strings.Join(tokens[1:], " ")
					fmt.Println("input:", input)
					s, _ := PrepareString(input)
					b := base64.StdEncoding.EncodeToString(s)
					fmt.Println(b)
				}
			}
		case "encb":
			{
				if len(tokens) > 1 {
					for i := range tokens[0:] {
						tokens[i] = strings.Replace(tokens[i], "%20", " ", -1)
					}
					input := strings.Join(tokens[1:], " ")
					s, _ := PrepareString(input)
					smkbits(s)
					b := base64.StdEncoding.EncodeToString(s)
					fmt.Println(b)
				}
			}
		case "dec":
			{
				if len(tokens) > 1 {
					input := tokens[1]
					ct, err := text.NewCiphertextFromBase64([]byte(input))
					if err != nil {
						fmt.Println(err)
					}
					pt, _ := ct.DecryptCBC(key, nil)
					fmt.Println(string(pt))
				}
			}
		case "attk":
			{
				if len(tokens) > 2 {
					input := tokens[1]
					desiredString := tokens[2]
					ct, _ := PrepareString(input)
					res, _ := attack(ct, desiredString, 3, 0)
					bs := base64.StdEncoding.EncodeToString(res)
					fmt.Println(bs)
				}
			}
		default:
			{
				fmt.Println("Unrecognized command")
			}
		}
	}
}

func attack(ct []byte, desiredString string, block, position int) ([]byte, error) {
	numBlocks := len(ct) / 16
	ctCopy := make([]byte, len(ct))
	copy(ctCopy, ct)
	if block == 0 || block > numBlocks {
		return nil, errors.New("block outside acceptable range")
	}
	if len(desiredString) > 16 {
		return nil, errors.New("desired string greater than block size")
	}
	for i := 0; i < len(desiredString); i++ {
		for j := 0; j < 256; j++ {
			pos := (block-1)*16 + position + i
			ctCopy[pos] = ct[pos] ^ byte(j)
			ciphertext := text.NewCiphertext(ctCopy, 16)
			pt, _ := ciphertext.DecryptCBC(key, nil)
			if pt[pos] == byte(desiredString[i]) {
				fmt.Println("Bingo")
				fmt.Println("ct:", ct[pos+16], "ctCopy:", ctCopy[pos+16])
				break
			}
		}
	}
	return ctCopy, nil
}

func smkbits(b []byte) []byte {
	x := byte('=')
	y := x ^ byte(1)
	fmt.Println(string(x), x)
	fmt.Println(string(y), y)
	block := 0
	//blockSize := 16
	index := 0
	//l := len(b)
	b[block*16+index] = b[block*16+index] & byte(1)
	return b
}

func PrepareString(input string) ([]byte, []byte) {
	input = strings.Replace(input, ";", "", -1)
	input = strings.Replace(input, "=", "", -1)
	result := []byte(prependString + input + appendString)

	plaintext := text.NewPlaintext(result)
	ct, iv := plaintext.EncryptCBC(key, nil)
	return ct, iv
}

func CheckString(ct, iv []byte) bool {
	ciphertext := text.NewCiphertext(ct, 16)
	pt, _ := ciphertext.DecryptCBC(key, iv)
	fmt.Println(pt)
	spt := string(pt)
	splitString := strings.Split(spt, ";")
	for i := range splitString {
		tuple := strings.Split(splitString[i], "=")
		if tuple[0] == "admin" && tuple[1] == "true" {
			return true
		}
	}
	return false
}
