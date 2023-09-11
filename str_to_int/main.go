package main

import (
	"fmt"
	"strings"
)

func strToInt(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	n := len(s)
	fmt.Println("after trim:", s)

	if n == 0 {
		return 0
	}

	var multiplier, start, resp int
	multiplier = 1
	start = 0
	if s[0] == '-' {
		multiplier = -1
		start = 1
	} else if s[0] == '+' {
		multiplier = 1
		start = 0
	}

	fmt.Println("start:", start, "multiplier:", multiplier)
	for i := start; i < n; i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return resp * multiplier
		}

		resp = resp*10 + int(s[i]-'0')
	}
	return resp * multiplier
}

func main() {
	s := "-1221  hi  "
	fmt.Println("Input: ", s)
	output := strToInt(s)
	fmt.Println("Output: ", output)
}
