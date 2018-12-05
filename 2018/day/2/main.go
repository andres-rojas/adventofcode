package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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

func characterCount(str string) map[rune]int64 {
	runes := make(map[rune]int64)
	for _, char := range str {
		if _, ok := runes[char]; !ok {
			runes[char] = 0
		}
		runes[char] = runes[char] + 1
	}
	return runes
}

func containsSetAmount(str string, amount int64) bool {
	characters := characterCount(str)
	for _, count := range characters {
		if count == amount {
			return true
		}
	}
	return false
}

func checksum(input []string) int64 {
	count := []int64{0, 0}

	for _, str := range input {
		if containsSetAmount(str, 2) {
			count[0]++
		}
		if containsSetAmount(str, 3) {
			count[1]++
		}
	}

	return count[0] * count[1]
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", checksum(input))
}
