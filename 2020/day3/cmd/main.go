package main

import (
	"fmt"
	"os"

	"github.com/ikaradimas/adventofcode/2020/day3"
)

func main() {
	m, err := day3.ReadMap(os.Args[1])
	if err != nil {
		fmt.Printf("reading file, %v\n", err)
	}

	m.Open = '.'
	m.Tree = '#'

	slope := day3.Slope{
		Movements: []day3.Movement{
			{
				Steps:     1,
				Direction: day3.Right,
			},
			{
				Steps:     2,
				Direction: day3.Bottom,
			},
		},
	}

	trees, err := m.CountTrees(slope)
	if err != nil {
		fmt.Printf("counting trees, %v\n", err)
	}

	fmt.Printf("Total trees in map: %d\n", trees)
}

// part 1 result is 280
// part 2: 77, 280, 74, 78, 35
