package xor

import (
	"fmt"
	"testing"
)

func TestHamming(t *testing.T) {
	s1 := []byte("aaa")
	s2 := []byte("bbb")
	fmt.Println(HammingDistance(s1, s2))
	x := byte(255)
	a := (x >> 1) & 0333
	b := (x >> 2) & 0111
	c := x - a - b
	fmt.Printf("%b %b %b %b\n", x, x>>1, x>>2, c)
	d := c + (c >> 3)
	fmt.Println(0307 & d)
	fmt.Printf("%b %b %b\n", 0333, 0111, 0307)
	fmt.Println(string(0307))
	fmt.Println(countBitsByte(x))
	for i := 0; i < 255; i++ {
		if countBitsByte(byte(i)) != countBits(byte(i)) {
			fmt.Println("ERROR")
		}
	}

}
