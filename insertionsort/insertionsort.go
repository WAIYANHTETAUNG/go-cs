package insertionsort

func InsertionSort(data []int) []int {
	for index := 1; index < len(data); index++ {
		for j := 0; j < index; j++ {
			if data[index] < data[j] {
				Swap(data, index, j)
			}
		}
	}
	return data
}

func Swap(data []int, firstIndex int, secondIndex int) []int {
	temp := data[firstIndex]
	data[firstIndex] = data[secondIndex]
	data[secondIndex] = temp
	return data
}
