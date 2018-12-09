package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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

func parseCoordinates(input []string) ([][2]int, error) {
	var coordinates [][2]int
	for _, line := range input {
		c := strings.Split(line, ", ")

		x, err := strconv.Atoi(c[1])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(c[0])
		if err != nil {
			return nil, err
		}

		coordinates = append(coordinates, [2]int{x, y})
	}

	return coordinates, nil
}

func initPlane(coordinates [][2]int) [][]int {
	max := 0
	for _, coordinate := range coordinates {
		for _, i := range coordinate {
			if i > max {
				max = i
			}
		}
	}
	max++

	plane := make([][]int, max)
	for i := range plane {
		plane[i] = make([]int, max)
	}

	return plane
}

func mapCoordinates(coordinates [][2]int) [][]int {
	plane := initPlane(coordinates)
	for i, coordinate := range coordinates {
		plane[coordinate[0]][coordinate[1]] = i + 1
	}

	return plane
}

func manhattanDistance(p [2]int, q [2]int) int {
	dist := math.Abs(float64(p[0])-float64(q[0])) + math.Abs(float64(p[1])-float64(q[1]))
	return int(dist)
}

func closestPoint(point [2]int, coordinates [][2]int) int {
	distances := make(map[int][]int)

	for i, coordinate := range coordinates {
		distance := manhattanDistance(point, coordinate)
		distances[distance] = append(distances[distance], i+1)
	}

	var sortDistances []int
	for i := range distances {
		sortDistances = append(sortDistances, i)
	}
	sort.Ints(sortDistances)
	closestDistance := sortDistances[0]

	if len(distances[closestDistance]) > 1 {
		return 0
	}

	return distances[closestDistance][0]
}

func plotClosest(coordinates [][2]int) [][]int {
	plane := mapCoordinates(coordinates)
	for i, row := range plane {
		for j, point := range row {
			if point == 0 {
				plane[i][j] = closestPoint([2]int{i, j}, coordinates)
			}
		}
	}

	return plane
}

func largestArea(plane [][]int) int {
	areas := make(map[int]int)

	for i, row := range plane {
		for j, point := range row {
			// If a point touches the edges, that indicates an infinte area
			if i == 0 || i == len(plane)-1 || j == 0 || j == len(row)-1 {
				areas[point] = -1
			}

			// Ignore infinites
			if areas[point] != -1 {
				areas[point]++
			}
		}
	}

	var sortedAreas []int
	for _, area := range areas {
		sortedAreas = append(sortedAreas, area)
	}
	sort.Ints(sortedAreas)

	last := len(sortedAreas) - 1
	if sortedAreas[last] == 0 {
		return sortedAreas[last-1]
	}
	return sortedAreas[last]
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	input, err := readInput("input.txt")
	check(err)

	coordinates, err := parseCoordinates(input)
	check(err)

	plane := plotClosest(coordinates)
	fmt.Printf("Part One: %d\n", largestArea(plane))
}
