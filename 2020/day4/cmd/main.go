package main

import (
	"fmt"
	"os"

	"github.com/ikaradimas/adventofcode/2020/day4"
)

func main() {
	passports, err := day4.ReadPassports(os.Args[1])
	if err != nil {
		fmt.Printf("reading passports %v\n", err)
		os.Exit(1)
	}

	validator := day4.NewSimplePassportValidator()

	count := 0
	for _, passport := range passports {
		if validator.IsValid(passport) {
			count++
		}
	}

	fmt.Printf("Valid passports: %d\n", count)
}
