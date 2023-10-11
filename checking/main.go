package main

import (
	"fmt"
)

func checking(in string) bool {
	res := []rune{}

	for _, s := range in {
		if s == '(' || s == '{' || s == '[' {
			res = append(res, s)
		} else if s == ')' || s == '}' || s == ']' {
			if len(res) == 0 {
				return false
			}
			l := len(res) - 1
			if s != getMatchingBracket(res[l]) {
				return false
			}

			res = res[:l]
		}
	}
	return len(res) == 0

}

func getMatchingBracket(in rune) rune {
	switch in {
	case '{':
		return '}'
	case '[':
		return ']'
	case '(':
		return ')'
	}
	return 0
}
func main() {
	//m1 := map[string]int{"a": 1, "b": 2, "c": 3, "e": 5}
	//m2 := map[string]int{"a": 1, "b": 2, "d": 3, "e": 6}

	in := "(())"
	out := checking(in)
	fmt.Println(out)
}

// palindrome 1221
// fibonacci: 0 1 1 2 3 5
// string to int "1234" --> 1234, "aaa -123" --> -123
// twosum  s := []int{1, 2, 3, 4, 5, 6} , target := 5   ==> [0 3]
// Monotonic Input:  [6 5 4 4] Is Monotonic:  true   ; Input:  [1 3 2] Is Monotonic:  false

/*
m1:{a:1, b:2,c:3,e:5}  m2:{a:1, b:2,d:2,e:6} out= {c:[3,0], d:[0,2], e:[5,6]}
*/

/*  in={"abc", "ab", "bc", "abc hello" } out: [bc abc hello]
 */

// in="(([])){}" out=true - string is balanced
