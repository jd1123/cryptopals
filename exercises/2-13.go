package exercises

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jd1123/cryptopals/aes"
	"github.com/jd1123/cryptopals/text"
)

var randKey2_13 = aes.RandomKey()

func Ex2_13() {
	Repl2_13()
	//runExercise(email)
}

func Repl2_13() {
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
		case "profile":
			{
				if len(tokens) > 1 {
					for i := range tokens[0:] {
						tokens[i] = strings.Replace(tokens[i], "%20", " ", -1)
					}
					input := strings.Join(tokens[1:], " ")
					fmt.Println("input:", input)
					cookie := encryptCookie(profileFor(input), key)
					str := base64.StdEncoding.EncodeToString(cookie)
					fmt.Println(str)
				}
			}
		case "dec":
			{
				if len(tokens) > 1 {
					ct, err := base64.StdEncoding.DecodeString(tokens[1])
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					prof := string(decryptCookie(ct, key))
					fmt.Println(prof)
					prof2 := decodeCookie(prof)
					fmt.Println(string(prof2))
				}
			}
		case "attack":
			{
				attackString := "0123456789admin           "
				cookie := encryptCookie(profileFor(attackString), key)
				adminBlock := cookie[16:32]
				fmt.Println("decoded block!", string(aes.ECBDecrypt(adminBlock, key)))
				attackString = "bill1@aol.com"
				cookie = encryptCookie(profileFor(attackString), key)
				position := len(cookie) - 16
				newCookie := cookie[0:position]
				newCookie = append(newCookie, adminBlock...)
				fmt.Println(base64.StdEncoding.EncodeToString(newCookie))

			}
		default:
			{
				fmt.Println("Unrecognized command")
			}
		}
	}
}

func runExercise(email string) {
	cookie := profileFor(email)
	ct := encryptCookie(cookie, key)
	pt := decryptCookie(ct, key)
	fmt.Println("pt:", string(pt))
	fmt.Println("len ct:", len(ct))
	fmt.Println("ct:", ct)
	fmt.Println(string(pt))

}

func encryptCookie(cookie, key []byte) []byte {
	pt := text.NewPlaintext(cookie)
	ct := pt.EncryptECB(key)
	return ct
}

func decryptCookie(cookie, key []byte) []byte {
	ct := text.NewCiphertext(cookie, 16)
	pt := ct.DecryptECB(key)
	return pt
}

func decodeCookie(cookie string) []byte {
	obj := make(map[string]string)
	codes := strings.Split(cookie, "&")
	for _, c := range codes {
		fields := strings.Split(c, "=")
		obj[fields[0]] = fields[1]
	}
	j, err := json.Marshal(obj)
	if err != nil {
		return nil
	} else {
		return j
	}
}

func profileFor(email string) []byte {
	// Sanitize input
	sanitized := strings.Join(strings.Split(email, "&"), "")
	sanitized = strings.Join(strings.Split(sanitized, "="), "")
	//result := "{\"email\":\"" + sanitized + "\",\"uid\":\"10\",\"role\":\"user\"}"
	result := "email=" + sanitized + "&uid=10&role=user"
	return []byte(result)
}
