package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func readInput(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Assumes input is on a single line
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	line := scanner.Text()

	input := make([]byte, len(line))
	copy(input, line)
	return input, nil
}

func reactPolymer(polymer []byte) []byte {
	// Assumes that all units are represented by standard ASCII alphabet chars
	// Further assumes that the absolute difference between a reactive pair
	// is equivalent to the distance between lowercase and uppercase (i.e. 32)
	for i, j := 0, 1; j < len(polymer); i, j = i+1, j+1 {
		if math.Abs(float64(polymer[i])-float64(polymer[j])) == 32 {
			return reactPolymer(append(polymer[:i], polymer[j+1:]...))
		}
	}

	return polymer
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d\n", len(reactPolymer(input)))
}
