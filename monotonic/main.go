package main

import "fmt"

func isMonotonic(num []int) bool {
	n := len(num)

	var isInc, isDec int = 0, 0
	for i := 0; i < n-1; i++ {
		if num[i] < num[i+1] {
			isInc++
		}
		if num[i] > num[i+1] {
			isDec++
		}
		if isInc != 0 && isDec != 0 {
			return false
		}
	}
	return true
}

func main() {
	num := []int{6, 5, 4, 4}
	fmt.Println("Input: ", num)
	output := isMonotonic(num)
	fmt.Println("Output: ", output)
}
