package day3

import (
	"bufio"
	"fmt"
	"os"
)

// Slope expresses a vector of arbitrary movements.
type Slope struct {
	Movements []Movement
}

// Movement expresses a move of X steps to an arbitrary direction.
type Movement struct {
	Direction Direction
	Steps     int
}

// Map expresses a map with a defined number of lines and columns.
type Map struct {
	Columns int
	Lines   []string
	Open    byte
	Tree    byte
}

// Direction expresses a direction to which movement is allowed.
type Direction string

// This is the set of directions possible.
const (
	Left   Direction = "left"
	Right  Direction = "right"
	Top    Direction = "top"
	Bottom Direction = "bottom"
)

// Position represents a typical 2D position.
type Position struct {
	X int
	Y int
}

// ReadMap reads a map from a file.
func ReadMap(fileName string) (*Map, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	columns := -1
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
		if columns == -1 {
			columns = len(text)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Map{
		Columns: columns,
		Lines:   lines,
	}, nil
}

// GetByteAt will return the rune at position (x, y)
// in the file. It may also pan arbitrarily
// by repeating lines or columns, as warranted.
func (m *Map) GetByteAt(position Position) (byte, error) {
	// x pans, so modulo it.
	aX := position.X % m.Columns
	// y doesn't pan, so check for bounds.
	if position.Y < 0 || position.Y >= len(m.Lines) {
		return ' ', fmt.Errorf("y is out of bounds")
	}

	return m.Lines[position.Y][aX], nil
}

// CountTrees will count trees encountered while moving
// on a determined slope, until the end of the map
// has been reached.
func (m *Map) CountTrees(slope Slope) (int, error) {
	var trees int
	position := Position{X: 0, Y: 0}

	for {
		position = m.MoveBy(position, slope)
		b, err := m.GetByteAt(position)
		if err != nil {
			// only one case for an error: end reached.
			break
		}
		if b == m.Tree {
			trees++
		}
	}

	return trees, nil
}

// MoveBy will attempt to move to a new position given a slope
// and return the new Position, or an error if it cannot
// do so.
func (m *Map) MoveBy(start Position, slope Slope) Position {
	offsetX, offsetY := 0, 0

	for _, m := range slope.Movements {
		switch m.Direction {
		case Left:
			offsetX = offsetX - m.Steps
			break
		case Right:
			offsetX = offsetX + m.Steps
			break
		case Top:
			offsetY = offsetY - m.Steps
		case Bottom:
			offsetY = offsetY + m.Steps
		}
	}

	return Position{
		X: start.X + offsetX,
		Y: start.Y + offsetY,
	}
}
