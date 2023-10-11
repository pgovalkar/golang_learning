package main

import (
	"fmt"
	"strconv"
)

func monotoneIncreasingDigits(n int) int {
	//tmp := n
	a := strconv.Itoa(n)
	alen := len(a)

	isInc, isDec, eq := 0, 0, 0
	for i := 0; i < alen-1; i++ {
		if a[i] > a[i+1] {
			isDec++
		}
		if a[i] < a[i+1] {
			isInc++
		}
		// if a[i] == a[i+1] {
		// 	eq++
		// }

	}
	fmt.Println("isInc: ", isInc, "IsDec: ", isDec, "eq: ", eq)
	// if isDec != 0 && eq != 0 {
	// 	return n - (isDec + eq) - 1
	// } else if isInc != 0 && eq != 0 {
	// 	return n + isInc + eq
	// }

	if isDec != 0 {
		return n - isDec
	} else if isInc != 0 {
		return n
	}

	return n
}

func main() {
	n := 1234
	out := monotoneIncreasingDigits(n)
	fmt.Println("Output: ", out)
}
