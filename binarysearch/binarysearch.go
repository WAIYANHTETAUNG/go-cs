package binarysearch

import (
	"math"
)

func BinarySearchRecursive(search int, data []int, min int, max int) int {
	mid := int(math.Floor((float64(min) + float64(max)) / 2))

	if max < min {
		return -1
	}

	if data[mid] == search {
		return data[mid]
	}

	if search < data[mid] {
		max = mid - 1
		return BinarySearchRecursive(search, data, min, max)
	}

	min = mid + 1
	return BinarySearchRecursive(search, data, min, max)
}

func BinarySearchLoop(search int, data []int, min int, max int) int {

	for max >= min {
		mid := int(math.Floor((float64(min) + float64(max)) / 2))

		if data[mid] == search {
			return data[mid]
		}

		if search < data[mid] {
			max = mid - 1
		}

		if search > data[mid] {
			min = mid + 1
		}
	}

	return -1
}
