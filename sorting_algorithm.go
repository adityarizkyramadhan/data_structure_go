package data_structure_go

import "fmt"

type sortingAlgorithm struct {
	Array *[]int
}

func NewSortingAlgorithm(array *[]int) *sortingAlgorithm {
	return &sortingAlgorithm{
		Array: array,
	}
}

func (s *sortingAlgorithm) BubbleSortDescending() {
	for i := 0; i < len(*s.Array); i++ {
		for j := 0; j < len(*s.Array)-1; j++ {
			if (*s.Array)[j] < (*s.Array)[j+1] {
				temp := (*s.Array)[j]
				(*s.Array)[j] = (*s.Array)[j+1]
				(*s.Array)[j+1] = temp
			}
		}
	}
}

func (s *sortingAlgorithm) BubbleSortAscending() {
	for i := 0; i < len(*s.Array); i++ {
		for j := 0; j < len(*s.Array)-1; j++ {
			if (*s.Array)[j] > (*s.Array)[j+1] {
				temp := (*s.Array)[j]
				(*s.Array)[j] = (*s.Array)[j+1]
				(*s.Array)[j+1] = temp
			}
		}
	}

}

func (s *sortingAlgorithm) PrintArray() {
	for _, data := range *s.Array {
		fmt.Print(data, ", ")
	}
	fmt.Println("")
}

func (s *sortingAlgorithm) QuickSortAscending() {
	s.quickSortRecursiveAscending(len(*s.Array)-1, 0)
}

func (s *sortingAlgorithm) quickSortRecursiveAscending(high, low int) {
	if low < high {
		pivot := s.partitionAscending(high, low)
		s.quickSortRecursiveAscending(pivot-1, low)
		s.quickSortRecursiveAscending(high, pivot+1)
	}
}

func (s *sortingAlgorithm) partitionAscending(high, low int) int {
	pivot := (*s.Array)[high]
	i := low - 1
	for j := low; j < high; j++ {
		if (*s.Array)[j] <= pivot {
			i++
			temp := (*s.Array)[i]
			(*s.Array)[i] = (*s.Array)[j]
			(*s.Array)[j] = temp
		}

	}
	temp := (*s.Array)[i+1]
	(*s.Array)[i+1] = (*s.Array)[high]
	(*s.Array)[high] = temp
	return i + 1
}

func (s *sortingAlgorithm) QuickSortDescending() {
	s.quickSortRecursiveDescending(len(*s.Array)-1, 0)

}

func (s *sortingAlgorithm) quickSortRecursiveDescending(high, low int) {
	if low < high {
		pivot := s.partitionDescending(high, low)
		s.quickSortRecursiveDescending(pivot-1, low)
		s.quickSortRecursiveDescending(high, pivot+1)
	}
}

func (s *sortingAlgorithm) partitionDescending(high, low int) int {
	pivot := (*s.Array)[high]
	i := low - 1
	for j := low; j < high; j++ {
		if (*s.Array)[j] >= pivot {
			i++
			temp := (*s.Array)[i]
			(*s.Array)[i] = (*s.Array)[j]
			(*s.Array)[j] = temp
		}

	}
	temp := (*s.Array)[i+1]
	(*s.Array)[i+1] = (*s.Array)[high]
	(*s.Array)[high] = temp
	return i + 1
}
