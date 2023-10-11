package main

import "fmt"

func transformArray(a []int) []int {
	n := len(a)
	b := make([]int, n)

	for i := 0; i < n; i++ {
		var sum int
		fmt.Println("i: ", i, "sum: ", sum)
		// Calculate the sum of a[i-1], a[i], and a[i+1]
		if i > 0 {
			sum += a[i-1]
		}
		sum += a[i]
		if i < n-1 {
			sum += a[i+1]
		}

		// Assign the calculated sum to b[i]
		b[i] = sum
	}

	return b
}

func main() {
	// Example usage
	input := []int{4, 0, 1, -2, 3}
	result := transformArray(input)
	fmt.Println(result)
}
