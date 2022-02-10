package main

func sorting_helper(arr []int) bool {
	for i := 1; i < len(arr) ; i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}

	return true
}
