package exercises

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Ex2_13() {
	fmt.Println(decodeCookie("this=that&mine=stuff&here=there"))
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
