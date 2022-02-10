package main

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {

	t.Run("test linear search", func(t *testing.T){
		arr := []int{1,2,5,8,10}
		want := 3
		got := linear_search(arr, 8)
		AssertCorrectMessage(t, got, want)
	})

	t.Run("test linear search not find", func(t *testing.T) {
		arr := []int{1,2,5,8,10}
		want := -1
		got := linear_search(arr, 7)
		AssertCorrectMessage(t, got, want)
	})
}
