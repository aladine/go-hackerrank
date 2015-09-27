package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	// in, _ := os.Open("in4.txt")
	// scanner := bufio.NewScanner(in)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())
	// var T int
	// fmt.Scan(&T)

	for i := 0; i < T; i++ {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		s := scanner.Text()
		str_arr := strings.Split(s, " ")

		int_arr := make([]int, N)
		for j := 0; j < N; j++ {
			d, _ := strconv.Atoi(str_arr[j])
			int_arr[j] = d
		}
		solve(N, int_arr)
	}

}

func solve(N int, arr []int) {
	m := arr[0]
	for i := 0; i < N; i++ {
		if m > arr[i] {
			m = arr[i]
		}
		// fmt.Println(arr[i])
	}
	var sum1, sum2, sum3 int

	for i := 0; i < N; i++ {
		diff := arr[i] - m
		// fmt.Println("---")
		// fmt.Println(diff)
		sum1 += int(diff/5) + int((diff%5)/2) + int((diff%5)%2)
		sum2 += int((diff+1)/5) + int(((diff+1)%5)/2) + int(((diff+1)%5)%2)
		sum3 += int((diff+2)/5) + int(((diff+2)%5)/2) + int(((diff+2)%5)%2)
	}
	if sum1 <= sum2 && sum1 <= sum3 {
		fmt.Println(sum1)
	} else if sum2 <= sum3 {
		fmt.Println(sum2)
	} else {
		fmt.Println(sum3)
	}

}

// func atois(S []string) (I []int) {
// 	I = make([]int, len(S))
// 	for i, a := range S {
// 		n, err := strconv.Atoi(a)
// 		check(err)
// 		I[i] = n
// 	}
// 	return
// }

// func check(err error) {
// 	if err != nil {
// 		if err.Error() == "EOF" {
// 			println("<EOF reached>")
// 		} else {
// 			panic(err)
// 		}
// 	}
// }
