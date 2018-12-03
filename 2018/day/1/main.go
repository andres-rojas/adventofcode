package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

func partOne() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	frequency := calcFrequency(input)
	fmt.Printf("Part One: %d\n", frequency)
}

func main() {
	partOne()
}
