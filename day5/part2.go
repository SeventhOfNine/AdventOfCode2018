package main

import "strings"

func part2main(line string) int {
	min := len(line)
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		newline := strings.Replace(line, string(r), "", -1)
		newline = strings.Replace(newline, strings.ToUpper(string(r)), "", -1)

		m := reactPolymer(newline)
		if m < min {
			min = m
		}
	}
	return min
}
