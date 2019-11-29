package main

import (
	"fmt"
)

func removeKdigits(num string, k int) string {
	length := len(num)
	if length <= k{
		return "0"
	}
	if k == 0{
		return num
	}
	fCount := length - k
	ret, count, start := "", 0, 0
	for {
		if count >= fCount{
			break
		}
		index, min := start, int(num[start])
		for i:=start+1;i<k+count+1;i++{
			if int(num[i]) < min{
				min = int(num[i])
				index = i
			}
		}
		ret += string(min)
		start = index + 1
		count += 1
	}
	if len(ret) == 0{
		return "0"
	}
	if ret[0] == '0'{
		s := 0
		for i, v := range ret{
			if v != '0'{
				s = i
				break
			}
		}
		if s == 0{
			return "0"
		}
		return ret[s:]
	}
	return ret
}

func main()  {
	fmt.Println(removeKdigits("100", 1))
}