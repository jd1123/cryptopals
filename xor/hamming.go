package xor

func byteHamming(x, y byte) int {
	return countBitsByte(x ^ y)
}

func countBitsByte(x byte) int {
	c := x - ((x >> 1) & 0333) - ((x >> 2) & 0111)
	return int((c+(c>>3))&0307) % 63
}

func HammingDistance(a, b []byte) int {
	sum := 0
	if len(a) != len(b) {
		return 0
	}
	for i := range a {
		sum += byteHamming(a[i], b[i])
	}
	return sum
}

func HammingUnequal(longerSlice, shorterSlice []byte) int {
	if len(longerSlice) < len(shorterSlice) {
		longerSlice, shorterSlice = shorterSlice, longerSlice
	}
	sliceToCompute := longerSlice[:len(shorterSlice)]
	return HammingDistance(sliceToCompute, shorterSlice)
}

func countBits(b byte) int {
	bits := 0
	for i := 0; i < 8; i++ {
		if b>>uint8(i)&1 == 1 {
			bits++
		}
	}
	return bits
}
