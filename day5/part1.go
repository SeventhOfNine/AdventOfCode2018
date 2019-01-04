package main

import (
	"AdventOfCode2018/utils"
	"fmt"
	"strings"
)

func main() {
	data := utils.ReadStringFromFile("day5.txt")
	line := strings.TrimSpace(data)

	part1Length := reactPolymer(line)
	fmt.Println("part 1 solution: ", part1Length)

	part2Length := part2main(line)
	fmt.Println("part 2 solution: ", part2Length)
}

func reactPolymer(line string) int {
	buffer := []byte{line[0]}

	for i := 1; i < len(line); i++ {
		length := len(buffer)
		lastCharInBuffer := ""
		currentCharInLine := string(line[i])

		if length != 0 {
			lastCharInBuffer = string(buffer[length-1])
		}

		if lastCharInBuffer != currentCharInLine && strings.EqualFold(lastCharInBuffer, currentCharInLine) {
			buffer[length-1] = 0
			buffer = buffer[:length-1]
		} else {
			buffer = append(buffer, line[i])
		}
	}

	return len(buffer)
}
