package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Parent   *Node
	Children []*Node
	Metadata []int
}

func readInput(file string) ([]int, error) {
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

	var input []int
	for _, str := range strings.Fields(line) {
		i, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		input = append(input, i)
	}

	return input, nil
}

func parseInput(input []int, parent *Node) (*Node, int) {
	node := Node{Parent: parent}
	i := 0

	qtyChildren := input[i]
	i++
	qtyMetadata := input[i]
	i++

	for j := 0; j < qtyChildren; j++ {
		child, size := parseInput(input[i:], &node)
		node.Children = append(node.Children, child)
		i = i + size
	}

	node.Metadata = input[i : i+qtyMetadata]
	i = i + qtyMetadata

	return &node, i
}

func (tree Node) totalMetadata() int {
	total := 0

	for _, child := range tree.Children {
		total = total + child.totalMetadata()
	}
	for _, metadata := range tree.Metadata {
		total = total + metadata
	}

	return total
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	tree, _ := parseInput(input, nil)
	fmt.Printf("Part One: %d\n", tree.totalMetadata())
}
