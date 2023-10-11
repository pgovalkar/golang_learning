/*

Given an array of strings, remove those that are prefixes of one or more of other strings, and return the remaining strings in the order as they appear in the input array

Example:

input = ["abc", "ab", "bc", "abc hello" ]

output = ["bc", "abc hello"]

"abc" is deleted -> it's a prefix of "abc hello"

*/

package main

import (
	"fmt"
	"strings"
	//"strings"
)

func checkSubStr(in []string) {
	//m := make(map[int]string)

	res := []string{}

	for _, v := range in {
		var isPrefix bool
		for _, str := range in {
			//fmt.Println(str, v)
			if str != v && strings.HasPrefix(str, v) {
				isPrefix = true
				break
			}
		}
		if !isPrefix {
			res = append(res, v)
		}
	}

	fmt.Println(res)

}

func main() {
	input := []string{"abc", "ab", "bc", "abc hello"}
	checkSubStr(input)
}
