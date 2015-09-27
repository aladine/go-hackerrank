package main

import (
	"bufio"
	"fmt"
	. "math"
	"os"
	"strconv"
	"strings"
)

const (
	EARTH_RADIUS int = 6371
)

var (
	fileIn  = os.Stdin
	fileOut = os.Stdout
	//fileIn, _    = os.Open(`e1.txt`)
	rw           = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
	bytesWritten = 0
)

type LonLat struct {
	id        int
	latitude  float64
	longitude float64
}

func main() {
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()
	defer report()

	val := readfields()[0]
	N, _ := strconv.Atoi(val)

	m := make([]LonLat, N)
	for i := 0; i < N; i++ {
		line := readfields()
		fmt.Print(len(line))
		a0, _ := strconv.Atoi(line[0])
		a1, _ := strconv.ParseFloat(line[0], 64)
		a2, _ := strconv.ParseFloat(line[0], 64)
		m[i] = new & LonLat{a0, a1, a2}
		// append(m, new LonLat{A[0], A[1], A[2] })

	}
}

func distance_between(point1 LonLat, point2 LonLat) float64 {
	point1_lat_in_radians := degree2radians(point1.latitude)
	point2_lat_in_radians := degree2radians(point2.latitude)
	point1_long_in_radians := degree2radians(point1.longitude)
	point2_long_in_radians := degree2radians(point2.longitude)

	return Acos(Sin(point1_lat_in_radians)*Sin(point2_lat_in_radians)+
		Cos(point1_lat_in_radians)*Cos(point2_lat_in_radians)*
			Cos(point2_long_in_radians-point1_long_in_radians)) * (float64)(EARTH_RADIUS)
}

func radian2degrees(a float64) float64 {
	return a * 57.295779513082
}

func degree2radians(a float64) float64 {
	return a * 0.017453292519
}

// IO ====================================================

func write(s string) {
	n, err := rw.WriteString(s)
	check(err)
	bytesWritten += n
}

func report() {
	println("<wrote", bytesWritten, "bytes>")
}

func readline() (line string) {
	line, err := rw.ReadString('\n')
	line = strings.TrimSpace(line)
	check(err)
	return
}

func readfields() (fields []string) {
	line, err := rw.ReadString('\n')
	check(err)
	fields = strings.Fields(line)
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
