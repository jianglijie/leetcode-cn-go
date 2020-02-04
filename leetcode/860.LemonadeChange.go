package main

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, i := range bills {
		if i == 5 {
			five++
		} else if i == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five > 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func main() {

}
