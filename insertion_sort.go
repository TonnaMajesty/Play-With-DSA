package main

// 逐个跟前面的元素比较，比前面的元素小就交换
// o(n^2)
func insertion_sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		//for j := i; j - 1 >= 0; j-- {
		//	if arr[j] < arr[j - 1] {
		//		swapSlice(arr, j, j - 1)
		//
		//	} else {
		//		break
		//	}
		//}

		for j := i; j - 1 >= 0 && arr[j] < arr[j - 1]; j-- {
				swapSlice(arr, j, j - 1)
		}
	}
}

// 找插入的位置，把前面的元素向后平移一位，直到找到合适的位置，将当前元素插入
// o(n^2)
func insertion_sort_v2(arr []int) {
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		var j int;
		for j = i; j - 1 >= 0 && arr[j - 1] > temp; j-- {
			arr[j] = arr[j - 1]
		}
		arr[j] = temp
	}
}
