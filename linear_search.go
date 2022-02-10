package main

import "fmt"

func linear_search (arr []int, value int) int {
	for key, val := range arr {
		if val == value {
			return key
		}
	}
	return -1
}

func main () {
	arr1 := []int{1,8,19,20,14}
	key := linear_search(arr1, 20)
	fmt.Println(key)

}