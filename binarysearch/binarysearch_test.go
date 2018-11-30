package binarysearch

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestBinarySearch(t *testing.T) {
	// sorted array
	data := []int{1, 4, 8, 10, 14, 18}
	result := BinarySearchRecursive(18, data, 0, len(data))
	assert.Equal(t, result, 18)
}

func TestBinarySearchLoop(t *testing.T) {
	// sorted array
	data := []int{1, 4, 8, 10, 14, 18}
	result2 := BinarySearchLoop(18, data, 0, len(data))
	assert.Equal(t, result2, 18)
}
