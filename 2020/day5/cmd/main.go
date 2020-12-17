package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Examples:")
	fmt.Println("-------------------------------------")
	result := getSeat("BFFFBBFRRR")
	fmt.Printf("BFFFBBFRRR, %v\n", result)
	result = getSeat("FFFBBBFRRR")
	fmt.Printf("FFFBBBFRRR, %v\n", result)
	result = getSeat("BBFFBBFRLL")
	fmt.Printf("BBFFBBFRLL, %v\n", result)

	bspCodes, err := readFile(os.Args[1])
	if err != nil {
		fmt.Printf("reading BSP codes: %v\n", err)
		os.Exit(1)
	}

	maxSeatID := 0
	seatIDs := []int{}
	for _, bspCode := range bspCodes {
		seat := getSeat(bspCode)
		if seat.ID > maxSeatID {
			maxSeatID = seat.ID
		}
		seatIDs = append(seatIDs, seat.ID)
	}

	fmt.Println("-------------------------------------")
	fmt.Printf("Max seat ID: %d\n", maxSeatID)

	seatIDs = bubbleSort(seatIDs)

	missingID := 0
	for i := 0; i < len(seatIDs)-1; i++ {
		if seatIDs[i+1]-seatIDs[i] > 1 {
			missingID = seatIDs[i] + 1
		}
	}

	fmt.Println("-------------------------------------")
	fmt.Printf("Missing seat ID: %d\n", missingID)
}

func readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

type seat struct {
	Row    int
	Column int
	ID     int
}

type partition struct {
	LowerBound int
	UpperBound int
}

func bubbleSort(values []int) []int {
	for {
		swaps := 0
		for i := 0; i < len(values)-1; i++ {
			if values[i] > values[i+1] {
				swaps++
				values[i+1], values[i] = values[i], values[i+1]
			}
		}

		if swaps == 0 {
			break
		}
	}

	return values
}

func getSeat(bspCode string) seat {
	const maxRows int = 128
	const maxColumns int = 8

	row := &partition{
		LowerBound: 0,
		UpperBound: maxRows - 1,
	}

	col := &partition{
		LowerBound: 0,
		UpperBound: maxColumns - 1,
	}

	for _, i := range []rune(bspCode) {
		ir := (row.UpperBound - row.LowerBound + 1) / 2
		ic := (col.UpperBound - col.LowerBound + 1) / 2

		switch i {
		case 'F':
			row.UpperBound = row.LowerBound + ir
			break
		case 'B':
			row.LowerBound = row.LowerBound + ir
			break
		case 'R':
			col.LowerBound = col.LowerBound + ic
			break
		case 'L':
			col.UpperBound = col.LowerBound + ic
			break
		}
	}

	return seat{
		Row:    row.LowerBound,
		Column: col.LowerBound,
		ID:     (row.LowerBound * 8) + col.LowerBound,
	}
}
