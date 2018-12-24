package main

import "fmt"

func main2(allSchedules schedules) {
	maxTimesPerMinute := -1
	guardId := -1
	minute := -1

	for g, minutes := range allSchedules {
		minutes_asleep := map[int]int{}
		for _, times := range minutes {
			for m, t := range times {
				minutes_asleep[m] += t
			}
		}

		for m, t := range minutes_asleep {
			if t > maxTimesPerMinute {
				maxTimesPerMinute = t
				minute = m
				guardId = g
			}
		}
	}

	fmt.Println("guardid: ", guardId)
	fmt.Println("value: ", guardId*minute)
}
