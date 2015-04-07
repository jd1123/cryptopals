package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jd1123/cryptopals/exercises"
)

type Exercise func()

var ex = map[string]Exercise{
	"1-1":  exercises.Ex1_1,
	"1-2":  exercises.Ex1_2,
	"1-3":  exercises.Ex1_3,
	"1-5":  exercises.Ex1_5,
	"1-6":  exercises.Ex1_6,
	"1-7":  exercises.Ex1_7,
	"1-8":  exercises.Ex1_8,
	"2-10": exercises.Ex2_10,
	"2-11": exercises.Ex2_11,
	"2-12": exercises.Ex2_12,
	"2-13": exercises.Ex2_13,
}

func main() {
	spl := make([]string, 0)
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("usage: cryptopals <exercise>")
		os.Exit(1)
	} else {
		exercise := os.Args[1]
		spl = strings.Split(exercise, "-")
		if len(spl) != 2 {
			Usage()
			fmt.Println("invalid exercise number")
			os.Exit(1)
		} else {
			if exFound(exercise) {
				ex[exercise]()
			} else {
				fmt.Println("Exercise", exercise, "not found")
			}
			os.Exit(0)
		}
	}
	fmt.Println(spl)
}

func exFound(tst string) bool {
	for k, _ := range ex {
		if tst == k {
			return true
		}
	}
	return false
}

func Usage() {
	fmt.Println("usage: cryptopals <exercise>")
}
