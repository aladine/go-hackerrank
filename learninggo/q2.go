package main

import (
	"fmt"
)

func main() {
	counter := 10
L:
	if counter > 0 {
		counter--
		goto L
	}
	fmt.Println(counter)
}
