package main

import "fmt"

func checkDup(m1, m2 map[string]int) map[string][]int {
	res := make(map[string][]int)
	for k1 := range m1 {
		v1, ok1 := m1[k1]
		v2, ok2 := m2[k1]

		if ok1 && ok2 {
			if v1 != v2 {
				res[k1] = []int{v1, v2}
			}
		} else if ok1 {
			res[k1] = []int{v1, 0}
		} else if ok2 {
			res[k1] = []int{0, v2}
		} else {
			res[k1] = []int{0, 0}
		}

	}

	for k2, val2 := range m2 {
		_, ok := m1[k2]
		if !ok {
			res[k2] = []int{0, val2}
		}
	}
	//fmt.Println(res)
	return res
}

func main() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 4, "e": 5}
	m2 := map[string]int{"a": 1, "b": 2, "d": 3, "e": 4}
	out := checkDup(m1, m2)
	fmt.Println(out)
}
