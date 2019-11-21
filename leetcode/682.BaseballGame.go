package main

import "strconv"

func calPoints(ops []string) int {
    var ret []int
	for _, i := range ops {
        l := len(ret)
		switch i {
		case "C":
			if l > 0 {
				ret = ret[:l-1]
			}
		case "D":
			if l > 0 {
				ret = append(ret, ret[l-1]*2)
			}
		case "+":
			if l > 1 {
				ret = append(ret, ret[l-1]+ret[l-2])
			}
		default:
            n, _ := strconv.Atoi(i)
			ret = append(ret, n)
		}
	}
    var sum int
	for _, i := range ret {
		sum += i
	}
	return sum
}
