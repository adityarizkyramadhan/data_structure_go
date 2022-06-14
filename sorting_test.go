package data_structure_go

import (
	"testing"
)

func TestSortingAlgorithm_BubbleSort(t *testing.T) {
	array := []int{45, 21, 24, 34, 56, 34, 20, 93, 1, 4}
	sorting := NewSortingAlgorithm(&array)
	sorting.BubbleSortDescending()
	sorting.PrintArray()
	sorting.BubbleSortAscending()
	sorting.PrintArray()
	/*
		=== RUN   TestSortingAlgorithm_BubbleSort
		93, 56, 45, 34, 34, 24, 21, 20, 4, 1,
		1, 4, 20, 21, 24, 34, 34, 45, 56, 93,
		--- PASS: TestBubbleSort (0.00s)
		PASS
		Process finished with the exit code 0
	*/
}

func TestSortingAlgorithm_QuickSort(t *testing.T) {
	array := []int{45, 21, 24, 34, 56, 37, 20, 93, 1, 4}
	sorting := NewSortingAlgorithm(&array)
	sorting.PrintArray()
	sorting.QuickSortAscending()
	sorting.PrintArray()
	sorting.QuickSortDescending()
	sorting.PrintArray()
	/*S
	=== RUN   TestSortingAlgorithm_QuickSort
	1, 4, 20, 21, 24, 34, 34, 45, 56, 93,
	--- PASS: TestSortingAlgorithm_QuickSort (0.00s)
	PASS
	Process finished with the exit code 0
	*/
}
