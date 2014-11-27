package xor

import (
	"fmt"
	"testing"
)

func TestHamming(t *testing.T) {
	s1 := []byte("this is a test")
	s2 := []byte("wokka wokka!!!")
	fmt.Println(HammingDistance(s1, s2))
}
