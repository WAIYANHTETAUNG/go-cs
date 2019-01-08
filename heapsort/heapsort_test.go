package heapsort_test

import (
	"testing"

	"github.com/pyaesone17/go-cs/heapsort"
)

func TestHash(t *testing.T) {
	heapsort.HeapSort([]int{8, 18, 45, 32, 12, 2, 4, 50})
}
