package main

func selection_sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				swapSlice(arr, j, i)
			}
		}
	}
}

func selection_sort_v2(arr []int) {
	for i := 0; i < len(arr) - 1; i++ { // 最后一个不需要再找了
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swapSlice(arr, i, minIndex)
	}

}
