package aoc2018

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

type dayFourChange string

const dayFourChangeStartShift = "start shift"

var dayFourChangeSleep = "falls asleep"
var dayFourChangeWake = "wakes up"

type dayFourRecord struct {
	month  int
	day    int
	guard  int
	hour   int
	minute int
	change dayFourChange
}

type dayFourSleepStartStop struct {
	start int
	stop  int
}

func DayFourPartOne(input string) int {
	records := dayFourParseInput(input)

	maxSleep := make(map[int]int, 0)

	sleep := dayFourChange(dayFourChangeSleep)
	wake := dayFourChange(dayFourChangeWake)

	//	Work out who slept when
	sleepSince := 0
	for _, record := range records {

		if record.change == sleep {
			sleepSince = record.minute
		}

		if record.change == wake {
			slept := record.minute - sleepSince

			if _, ok := maxSleep[record.guard]; ok == false {
				maxSleep[record.guard] = 0
			}

			maxSleep[record.guard] += slept
		}
	}

	//	Find out who slept the longest
	longestSleeper := 0
	longestSleep := 0

	for guard, sleep := range maxSleep {
		if sleep > longestSleep {
			longestSleep = sleep
			longestSleeper = guard
		}
	}

	sleptMinuteNr, _ := dayFourFindMostSleptMinuteOfGuard(records, longestSleeper, sleep, wake)

	log.Printf("%d slept %d minutes", longestSleeper, sleptMinuteNr)

	return sleptMinuteNr * longestSleeper
}

func DayFourPartTwo(input string) int {
	records := dayFourParseInput(input)
	sleep := dayFourChange(dayFourChangeSleep)
	wake := dayFourChange(dayFourChangeWake)

	//	Make a map, for uniqueness
	uniqueGuardMap := make(map[int]bool, 0)
	guards := make([]int, 0)
	for _, r := range records {
		if _, ok := uniqueGuardMap[r.guard]; ok == false {
			uniqueGuardMap[r.guard] = true
			guards = append(guards, r.guard)
		}
	}

	mostSleptMinute := 0
	mostSleptCount := 0
	mostSleptGuard := 0

	for _, guard := range guards {
		minute, count := dayFourFindMostSleptMinuteOfGuard(records, guard, sleep, wake)

		if count > mostSleptCount {
			mostSleptMinute = minute
			mostSleptCount = count
			mostSleptGuard = guard
		}
	}

	return mostSleptMinute * mostSleptGuard
}

func dayFourFindMostSleptMinuteOfGuard(records []dayFourRecord, guard int, sleep dayFourChange, wake dayFourChange) (int, int) {
	//	Find out when that guard slept the most often
	sleepPeriods := make([]dayFourSleepStartStop, 0)
	lastSleepStart := 0
	for _, record := range records {
		if record.guard == guard {
			if record.change == sleep {
				lastSleepStart = record.minute
			}

			if record.change == wake {
				sleepPeriods = append(sleepPeriods, dayFourSleepStartStop{lastSleepStart, record.minute})
			}
		}
	}
	sleptTimes := make(map[int]int, 0)
	for _, sleepPeriod := range sleepPeriods {
		for i := sleepPeriod.start; i < sleepPeriod.stop; i++ {
			if _, ok := sleptTimes[i]; ok == false {
				sleptTimes[i] = 0
			}
			sleptTimes[i]++
		}
	}
	sleptMinuteNr := 0
	sleptMinuteCnt := 0
	for minute, cnt := range sleptTimes {
		if cnt > sleptMinuteCnt {
			sleptMinuteNr = minute
			sleptMinuteCnt = cnt
		}
	}
	return sleptMinuteNr, sleptMinuteCnt
}

func dayFourParseInput(input string) []dayFourRecord {
	lines := strings.Split(input, "\n")
	sort.Strings(lines)

	guard := 0
	records := make([]dayFourRecord, 0)
	for _, line := range lines {
		parts := strings.Split(line, `] `)

		r := dayFourRecord{}
		r.month, _ = strconv.Atoi(parts[0][6:8])
		r.day, _ = strconv.Atoi(parts[0][9:11])
		r.hour, _ = strconv.Atoi(parts[0][12:14])
		r.minute, _ = strconv.Atoi(parts[0][15:17])

		switch parts[1] {
		case dayFourChangeSleep:
			r.guard = guard
			r.change = dayFourChange(parts[1])
			break
		case dayFourChangeWake:
			r.guard = guard
			r.change = dayFourChange(parts[1])
			break
		default:
			guard, _ = strconv.Atoi(strings.Split(strings.Split(parts[1], `#`)[1], ` `)[0])
			r.guard = guard
			r.change = dayFourChange(dayFourChangeStartShift)
		}

		records = append(records, r)
	}

	return records
}
