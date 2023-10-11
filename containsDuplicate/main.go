package main

import "fmt"

func containsDuplicate(nums []int) bool {
	n := len(nums)

	isdup := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				isdup++
			}
		}
	}
	return isdup != 0
}

func containsDuplicates(nums []int) bool {
	//n := len(nums)
	check := make(map[int]bool)

	for _, num := range nums {
		if check[num] {
			return true
		}
		check[num] = true
	}
	return false
}

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println("Inuput: ", nums)
	out := containsDuplicate(nums)
	out1 := containsDuplicates(nums)
	fmt.Println("Output: ", out)
	fmt.Println("Output1: ", out1)
}
