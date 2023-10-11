package main

import (
	"fmt"
)

func sequencialCount(s []int) bool {
	//s = strings.ReplaceAll(s, " ", "")
	n := len(s)

	isInc, isDec := 0, 0

	if n == 0 {
		return false
	}

	for i := 0; i < n-1; i++ {
		if s[i] < s[i+1] {
			isInc++
		}

		if s[i] > s[i+1] {
			isDec++
		}

		if isDec != 0 && isInc != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := []int{6, 5, 4, 4}
	fmt.Println("Input: ", s)
	out := sequencialCount(s)
	fmt.Println("Is Monotonic: ", out)
}
