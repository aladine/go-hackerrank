package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// palindrome returns the minimum number of operations carried out
// to convert a word into a palindrome.
// A word is a set of lower-case ASCII letters.
//scanner := bufio.NewScanner(os.Stdin)
// scanner.Split(bufio.ScanLines)
func palindrome(word string) (ops int) {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		n := int(word[i]) - int(word[j])
		if n < 0 {
			n = -n
		}
		ops += n
	}
	return ops
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		t, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
			t = 0
		}
		for ; t > 0 && scanner.Scan(); t-- {
			fmt.Println(palindrome(scanner.Text()))
		}
		if t > 0 {
			fmt.Fprintln(os.Stderr, "reading input:", io.ErrUnexpectedEOF)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input", err)
	}

}
