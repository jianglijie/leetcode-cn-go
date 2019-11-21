package algorithm

type Heap struct {
	SortBase
}

func (s *Heap) Sort(arr []int64) {
	length := len(arr)
	// max heap
	for i := length; i >= 0; i-- {
		s.heapify(arr, length, i)
	}
	for i := length - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		s.heapify(arr, i, 0)
	}
}

func (s *Heap) heapify(arr []int64, n int, i int) {
	largest := i
	l := 2*i + 1 //left = 2*i + 1
	r := 2*i + 2 //right = 2*i + 2
	if l < n && arr[i] < arr[l] {
		largest = l
	}
	if r < n && arr[largest] < arr[r] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		s.heapify(arr, n, largest)
	}
}
