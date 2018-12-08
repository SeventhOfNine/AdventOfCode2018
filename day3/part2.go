package main

import "fmt"

func main2(canvas map[coordinate]int, claims []claim) {
	for _, claim := range claims {
		found := true

		for _, coord := range claim.coords {
			if canvas[coord] != 1 {
				found = false
				break
			}
		}

		if found {
			fmt.Println("Unique claim: ", claim.id)
		}
	}
}
