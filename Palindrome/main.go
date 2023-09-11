package main

import "fmt"

func Palindrome(a int) bool {
	tmp := a
	res := 0

	for tmp > 0 {
		lastval := tmp % 10
		res = res*10 + lastval
		tmp = tmp / 10
	}

	return res == a
}

func main() {
	a := 10011001
	fmt.Println("Input:", a)
	out := Palindrome(a)
	fmt.Println(out)
}
