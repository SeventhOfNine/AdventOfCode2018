package main

import (
	"AdventOfCode2018/utils"
	"fmt"
)

func main2(data []string) {
	counter := 0
	seen := map[int]bool{}
	seen[counter] = true
	dupe := false

	for dupe == false {
		for _, v := range data {
			counter = counter + utils.FormatInt(v)

			if seen[counter] {
				fmt.Println("Duplicate value is: ", counter)
				dupe = true
				break
			} else {
				seen[counter] = true
			}
		}
	}
}
