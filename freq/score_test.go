package freq

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"
)

/*
func TestScore(t *testing.T) {
	s := "this is an actual sentance"
	n := "jfklsjalfsjkdsjfklsdjl=k slajflkjsdkl =f3ou8#@@R EW"
	a := "abc"
	fmt.Println(ScoreString(s))
	fmt.Println(ScoreString(n))
	fmt.Println(ScoreString(a))
}

func TestScoreForNumber(t *testing.T) {
	numTests := 100
	score := make([]float64, numTests)
	actualEnglish := "This is an actual english string. What's the score?"
	actualEnglish2 := "This is another actual english string. I think this is a decent test, but I'm not sure"
	actualEnglish3 := "TEHCNICALLYTHISISENGLISHXXXBUTTHEREARENOSPACESORPUNCTUATION"
	for i := 0; i < numTests; i++ {
		s := make([]byte, 47)
		if _, err := io.ReadFull(rand.Reader, s); err != nil {
			fmt.Println("Random number gen error")
		}
		score[i] = ScoreString(string(s))
		if score[i] <= englishThreshold {
			fmt.Println(string(s))
		}
	}
	fmt.Println(ScoreString(string(actualEnglish)))
	fmt.Println(ScoreString(string(actualEnglish2)))
	fmt.Println(ScoreString(string(actualEnglish3)))
	bs := []byte("This is a test string")
	fmt.Println(bs[:3])
	fmt.Println(bs[3:6])
	fmt.Println(bs)
}

*/

func TestCount(t *testing.T) {
	b := []byte("abcd Ef. ghi.")
	if countLetters(b) != 9 {
		t.Errorf("The number of letters should be 9, it is not")
	}
	if countSpaces(b) != 2 {
		t.Errorf("The number of spaces should be 2, it is not")
	}
}

func TestScore(t *testing.T) {
	b := []byte("ABC")
	b1 := []byte("abc")
	randBytes := make([]byte, 102)
	if _, err := io.ReadFull(rand.Reader, randBytes); err != nil {
		t.Errorf("Io Error")
	}
	fmt.Println("Score Random:", ScoreBytes(randBytes))
	fmt.Println("Score b:", ScoreBytes(b))
	fmt.Println("Score b1:", ScoreBytes(b1))
}
