package algorithm

type SortSelection struct {
	SortBase
}

func (s *SortSelection) Sort(arr []int64) {
	length := len(arr)
	if length <= 1 {
		return
	}
	for i:=0; i<length;i++ {
		min := i
		for j:=i+1;j<length;j++{
			if s.less(arr[j], arr[min]){
				min = j
			}
		}
		s.swap(arr, i, min)
	}
}
