package exercises

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jd1123/cryptopals/aes"
	"github.com/jd1123/cryptopals/text"
)

func Ex2_13() {
	email := "hello@admin               "
	runExercise(email)
}

func runExercise(email string) {
	key := aes.RandomKey()
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
	result := "{\"email\":\"" + sanitized + "\",\"uid\":\"10\",\"role\":\"user\"}"
	return []byte(result)
}
