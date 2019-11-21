package algorithm

import (
	"math/rand"
	"time"
)

type Quick struct {
	SortBase
}

func (s *Quick) Shuffle(slice []int64) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

func (s *Quick) Sort(arr []int64) {
	s.Shuffle(arr)
	length := len(arr)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if s.less(arr[j], arr[min]) {
				min = j
			}
		}
		s.swap(arr, i, min)
	}
}

func (s *Quick) sort(arr []int64, start int, end int) {
	if end <= start {
		return
	}
	j := s.partition(arr, start, end)
	s.sort(arr, start, j-1)
	s.sort(arr, j+1, end)
}

func (s *Quick) partition(arr []int64, start int, end int) int {
	i := start
	j := end + 1
	v := arr[start]
	for {
		for {
			i++
			if s.less(arr[i], v) {
				if i == end {
					break
				}
			}
		}
		for {
			j--
			if s.less(v, arr[j]) {
				if j == start {
					break
				}
			}
		}
		if i >= j {
			break
		}
	}
	s.swap(arr, i, j)
	return j
}
