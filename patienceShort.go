package main

import (
	"container/heap"
	"fmt"
)

type PileHeap [][]int

func (h PileHeap) Len() int      { return len(h) }
func (h PileHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h PileHeap) Less(i, j int) bool {
	return h[i][len(h[i])-1] < h[j][len(h[j])-1]
}

func (h *PileHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *PileHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func bisectPilesRight(piles [][]int, x int) int {
	lo, hi := 0, len(piles)
	for lo < hi {
		// invariant: x maybe between a[lo]...a[hi-1]
		mid := lo + (hi-lo)/2 // don't use (lo+hi)/2 to avoid overflow
		// Note that since (hi-lo)/2 >= 0, lo <= mid < hi
		pile := piles[mid]
		if x < pile[len(pile)-1] { // compare x to top of pile
			hi = mid // x may be between a[lo]...a[mid-1]
		} else {
			lo = mid + 1 // x may be between a[mid+1]...a[hi]
		}
		// The new range is either lo...mid or mid+1...hi and
		// because lo<=mid<hi, the new range is always smaller than lo..hi
	}
	return lo
}

func PatienceSort(a []int) []int {
	piles := make([][]int, 0, 10) // each pile will be a slice.
	for _, x := range a {
		i := bisectPilesRight(piles, x)
		if i < len(piles) {
			piles[i] = append(piles[i], x)
		} else {
			piles = append(piles, []int{x}) // make a new pile
		}
		// fmt.Println(piles)
	}

	h := PileHeap(piles) // Use piles as a heap
	// heap.Init(&h) is not need because piles are already sorted by top card
	n := len(a)
	sorted := make([]int, n)
	for i := 0; i < n; i++ {
		pile := heap.Pop(&h).([]int)
		top := len(pile) - 1
		sorted[i] = pile[top]
		if top > 0 {
			// Put pile minus the top card back in heap if it is not empty
			heap.Push(&h, pile[:top])
		}
	}
	return sorted
}

func main() {
	a := []int{2, 6, 3, 1, 5, 9, 2}
	fmt.Print(PatienceSort(a))
}
