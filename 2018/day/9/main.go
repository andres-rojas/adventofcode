package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func readInput(file string) (int, int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	// Assumes input is on a single line
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	if scanner.Err() != nil {
		return 0, 0, scanner.Err()
	}
	line := scanner.Text()

	r := regexp.MustCompile("^([0-9]+) players; last marble is worth ([0-9]+) points$")
	match := r.FindStringSubmatch(line)

	playerCount, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, 0, nil
	}
	lastMarble, err := strconv.Atoi(match[2])
	if err != nil {
		return 0, 0, nil
	}

	return playerCount, lastMarble, nil
}

func playGame(players int, marbles int) []int {
	var turn, current, player, next int
	scores := make([]int, players)
	circle := []int{0}

	for turn = 1; turn <= marbles; turn++ {
		if turn%23 == 0 {
			next = current - 7
			if next < 0 { // Wrap around circle counter-clockwise
				next = len(circle) + next
			}

			scores[player] = scores[player] + turn + circle[next]
			circle = append(circle[:next], circle[next+1:]...) // Remove marble from circle
		} else {
			if turn > 1 {
				next = current + 2
				if next > len(circle) { // Wrap around circle clockwise
					next = next % len(circle)
				}
			} else {
				next = 1
			}

			if next == len(circle) { // Add to "end" of circle
				circle = append(circle, turn)
			} else { // Insert marble into circle
				circle = append(circle, 0)
				copy(circle[next+1:], circle[next:])
				circle[next] = turn
			}
		}
		current = next

		player = (player + 1) % players // Next player
	}

	return scores
}

func highScore(scores []int) int {
	sort.Ints(scores)
	return scores[len(scores)-1]
}

func main() {
	playerCount, lastMarble, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	scores := playGame(playerCount, lastMarble)
	fmt.Printf("Part One: %d\n", highScore(scores))
}
