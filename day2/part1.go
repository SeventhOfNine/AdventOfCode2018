package main

import (
	"AdventOfCode2018/utils"
	"fmt"
)

func main() {
	data := utils.ReadFromFile("day2.txt")
	twoLettersCounter := 0
	threeLettersCounter := 0

	for _, line := range data {
		code := map[byte]int{}

		for k := 0; k < len(line); k++ {
			code[line[k]] += 1
		}

		if seen(code, 2) {
			twoLettersCounter += 1
		}
		if seen(code, 3) {
			threeLettersCounter += 1
		}
	}

	total := twoLettersCounter * threeLettersCounter
	fmt.Println("Checksum is ", total)

	main2(data)
}

func seen(vals map[byte]int, number int) bool {
	for v := range vals {
		if vals[v] == number {
			return true
		}
	}
	return false
}
