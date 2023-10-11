package main

import "fmt"

func isValid(in string) bool {
	res := []rune{}
	if len(in)%2 != 0 {
		return false
	}

	for _, r := range in {
		if r == '{' || r == '(' || r == '[' {
			res = append(res, r)
		} else if r == '}' || r == ')' || r == ']' {
			if len(res) == 0 {
				return false
			}

			l := len(res) - 1
			if r != getMatching(res[l]) {
				return false
			} else {
				res = res[:l]
			}
		}
		fmt.Println("res: ", string(res))
	}
	return len(res) == 0

}

func getMatching(r rune) rune {
	switch r {
	case '{':
		return '}'
	case '(':
		return ')'
	case '[':
		return ']'
	}
	return 0
}

func main() {
	s := "(())22()"
	out := isValid(s)

	if out {
		fmt.Println("Is balanced")
	} else {
		fmt.Println("Not balanced")
	}

}
