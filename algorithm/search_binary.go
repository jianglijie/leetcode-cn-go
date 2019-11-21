package algorithm

type Binary struct {
}

func (s *Binary) Search(arr []int64, key int64) int {
	start := 0
	end := len(arr) - 1
	for {
		if start > end {
			break
		}
		mid := start + (end-start)/2
		if key < arr[mid] {
			end = mid - 1
		} else if key > arr[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
