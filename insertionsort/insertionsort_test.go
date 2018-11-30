package insertionsort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	// sorted array
	data := []int{1, 4, 8, 10, 14, 18}
	result2 := InsertionSort(data)
	fmt.Println(result2)
}
