package insertionsort

func FindMinOnSplit(data []int, startIndex int) int {
	startValue := data[startIndex]
	foundIndex := startIndex

	for index := startIndex + 1; index < len(data); index++ {
		if data[index] < startValue {
			foundIndex = index
			startValue = data[index]
		}
	}

	return foundIndex
}

func Swap(data []int, firstIndex, secondIndex int) {
	temp := data[firstIndex]
	data[firstIndex] = data[secondIndex]
	data[secondIndex] = temp
}

func SelectionSort(data []int) []int {
	for index := 0; index < len(data); index++ {
		minIndex := FindMinOnSplit(data, index)
		Swap(data, index, minIndex)
	}

	return data
}
