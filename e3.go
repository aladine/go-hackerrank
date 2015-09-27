package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hotel struct {
	id         int
	price      int
	facilities []string
}

type hotels []*hotel

func (b hotels) Len() int {
	return len(b)
}

func (b hotels) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (s hotels) Less(i, j int) bool {
	if len(s[i].facilities) != len(s[j].facilities) {
		return len(s[i].facilities) > len(s[j].facilities)
	}
	if (s[i].price) != (s[j].price) {
		return (s[i].price) > (s[j].price)
	}

	return (s[i].id) < (s[j].id)

}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	in, _ := os.Open("in2_1.txt")
	scanner := bufio.NewScanner(in)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	l := make(hotels, N)
	for i := 0; i < N; i++ {
		var s string
		scanner.Scan()
		s = scanner.Text()
		str_arr := strings.Split(s, " ")
		id, _ := strconv.Atoi(str_arr[0])
		price, _ := strconv.Atoi(str_arr[1])
		facilities := str_arr[2:]
		l[i] = &hotel{id, price, facilities}
	}

	sort.Sort(l)

	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < M; i++ {
		scanner.Scan()
		s := scanner.Text()
		str_arr := strings.Split(s, " ")
		price, _ := strconv.Atoi(str_arr[0])
		// fmt.Printf(" test case %d: %d \n", i, price)
		for i := 0; i < N; i++ {
			if (price >= l[i].price) && containArray(l[i].facilities, str_arr[1:]) {
				fmt.Printf("%d ", l[i].id)
			}

		}
		fmt.Println("")

	}

}

func containArray(a []string, b []string) bool {

	for i := 0; i < len(b); i++ {
		flag := false
		for j := 0; j < len(a); j++ {
			if a[j] == b[i] {
				flag = true
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
