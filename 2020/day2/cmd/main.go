package main

import (
	"fmt"
	"os"

	"github.com/ikaradimas/adventofcode/2020/day2"
)

func main() {
	passwords, err := day2.ReadPasswords(os.Args[1])
	if err != nil {
		fmt.Printf("reading file %v\n", err)
		os.Exit(1)
	}

	var count int
	for _, p := range passwords {
		if p.IsValidNew() {
			count++
		}
	}

	fmt.Printf("Total correct passwords: %d\n", count)
}
