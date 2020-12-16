package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fileName := os.Args[1]
	numbers, err := readNumbers(fileName)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	complementOf2, err := complementOf(2020, 2, numbers)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("complementOf(2): %v, Product: %v\n", complementOf2, product(complementOf2))

	complementOf3, err := complementOf(2020, 3, numbers)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("complementOf(3): %v, Product: %v\n", complementOf3, product(complementOf3))
}

func complementOf(target, parts int, numbers []int) ([]int, error) {
	iters := int(math.Pow(float64(len(numbers)), float64(parts)))

	for iter := 0; iter < iters; iter++ {
		sum := 0
		candidates := []int{}
		bound := int(math.Pow(float64(len(numbers)), float64(parts)))
		acc := iter

		for part := parts; part >= 1; part-- {
			var index, candidate int

			if part == 1 {
				index = acc % bound
			} else {
				bound = bound / len(numbers)
				index = acc / bound
				acc = acc - (index * bound)
			}

			candidate = numbers[index]
			candidates = append(candidates, candidate)
			sum = sum + candidate
		}

		if sum == target {
			return candidates, nil
		}
	}

	return nil, fmt.Errorf("Could not find pair to add up to %d", target)
}

func product(input []int) int64 {
	var result int64
	for _, i := range input {
		if result == 0 {
			result = int64(i)
			continue
		}
		result = result * int64(i)
	}
	return result
}

func readNumbers(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []int{}
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}
