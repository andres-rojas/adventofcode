package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type Record struct {
	Timestamp time.Time
	Guard     int
	Action    string
}

type Records []Record

// Implement Sort interface
func (r Records) Len() int           { return len(r) }
func (r Records) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Records) Less(i, j int) bool { return r[i].Timestamp.Before(r[j].Timestamp) }

func readInput(file string) ([]string, error) {
	var input []string

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	// Known issue: This doesn't handle potential buffer overruns in bufio.ReadLine
	// The shape of the input makes this unlikely to be an issue
	line, _, err := reader.ReadLine()
	for err == nil {
		str := string(line)
		if str != "" {
			input = append(input, str)
		}
		line, _, err = reader.ReadLine()
	}
	if err != io.EOF {
		return nil, err
	}

	return input, nil
}

func parseRecords(input []string) (Records, error) {
	var records Records

	for _, line := range input {
		r := regexp.MustCompile("^[[]([0-9]{4}-[0-1][0-9]-[0-3][0-9] [0-2][0-3]:[0-5][0-9])] (.*)$")
		match := r.FindStringSubmatch(line)

		var err error
		record := Record{}

		timeLayout := "2006-01-02 15:04"
		record.Timestamp, err = time.Parse(timeLayout, match[1])
		if err != nil {
			return records, err
		}

		r = regexp.MustCompile("^Guard #([0-9]+) (.*)")
		if r.MatchString(match[2]) {
			match = r.FindStringSubmatch(match[2])
			record.Guard, err = strconv.Atoi(match[1])
			if err != nil {
				return records, err
			}
			record.Action = match[2]
		} else {
			record.Action = match[2]
		}

		records = append(records, record)
	}

	sort.Sort(Records(records))
	return records, nil
}

func minutesAsleep(records Records) map[int][60]int {
	var guard, from int

	parsedRecords := make(map[int]map[int]int)
	for _, record := range records {
		switch record.Action {
		case "begins shift":
			guard = record.Guard
			if _, ok := parsedRecords[guard]; !ok {
				parsedRecords[guard] = make(map[int]int)
			}
		case "falls asleep":
			from = record.Timestamp.Minute()
		case "wakes up":
			for i := from; i < record.Timestamp.Minute(); i++ {
				parsedRecords[guard][i]++
			}
		}
	}

	minutesSlept := make(map[int][60]int)
	for guard, times := range parsedRecords {
		var minutes [60]int
		for minute, duration := range times {
			minutes[minute] = duration
		}
		minutesSlept[guard] = minutes
	}

	return minutesSlept
}

func mostAsleep(minutesSlept map[int][60]int) int {
	var laziestGuard, currentMax int

	for guard, times := range minutesSlept {
		var guardTotal int
		for _, minutes := range times {
			guardTotal = guardTotal + minutes
		}
		if guardTotal > currentMax {
			currentMax = guardTotal
			laziestGuard = guard
		}
	}

	return laziestGuard
}

func mostFrequentMinuteAsleep(minutesSlept map[int][60]int, guard int) int {
	var laziestMinute, currentMax int

	for minute, duration := range minutesSlept[guard] {
		if duration > currentMax {
			currentMax = duration
			laziestMinute = minute
		}
	}

	return laziestMinute
}

func peakSleepByMinute(minutesSlept map[int][60]int) [60][2]int {
	var psbm [60][2]int

	for guard, times := range minutesSlept {
		for minute, duration := range times {
			if duration > psbm[minute][1] {
				psbm[minute] = [2]int{guard, duration}
			}
		}
	}

	return psbm
}

func mostFrequentSleeper(psbm [60][2]int) (int, int) {
	var guard, peakMinute, currentMax int

	for minute, record := range psbm {
		if record[1] > currentMax {
			guard, peakMinute, currentMax = record[0], minute, record[1]
		}
	}

	return guard, peakMinute
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	input, err := readInput("input.txt")
	check(err)

	records, err := parseRecords(input)
	check(err)

	minutesSlept := minutesAsleep(records)
	laziestGuard := mostAsleep(minutesSlept)
	laziestMinute := mostFrequentMinuteAsleep(minutesSlept, laziestGuard)
	fmt.Printf("Part One: %d\n", laziestGuard*laziestMinute)

	psbm := peakSleepByMinute(minutesSlept)
	sleepiestGuard, peakMinute := mostFrequentSleeper(psbm)
	fmt.Printf("Part Two: %d\n", sleepiestGuard*peakMinute)
}
