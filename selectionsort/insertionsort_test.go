package insertionsort

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	// sorted array
	data := []int{1, 10, 2, 28, 3}
	FindMinOnSplit(data, 1)
}

func TestSelectionSort(t *testing.T) {
	// sorted array
	data := []int{1, 10, 2, 28, 3}
	res := SelectionSort(data)
	fmt.Println(res)
}
