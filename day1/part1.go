package main

import (
	"AdventOfCode2018/utils"
	"fmt"
)

func main() {
	data := utils.ReadFromFile("day1.txt")
	counter := 0

	for _, v := range data {
		counter = counter + utils.FormatInt(v)
	}

	fmt.Println("End counter is: ", counter)

	main2(data)
}
