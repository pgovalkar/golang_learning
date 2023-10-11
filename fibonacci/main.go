package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)

}

func printFibonacciSeries(num int) {
	a := 0
	b := 1
	c := b
	fmt.Printf("Series is: %d %d", a, b)
	for {
		c = b     //
		b = a + b //
		if b >= num {
			fmt.Println()
			break
		}
		a = c ///1
		fmt.Printf(" %d", b)
	}
}

// fiboncci n=5 0,1,1,2,3
func main() {
	n := 6
	fmt.Println(fibonacci(n))
	printFibonacciSeries(n)

}
