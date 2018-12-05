package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/deckarep/golang-set"
)

func readInput(file string) ([]int64, error) {
	var input []int64

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
			i, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return nil, err
			}
			input = append(input, i)
		}
		line, _, err = reader.ReadLine()
	}
	if err != io.EOF {
		return nil, err
	}

	return input, nil
}

func calcFrequency(input []int64) int64 {
	var frequency int64 = 0
	for _, change := range input {
		frequency = frequency + change
	}
	return frequency
}

func repeatedFrequency(input []int64) int64 {
	var frequency int64 = 0
	frequencies := mapset.NewSet()

	for {
		for _, change := range input {
			frequency = frequency + change
			if frequencies.Contains(frequency) {
				return frequency
			}
			frequencies.Add(frequency)
		}
	}
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", calcFrequency(input))
	fmt.Printf("Part Two: %d\n", repeatedFrequency(input))
}
