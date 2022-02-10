package main

//按从小到大排列，循环一次比较两个相邻两个数的大小，然后交换，循环一次即可找出最大的一个数并防在最后一位
func pop_sort(arr []int) {
	for i := 1; i < len(arr); i++ { // 一共比len(arr)-1 轮
		index := 0
		sorted := true
		for j := 0; j < len(arr) - i; j++ {
			if arr[j + 1] < arr[j] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
				index = j // 如果尾部已经排好序了， 我们可以跳过尾部已经排好的那几次循环
				sorted = false
			}
		}

		i = len(arr) - index - 1
		if sorted {
			break
		}
	}
}
