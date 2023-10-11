package main

import "fmt"

func removeElement(nums []int, val int) int {
	res := []int{}

	for _, v := range nums {
		if v != val {
			res = append(res, v)
		}
	}
	fmt.Println(res)
	return len(res)
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	result := removeElement(nums, val)
	fmt.Println("result: ", result)

}
