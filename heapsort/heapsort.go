package heapsort

import (
	"fmt"
	"math"
)

func Heapify(arr []int, length int, root int) {
	largest := root
	leftIndex := 2*root + 1
	rightIndex := 2*root + 2

	if leftIndex < length && arr[leftIndex] > arr[largest] {
		largest = leftIndex
	}

	if rightIndex < length && arr[rightIndex] > arr[largest] {
		largest = rightIndex
	}

	if root != largest {
		temp := arr[root]
		arr[root] = arr[largest]
		arr[largest] = temp

		Heapify(arr, length, largest)
	}

}

func HeapSort(arr []int) {
	n := len(arr)

	for index := math.Floor(float64(n/2) - 1); index >= 0; index-- {
		Heapify(arr, n, int(index))
	}

	for j := n - 1; j >= 0; j-- {
		temp := arr[0]
		arr[0] = arr[j]
		arr[j] = temp

		Heapify(arr, j, 0)
	}

	fmt.Println(arr)
}
