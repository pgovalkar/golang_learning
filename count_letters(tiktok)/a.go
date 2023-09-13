package main

import "fmt"

func countLetters(in string) {
	var curChar rune
	curCount := 0

	for _, char := range in {
		if char == curChar {
			curCount++
		} else {
			if curCount > 0 {
				fmt.Printf("%c : %d\n", curChar, curCount+1)
			}
			curChar = char
			curCount = 0
		}
	}

	if curCount > 0 {
		fmt.Printf("%c : %d\n", curChar, curCount+1)
	}
}

func main() {
	in := "ttttiikkk ttooookkk"
	countLetters(in)
}

