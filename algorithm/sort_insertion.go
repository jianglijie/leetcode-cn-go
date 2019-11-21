package algorithm

type SortInsertion struct {
	SortBase
}

func (s *SortInsertion) Sort(arr []int64) {
	length := len(arr)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		for j := i; j > 0 && s.less(arr[j], arr[j-1]); j-- {
			s.swap(arr, j, j - 1)
		}
	}
}
