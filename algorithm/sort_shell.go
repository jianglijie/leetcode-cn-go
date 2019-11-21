package algorithm

type Shell struct {
	SortBase
}

func (s *Shell) Sort(arr []int64) {
	length := len(arr)
	gap := 1
	for {
		if gap >= length/3 {
			break
		}
		gap = 3*gap + 1
	}
	for {
		if gap < 1 {
			break
		}
		for i := gap; i < length; i++ {
			for j := i; j >= gap && s.less(arr[j], arr[j-gap]); j -= gap {
				s.swap(arr, j, j-gap)
			}
		}
		gap /= 3
	}
}
