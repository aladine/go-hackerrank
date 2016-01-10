package main

import (
	"fmt"
)

func main() {
	fmt.Println(2)
	p := plusTwo()
	fmt.Println(p(3))
}

func printLine(args ...int) {
	for _, x := range args {
		fmt.Println(x)
	}
}

// 12
func Map(f func(int) int, l []int) []int {
	j := make([]int, len(l))
	for k, v := range l {
		j[k] = f(v)
	}
	return j
}

//15

func plusTwo() func(int) int {
	return func(x int) int {
		return x + 2
	}
}
