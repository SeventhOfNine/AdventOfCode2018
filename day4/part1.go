package main

import (
	"AdventOfCode2018/utils"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

// minutes -> num of times found asleep
type timesAsleep map[int]int

// guardid -> timesAsleep
type schedules map[int][]timesAsleep

func main() {
	allSchedules := schedules{}
	guards_to_timeAsleep := map[int]int{}

	parseSchedule(allSchedules, guards_to_timeAsleep)

	guardId := findGuardWhoSpendsMostTimeAsleep(guards_to_timeAsleep)

	findMinuteWhenMostAsleep(guardId, allSchedules)

	main2(allSchedules)
}

func parseSchedule(allSchedules schedules, guards_to_timeAsleep map[int]int) {
	data := utils.ReadFromFile("day4.txt")
	sort.Strings(data)

	var asleepPeriod timesAsleep
	currentGuardId := -1
	asleep_start := -1

	for _, line := range data {
		re := regexp.MustCompile(`^\[(.*)\] (.*)$`)
		matches := re.FindStringSubmatch(line)
		status := matches[2]

		datetime, _ := time.Parse("2006-01-02 15:04", matches[1])

		if strings.Contains(status, "Guard") {
			// guard arrives
			// previous shift ends
			if currentGuardId != -1 {
				changeGuard(allSchedules, asleepPeriod, currentGuardId)
			}

			// save guardid
			reId := regexp.MustCompile(`(\d+)`)
			currentGuardId = utils.FormatInt(reId.FindString(status))
			// create current shift
			asleepPeriod = timesAsleep{}
		} else if strings.Contains(status, "asleep") {
			// falls asleep
			asleep_start = datetime.Minute()
		} else if strings.Contains(status, "wakes") {
			// mark interval as asleep
			wakesUp_start := datetime.Minute()

			for m := asleep_start; m < wakesUp_start; m++ {
				asleepPeriod[m] += 1
				guards_to_timeAsleep[currentGuardId] += 1
			}
		}
	}

	// save last period
	changeGuard(allSchedules, asleepPeriod, currentGuardId)
}

func changeGuard(scheds schedules, period timesAsleep, guardId int) {
	savedSchedule := scheds[guardId]
	savedSchedule = append(savedSchedule, period)
	scheds[guardId] = savedSchedule
}

func findGuardWhoSpendsMostTimeAsleep(guards_to_timeAsleep map[int]int) int {
	// sort guards_to_timeAsleep map
	mirrorMap := map[int]int{}
	minutesAsKeys := []int{}

	for key, val := range guards_to_timeAsleep {
		mirrorMap[val] = key
		minutesAsKeys = append(minutesAsKeys, val)
	}
	sort.Ints(minutesAsKeys)

	longestTimeAsleep := minutesAsKeys[len(minutesAsKeys)-1]
	guardId := mirrorMap[longestTimeAsleep]

	fmt.Println("GuardId: ", guardId)
	fmt.Println("longestTimeAsleep: ", longestTimeAsleep)
	fmt.Println("=================================================")

	return guardId
}

func findMinuteWhenMostAsleep(guardId int, allSchedules schedules) {
	// For identified guardId, get amount of times asleep for each minute
	minutes_to_timesAsleep := map[int]int{}
	for _, minutes := range allSchedules[guardId] {
		for m, t := range minutes {
			minutes_to_timesAsleep[m] += t
		}
	}

	// sort by value to find the minute
	sortedKeys := []int{}
	mirror_minutes_to_timesAsleep := map[int]int{}
	for key, val := range minutes_to_timesAsleep {
		mirror_minutes_to_timesAsleep[val] = key
		sortedKeys = append(sortedKeys, val)
	}
	sort.Ints(sortedKeys)

	idx := sortedKeys[len(sortedKeys)-1]
	minuteMostAsleep := mirror_minutes_to_timesAsleep[idx]

	fmt.Println("minute when most asleep: ", minuteMostAsleep)
	fmt.Println("guardid X minute: ", guardId*minuteMostAsleep)
	fmt.Println("=================================================")
}
