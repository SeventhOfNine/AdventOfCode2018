package main

import (
	"fmt"
)

func main2(data []string) {

	for idx, line := range data[0 : len(data)-1] {
		for i := idx + 1; i < len(data)-1-idx; i++ {
			res := intersect(line, data[i])
			if len(res) == len(line)-1 {
				fmt.Println("intersect     : ", res)
				fmt.Println("first line is : ", line)
				fmt.Println("second line is: ", data[i])
			}
		}
	}
}

func intersect(a string, b string) string {
	res := ""
	mismatchCount := 0

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			res += string(a[i])
		} else {
			mismatchCount += 1
		}
		if mismatchCount > 1 {
			return ""
		}
	}

	return res
}
