package main

import "fmt"

type MinStack struct {
	arr []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{

	}
}


func (this *MinStack) Push(x int)  {
    this.arr = append(this.arr, x)
}


func (this *MinStack) Pop()  {
    this.arr = this.arr[:len(this.arr)-1]
}


func (this *MinStack) Top() int {
    return this.arr[len(this.arr)-1]
}


func (this *MinStack) GetMin() int {
    min := this.arr[0]
    length := len(this.arr)
    for i:=1;i<length;i++{
        if min > this.arr[i]{
            min = this.arr[i]
        }
    }
    return min
}

func main() {
	obj := Constructor()
	obj.Push(1)
	param_3 := obj.Top()
	fmt.Println(param_3)
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */