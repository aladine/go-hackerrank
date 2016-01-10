package main

func main() {
	bubleSort([...]int{2, 3, 5, 4})
}

func bubleSort(list []int) {
	swapped := false

BEGIN:
	swapped = false

	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			list[i-1], list[i] = list[i], list[i-1]
			swapped = true
		}
	}
	if swapped == true {
		goto BEGIN
	}

HERE:
}
