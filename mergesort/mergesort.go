package mergesort

func MergeSort(data []int) []int {
	mid := len(data) / 2
	left := MergeSort(data[0:mid])
	right := MergeSort(data[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	var result []int
	return result
}
