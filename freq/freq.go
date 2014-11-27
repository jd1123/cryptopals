package freq

import "fmt"

func ScoreList(list [][]byte) int {
	bestScore := 0.0
	keyIndex := 0
	for i := range list {
		score := ScoreBytes(list[i])
		if i == 0 {
			bestScore = score
		} else {
			if bestScore > score {
				keyIndex = i
				bestScore = score
			}
		}
	}
	return keyIndex
}

func ScoreBytes(b []byte) float64 {
	sum := 0.0
	//b = []byte(strings.ToUpper(string(b)))
	for i := range Alphabet {
		p := letterProb(Alphabet[i], b)
		f := FREQ[Alphabet[i]]
		sum += (f - p) * (f - p)
	}
	space_prob := float64(countSpaces(b)) / float64(len(b))
	other_prob := float64(len(b)-countSpaces(b)-countLetters(b)) / float64(len(b))
	sum += (space_prob - 0.15) * (space_prob - 0.15)
	sum += (other_prob - 0.02) * (other_prob - 0.02)
	if sum < 0.0 {
		fmt.Println("SUM ERROR!")
	}
	return sum
}

func IsEnglish(b []byte) bool {
	if ScoreBytes(b) < englishThreshold {
		return true
	} else {
		return false
	}
}

/*
func GetKey(b []byte) byte {
	bestScore := 10000.0
	probableKey := uint8(0)
	for i := 0; i < 256; i++ {
		pt := xor.XORSingleChar(b, byte(i))
		score := ScoreString(string(pt))
		if score < bestScore {
			bestScore = score
			probableKey = byte(i)
		}
	}
	return probableKey
}
*/
