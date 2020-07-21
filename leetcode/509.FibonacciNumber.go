package main

func fib(N int) int {
    if N == 0 {
		return 0
	}else if N == 1{
		return 1
	}else{
		a, b, c := 0, 1, 0
		for N >= 2{
			c = a + b
			a, b = b, c
			N--
		}
		return c
	}
}

func main() {
	
}
