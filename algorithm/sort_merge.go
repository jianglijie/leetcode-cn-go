package algorithm

type Merge struct {
	SortBase
	aux []int64
}

// top->bottom
func (s *Merge) Sort(arr []int64) {
	length := len(arr)

	s.sort(arr, 0, length-1)
}

// bottom->top
func (s *Merge) SortBU(arr []int64) {
	length := len(arr)
	s.aux = make([]int64, length)
	for size := 1; size < length; size += size {
		for start := 0; start < length-size; start += size + size {
			s.merge(arr, start, start+size-1, s.min(start+size+size-1, length-1))
		}
	}
}

func (s *Merge) min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (s *Merge) merge(arr []int64, start int, mid int, end int) {
	i := start
	j := mid + 1
	s.aux = make([]int64, len(arr))
	for k := start; k <= end; k++ {
		s.aux[k] = arr[k]
	}
	for k := start; k <= end; k++ {
		if i > mid {
			arr[k] = s.aux[j]
			j++
		} else if j > end {
			arr[k] = s.aux[i]
			i++
		} else if s.less(s.aux[j], s.aux[i]) {
			arr[k] = s.aux[j]
			j++
		} else {
			arr[k] = s.aux[i]
			i++
		}
	}
}

func (s *Merge) sort(arr []int64, start int, end int) {
	if end <= start {
		return
	}
	mid := start + (end-start)/2
	s.sort(arr, start, mid)
	s.sort(arr, mid+1, end)
	s.merge(arr, start, mid, end)
}
