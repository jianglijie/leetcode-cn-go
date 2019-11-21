package main

import (
	"fmt"
)

func removeDuplicates(nums []int64) int {
    length := len(nums)
    for i:=1;i<length;i++{
    	fmt.Println(i-1, i)
        if nums[i-1] == nums[i]{
            if i != length - 1{
                nums = append(nums[:i], nums[i+1:]...)
            }else{
                nums = nums[:i]
            }
            length--
        }
    	fmt.Println(nums)
    }
    return length
}

func main() {
	//sort
	//arr := []int64{3, 2, 5, 78, 89, 4, 55, 43, 18, 123, 24, 75}
	//cls := &algorithm.SortSelection{}
	//cls := algorithm.SortInsertion{}
	//cls := algorithm.Shell{}
	//cls := algorithm.Merge{}
	//cls := algorithm.Quick{}
	//cls := algorithm.Heap{}
	//cls.Sort(arr)
	//cls.SortBU(arr)
	//cls.Show(arr)

	//find
	//findCls := algorithm.Binary{}
	//ret := findCls.Search(arr, 31)
	//if ret != -1{
	//	fmt.Println(ret, arr[ret])
	//}else{
	//	fmt.Println("not found")
	//}

	arr := []int64{1, 1, 2}
	leng := removeDuplicates(arr)
	fmt.Println(leng)

}