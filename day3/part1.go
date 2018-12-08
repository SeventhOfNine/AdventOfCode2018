package main

import (
	"AdventOfCode2018/utils"
	"fmt"
)

type claim struct {
	x, y          int
	height, width int
	id            int
	coords        []coordinate
}

type coordinate struct {
	x, y int
}

func main() {
	data := utils.ReadFromFile("day3.txt")
	canvas, claims := parseClaims(data)
	totalClaims := 0

	for _, v := range canvas {
		if v >= 2 {
			totalClaims += 1
		}
	}

	fmt.Println("Total overlapping claims: ", totalClaims)

	main2(canvas, claims)
}

func parseClaims(data []string) (map[coordinate]int, []claim) {
	canvas := map[coordinate]int{}
	claims := []claim{}

	for _, line := range data {
		claim := claim{}
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &claim.id, &claim.x, &claim.y, &claim.width, &claim.height)

		for row := claim.x; row < claim.x+claim.width; row++ {
			for col := claim.y; col < claim.y+claim.height; col++ {
				coord := coordinate{
					x: row,
					y: col,
				}
				canvas[coord] += 1
				claim.coords = append(claim.coords, coord)
			}
		}
		claims = append(claims, claim)
	}
	return canvas, claims
}
