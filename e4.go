package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x float64
	y float64
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	// in, _ := os.Open("in4.txt")
	// scanner := bufio.NewScanner(in)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	radius := make([]int, N)

	scanner.Scan()
	s := scanner.Text()
	str_arr := strings.Split(s, " ")
	for i := 0; i < N; i++ {
		radius[i], _ = strconv.Atoi(str_arr[i])

	}
	sort.Ints(radius)
	MAX := radius[N-1]
	sum_arr := make([]int, MAX+1)
	for i := 0; i < N-1; i++ {
		for j := radius[i]; j < radius[i+1]; j++ {
			sum_arr[j] = i + 1
		}
	}
	sum_arr[MAX] = N
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())
	total := 0
	for i := 0; i < M; i++ {
		scanner.Scan()
		s := scanner.Text()
		num_arr := atois(strings.Split(s, " "))
		p1 := Point{(float64)(num_arr[0]), (float64)(num_arr[1])}
		p2 := Point{(float64)(num_arr[2]), (float64)(num_arr[3])}
		h := getHeight(p1, p2)
		// fmt.Println(h)
		l1 := getNumIntersections(p1, h, sum_arr)
		l2 := getNumIntersections(p2, h, sum_arr)

		total += int(math.Abs((float64)(l1 - l2)))
	}
	fmt.Println(total)

}

func getNumIntersections(p Point, h float64, arr []int) int {
	d := math.Sqrt(p.x*p.x + p.y*p.y)
	c := math.Ceil(h)
	g := math.Floor(d)
	l := float64(len(arr))
	if c >= l {
		c = l - 1
	}
	if g >= l {
		g = l - 1
	}
	// fmt.Printf("%f-%f->%f %f\n", l, g, c, h)
	return arr[int(g)] - arr[int(c)]
	// return (int)(math.Floor(d) - math.Ceil(h) + 1)
}
func getHeight(a, b Point) (r float64) {
	//assume m,0 a.x:a.y b.x:b.y

	if a.y == b.y {
		return math.Abs(a.y)
	}
	if a.x == b.x {
		return math.Abs(a.y)
	}
	v := (a.x - b.x) / (a.y - b.y)
	//(a.x - m) / (a.y) =  v
	// a.x / (a.y - n) = v
	m := a.x - a.y*v
	n := a.y - (a.x)/v

	// fmt.Printf("%f:::%f %f:::%f\n", a.x, a.y, m, n)
	edge := math.Sqrt(m*m + n*n)
	r = math.Abs(m * n / edge)
	return
}

func atois(S []string) (I []int) {
	I = make([]int, len(S))
	for i, a := range S {
		n, err := strconv.Atoi(a)
		check(err)
		I[i] = n
	}
	return
}

func check(err error) {
	if err != nil {
		if err.Error() == "EOF" {
			println("<EOF reached>")
		} else {
			panic(err)
		}
	}
}
