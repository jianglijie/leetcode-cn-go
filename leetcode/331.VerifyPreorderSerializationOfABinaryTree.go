package main

import "strings"

func isValidSerialization(preorder string) bool {
    leaves := 0
    nodes := 0
    tmp := strings.Split(preorder,",")
	for i, s := range tmp {
		if s == "#" {
			leaves ++
		} else {
			nodes ++
		}
		if leaves > nodes + 1{
			return false
		}
		if leaves == nodes + 1 && i < len(tmp)- 1 {
			return false
		}

	}
	if leaves == nodes + 1{
		return true
	} else {
		return false
	}
}
