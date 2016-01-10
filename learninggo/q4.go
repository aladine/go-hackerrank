package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	counter := 10
	line := []string{}
L:
	if counter > 0 {
		line = append(line, "a")
		fmt.Println(line)
		counter--
		goto L
	}
	fmt.Println(counter)

	// utf8.RuneCount
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Println(utf8.RuneCount([]byte(str)))
	// fmt.Println(strings.Replace(str, "a", "b", 1))
	s := "￿￿￿ ￿￿￿￿ ￿￿ ￿￿￿￿￿"
	r := []rune(s)
	copy(r[4:4+3], []rune("abc"))

	// reverse
	arr := []byte(str)
	arr2 := []rune(str)

	result := []byte{}
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	fmt.Println(string(result))

	result2 := []rune{}
	for i := len(arr2) - 1; i >= 0; i-- {
		result2 = append(result2, arr2[i])
	}
	fmt.Println(string(result2))

	// another solution
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i] //← Parallel assignment
	}
	fmt.Printf("%s\n", string(a))
}
