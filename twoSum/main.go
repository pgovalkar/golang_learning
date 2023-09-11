package main

import "fmt"

func twoSum(num []int, target int) []int {
	n := len(num)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if num[i]+num[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func main() {
	num := []int{2, 3, 4, 5, 6}
	taget := 10
	out := twoSum(num, taget)
	fmt.Println(out)
}
