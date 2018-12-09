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

func cloneByteSlice(s []byte) []byte {
	clone := make([]byte, len(s))
	copy(clone, s)
	return clone
}

func stripType(polymer []byte, unitType byte) []byte {
	// Normalize input to uppercase A-Z
	if unitType > 90 {
		unitType = unitType - 32
	}

	for i := 0; i < len(polymer); {
		if polymer[i] == unitType || polymer[i] == unitType+32 {
			polymer = append(polymer[:i], polymer[i+1:]...)
		} else {
			i++
		}
	}

	return polymer
}

func shortestStrippedPolymer(polymer []byte) []byte {
	var i byte
	ssp := polymer

	for i = 65; i <= 90; i++ { // A - Z == 65 - 90
		strippedPolymer := cloneByteSlice(polymer)
		strippedPolymer = reactPolymer(stripType(strippedPolymer, i))
		if len(strippedPolymer) < len(ssp) {
			ssp = strippedPolymer
		}
	}

	return ssp
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	polymer := cloneByteSlice(input)
	fmt.Printf("Part One: %d\n", len(reactPolymer(polymer)))

	polymer = cloneByteSlice(input)
	fmt.Printf("Part Two: %d\n", len(shortestStrippedPolymer(polymer)))
}
