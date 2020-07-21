package main

func reverseBits(num uint32) uint32 {
    res, p := uint32(0), 31
    for num > 0{
        res += (num & 1) << p
        num = num >> 1
        p--
    }
    return res
}

func main() {
	
}
