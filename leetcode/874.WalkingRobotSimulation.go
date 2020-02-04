package main

import "fmt"

func robotSim(commands []int, obstacles [][]int) int {
    x, y, direct, max := 0, 0, 0, 0
	set := make(map[string]bool)
	for _, i := range obstacles{
		set[fmt.Sprintf("%s-%s",i[0],i[1])] = true
	}
	for _, c := range commands{
		if c == -1{
			direct = (direct+1)%4
		}else if c == -2{
			direct = (direct-1+4)%4
		}else{
			if direct == 0{
				for{
					if _, ok := set[fmt.Sprintf("%s-%s", x,y+1)]; ok {
						break
					}
					if c == 0{
						break
					}
					y += 1
					c -= 1
				}
			}else if direct == 2{
				for{
					if _, ok := set[fmt.Sprintf("%s-%s", x,y-1)]; ok {
						break
					}
					if c == 0{
						break
					}
					y -= 1
					c -= 1
				}
			}else if direct == 1{
				for{
					if _, ok := set[fmt.Sprintf("%s-%s", x+1,y)]; ok {
						break
					}
					if c == 0{
						break
					}
					x += 1
					c -= 1
				}
			}else if direct == 3{
				for{
					if _, ok := set[fmt.Sprintf("%s-%s", x-1,y)]; ok {
						break
					}
					if c == 0{
						break
					}
					x -= 1
					c -= 1
				}
			}
			tmp := x * x + y * y
			if tmp > max{
				max = tmp
			}
		}
	}
	return max
}

func main() {
	
}
