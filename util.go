package main

import (
	"math/rand"
	"testing"
)

func AssertCorrectMessage(t *testing.T, got, want interface{}) {
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func swapSlice(arr []int, i, j int) {
	temp := arr[j]
	arr[j] = arr[i]
	arr[i] = temp
}

func generateRandomArray(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Int()
	}
	return data
}