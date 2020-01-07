package main

import (
	"fmt"
)

func asteroidCollision(asteroids []int) []int {
	var s []int
	for _, i := range asteroids {
		if i < 0 {
			for {
				if len(s) == 0 || s[len(s)-1] < 0 {
					s = append(s, i)
					break
				}
				ret := s[len(s)-1] + i
				if ret == 0 {
					s = s[:len(s)-1]
					break
				}else if ret >0{
					break
				}
				s = s[:len(s)-1]
			}
		}else{
			s = append(s, i)
		}
	}
	return s
}

func main() {
	n := []int{-2, -1, 1, 2}
	fmt.Println(asteroidCollision(n))
}
