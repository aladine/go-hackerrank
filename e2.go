package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var N int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	l := make([]string, N)
	for i := 0; i < N; i++ {
		var s string
		scanner.Scan()
		s = scanner.Text()
		l[i] = s
	}
	sort.Strings(l)

	var desc string
	scanner.Scan()
	desc = strings.ToLower(scanner.Text())

	for i := 0; i < N; i++ {
		if strings.Contains(desc, strings.ToLower(l[i])) {

			fmt.Println(l[i])
		}
	}

}
