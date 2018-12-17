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
		return 0, 0, err
	}
	lastMarble, err := strconv.Atoi(match[2])
	if err != nil {
		return 0, 0, err
	}

	return playerCount, lastMarble, nil
}

type Marble struct {
	Prev, Next *Marble
	Value      int
}

func (marble *Marble) InsertAfter(value int) *Marble {
	inserted := &Marble{Prev: marble, Next: marble.Next, Value: value}

	marble.Next.Prev = inserted
	marble.Next = inserted

	return inserted
}

func (marble *Marble) Remove() (int, *Marble) {
	marble.Prev.Next = marble.Next
	marble.Next.Prev = marble.Prev

	return marble.Value, marble.Next
}

func playGame(players int, marbles int) []int {
	scores := make([]int, players)
	current := &Marble{Value: 0}
	current.Prev, current.Next = current, current

	for player, marble := 0, 1; marble <= marbles; player, marble = (player+1)%players, marble+1 {
		if marble%23 == 0 {
			current = current.Prev.Prev.Prev.Prev.Prev.Prev.Prev
			removed, next := current.Remove()
			scores[player] = scores[player] + marble + removed
			current = next
		} else {
			current = current.Next.InsertAfter(marble)
		}
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

	scores = playGame(playerCount, lastMarble*100)
	fmt.Printf("Part Two: %d\n", highScore(scores))
}
