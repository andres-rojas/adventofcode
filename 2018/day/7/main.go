package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
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

func parseDependencies(input []string) map[rune][]rune {
	steps := make(map[rune][]rune)
	r := regexp.MustCompile("^Step ([A-Z]) must be finished before step ([A-Z]) can begin[.]$")

	for _, line := range input {
		match := r.FindStringSubmatch(line)
		dependency, dependent := rune(match[1][0]), rune(match[2][0])
		if _, exists := steps[dependency]; !exists {
			steps[dependency] = make([]rune, 0)
		}
		steps[dependent] = append(steps[dependent], dependency)
	}

	return steps
}

func orderedSteps(steps map[rune][]rune) []rune {
	var oSteps []rune

	for len(steps) > 0 {
		var nextStep rune
		for i, dependencies := range steps {
			if len(dependencies) == 0 && (nextStep == 0 || i < nextStep) {
				nextStep = i
			}
		}

		oSteps = append(oSteps, nextStep)
		delete(steps, nextStep)

		for i, dependencies := range steps {
			for j, dependency := range dependencies {
				if dependency == nextStep {
					steps[i] = append(steps[i][:j], steps[i][j+1:]...)
				}
			}
		}
	}

	return oSteps
}

func timeInParallel(steps map[rune][]rune, workers int) int {
	var second int
	activeWork := make(map[rune]int) // maps [step] to time (second) job is done

	for len(steps) > 0 || len(activeWork) > 0 {
		// Check if any activeWork is done
		for step, done := range activeWork {
			if done == second {
				delete(activeWork, step)

				for i, dependencies := range steps {
					for j, dependency := range dependencies {
						if dependency == step {
							steps[i] = append(steps[i][:j], steps[i][j+1:]...)
						}
					}
				}
			}
		}

		// Start work if workers are available
		if len(activeWork) < workers {
			var readyWork []int
			for i, dependencies := range steps {
				if len(dependencies) == 0 {
					readyWork = append(readyWork, int(i))
				}
			}
			sort.Ints(readyWork)

			freeWorkers := workers - len(activeWork)
			if len(readyWork) > freeWorkers {
				readyWork = readyWork[:freeWorkers]
			}
			for _, step := range readyWork {
				// [A-Z] == [65-90]
				// step A will take 61 seconds (65 - 4 = 61)
				// step B will take 62 seconds (66 - 4 = 62)
				// ...
				done := second + step - 4
				activeWork[rune(step)] = done
				delete(steps, rune(step))
			}
		}

		// Time marches on
		second++
	}

	return second - 1
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	steps := parseDependencies(input)
	fmt.Printf("Part One: %s\n", string(orderedSteps(steps)))

	steps = parseDependencies(input)
	fmt.Printf("Part Two: %d\n", timeInParallel(steps, 5))
}
