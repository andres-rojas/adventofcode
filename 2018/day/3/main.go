package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
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

func parseClaim(claim string) (int64, [2]int64, [2]int64, error) {
	var id int64
	var position [2]int64
	var size [2]int64

	r := regexp.MustCompile("^#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)$")
	match := r.FindStringSubmatch(claim)

	id, err := strconv.ParseInt(match[1], 10, 64)
	if err != nil {
		return id, position, size, err
	}

	posX, err := strconv.ParseInt(match[2], 10, 64)
	if err != nil {
		return id, position, size, err
	}
	posY, err := strconv.ParseInt(match[3], 10, 64)
	if err != nil {
		return id, position, size, err
	}
	position = [2]int64{posX, posY}

	width, err := strconv.ParseInt(match[4], 10, 64)
	if err != nil {
		return id, position, size, err
	}
	height, err := strconv.ParseInt(match[5], 10, 64)
	if err != nil {
		return id, position, size, err
	}
	size = [2]int64{width, height}

	return id, position, size, nil
}

//// This was used to confirm the max size of the fabric is 1000x1000
// func totalSize(input []string) ([2]int64, error) {
// 	var maxSize [2]int64

// 	for _, claim := range input {
// 		var localMaxSize [2]int64

// 		_, claimPosition, claimSize, err := parseClaim(claim)
// 		if err != nil {
// 			return maxSize, err
// 		}
// 		localMaxSize = [2]int64{claimPosition[0] + claimSize[0], claimPosition[1] + claimSize[1]}

// 		if localMaxSize[0] > maxSize[0] {
// 			maxSize[0] = localMaxSize[0]
// 		}
// 		if localMaxSize[1] > maxSize[1] {
// 			maxSize[1] = localMaxSize[1]
// 		}
// 	}

// 	maxSize[0]++
// 	maxSize[1]++

// 	return maxSize, nil
// }

func mapClaims(input []string) ([1000][1000][]int64, error) {
	var fabric [1000][1000][]int64

	for _, claim := range input {
		claimID, claimPosition, claimSize, err := parseClaim(claim)
		if err != nil {
			return fabric, err
		}

		for i := claimPosition[0]; i < claimPosition[0]+claimSize[0]; i++ {
			for j := claimPosition[1]; j < claimPosition[1]+claimSize[1]; j++ {
				fabric[i][j] = append(fabric[i][j], claimID)
			}
		}
	}

	return fabric, nil
}

func overlappedClaims(input []string) (int64, error) {
	var count int64

	fabric, err := mapClaims(input)
	if err != nil {
		return 0, err
	}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if len(fabric[i][j]) > 1 {
				count++
			}
		}
	}

	return count, nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	input, err := readInput("input.txt")
	check(err)

	count, err := overlappedClaims(input)
	check(err)
	fmt.Printf("Part One: %d\n", count)
}
