package main

import "fmt"

// func reverseStr(s string) {
// 	n := len(s)
// 	res := []string{}

// 	for i := n - 1; i >= 0; i-- {
// 		res = append(res, string(s[i]))
// 	}
// 	fmt.Println(res)

// }

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	s := "abc123 agag"
	//reverseStr(string(s))
	out := Reverse(s)
	fmt.Println(out)

}
