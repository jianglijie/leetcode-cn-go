package algorithm

import "fmt"

type SortBase struct {
	
}

func (s *SortBase) less(a int64, b int64) bool {
	return a < b
}

func (s *SortBase) Show(arr []int64) {
	for _, item := range arr {
        fmt.Printf("%d\n", item)
    }
}

func (s *SortBase) swap(arr []int64, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
