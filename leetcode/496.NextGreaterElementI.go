package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    m := make(map[int]int)
    l := len(nums2)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums2[j] > nums2[i] {
				m[nums2[i]] = nums2[j]
				break
			}
		}
	}
	ret := []int{}
	for _, v := range nums1{
		if _, ok := m[v];!ok{
			ret = append(ret, -1)
		}else{
			ret = append(ret, m[v])
		}
	}
	return ret
}
