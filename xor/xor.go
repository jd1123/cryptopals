package xor

import "errors"

func XOR(b1, b2 []byte) ([]byte, error) {
	if len(b1) != len(b2) {
		return nil, errors.New("byte slices are not same length")
	}
	b3 := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		b3[i] = b1[i] ^ b2[i]
	}
	return b3, nil
}

func XORSingleChar(byteSlice []byte, b byte) []byte {
	if len(byteSlice) <= 0 {
		return nil
	}
	result := make([]byte, len(byteSlice))
	for i := range byteSlice {
		result[i] = byteSlice[i] ^ b
	}
	return result
}

func XORRepeatingKey(byteSlice []byte, key []byte) []byte {
	keyLength := len(key)
	result := make([]byte, len(byteSlice))
	for i := range byteSlice {
		result[i] = byteSlice[i] ^ key[i%keyLength]
	}
	return result
}
