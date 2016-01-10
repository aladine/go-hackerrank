package main

import (
	"fmt"
)

func main() {
	fmt.Println(2)
	sort(7, 2)
}

// func order(a, b int) (int, int) {}
func sort(a, b int) (c, d int) {
	if a < b {
		c, d = a, b
	} else {

		c, d = b, a
	}
	return
}
