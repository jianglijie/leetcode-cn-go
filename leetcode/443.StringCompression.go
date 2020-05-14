package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	left, right, size := 0, 0, 0
	for ; right <= len(chars); right++ {
		if right == len(chars) || chars[right] != chars[left] {
			chars[size] = chars[left]
			size++
			if right-left > 1 {
				l := strconv.Itoa(right - left)
				for i, _ := range l {
					chars[size] = l[i]
					size++
				}
			}
			left = right
		}
	}
	return size
}

func main() {
	l := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(l))
}
