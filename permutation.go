package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	str = "abcd"
	N   = 4
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	getPermutation("", s)

}

func getPermutation(prefix string, s string) {
	if len(s) == 0 {
		fmt.Println(prefix)
	} else {
		for i := 0; i < len(s); i++ {
			getPermutation(prefix+string(s[i]), string(s[0:i])+string(s[i+1:]))
		}
	}

}
