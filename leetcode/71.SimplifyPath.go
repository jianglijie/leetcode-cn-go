package main

import (
	"bytes"
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	if path == "/../" {
		return "/"
	}
	var s []string
	tmp := strings.Split(path, "/")
	for _, i := range tmp {
		if i == "." || i == "" {
			continue
		} else if i == ".."{
			if len(s) > 0{
				s = s[:len(s)-1]
			}
		} else {
			s = append(s, i)
		}
	}
	if len(s) == 0 {
		return "/"
	} else {
		var ret bytes.Buffer
		for _, i := range s {
			ret.WriteRune('/')
			ret.WriteString(i)
		}
		return ret.String()
	}

}

func main() {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}
