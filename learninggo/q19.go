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

type obj interface{}

// 12
func Map(f func(obj) obj, l []obj) []obj {
	j := make([]obj, len(l))
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
