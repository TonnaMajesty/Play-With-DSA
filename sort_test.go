package main

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	t.Run("selection_sort", func(t *testing.T) {
		arr := []int{98, 56, 67, 1, 4, 55, 43, 66}
		selection_sort(arr)
		res := sorting_helper(arr)
		if !res {
			t.Errorf("selection_sort fail")
		}
	})

	t.Run("selection_sort_v2", func(t *testing.T) {
		arr := []int{98, 56, 67, 1, 4, 55, 43, 66}
		selection_sort_v2(arr)
		res := sorting_helper(arr)
		if !res {
			t.Errorf("selection_sort fail")
		}
	})

	t.Run("insertion_sort", func(t *testing.T) {
		arr := []int{98, 56, 67, 1, 4, 55, 43, 66}
		insertion_sort(arr)
		res := sorting_helper(arr)
		if !res {
			t.Errorf("insertion_sort fail")
		}
	})

	t.Run("insertion_sort_v2", func(t *testing.T) {
		arr := []int{98, 56, 67, 1, 4, 55, 43, 66}
		insertion_sort_v2(arr)
		res := sorting_helper(arr)
		if !res {
			t.Errorf("insertion_sort_v2 fail")
		}
	})

	t.Run("pop_sort", func(t *testing.T) {
		arr := []int{98, 56, 67, 1, 4, 55, 43, 66}
		pop_sort(arr)
		res := sorting_helper(arr)
		if !res {
			t.Errorf("pop_sort fail")
		}
	})
}

func BenchmarkSelectionSortV2(b *testing.B) {
	b.StopTimer()
	n := 10000
	data := generateRandomArray(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		selection_sort_v2(data)
	}
}

