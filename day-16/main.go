package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	X int
	Y int
}

type Arrow struct {
	Position Coord
	Dir      Direction
}

type Direction string

const (
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
	North Direction = "N"
)

type Component string

const (
	VerticalSplitter   Component = "|"
	HorizontalSplitter Component = "-"
	EmptySpace         Component = "."
	MirrorUpRight      Component = "/"
	MirrorUpLeft       Component = "\\"
)

func stringsToComponents(strList []string) []Component {
	components := make([]Component, len(strList))
	for i, str := range strList {
		components[i] = Component(str)
	}
	return components
}

func parseFile(filePath string) ([][]Component, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines [][]Component
	for scanner.Scan() {
		lineStr := scanner.Text()
		line := strings.Split(lineStr, "")
		lines = append(lines, stringsToComponents(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return lines, nil
}

func move(grid [][]Component, arrow Arrow, visited map[Arrow]bool) {
	maxY := len(grid) - 1
	maxX := len(grid[0]) - 1
	component := grid[arrow.Position.Y][arrow.Position.X]
	newArrows := []Arrow{}

	if _, ok := (visited)[arrow]; ok {
		return
	}
	(visited)[arrow] = true

	switch arrow.Dir {
	case East:
		switch component {
		case VerticalSplitter:
			if arrow.Position.Y > 0 {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X, Y: arrow.Position.Y - 1},
					Dir:      North,
				})
			}
			if arrow.Position.Y < maxY {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X, Y: arrow.Position.Y + 1},
					Dir:      South,
				})
			}
		case HorizontalSplitter:
			if arrow.Position.X < maxX {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X + 1, Y: arrow.Position.Y},
					Dir:      East,
				})
			}
		case MirrorUpRight:
			if arrow.Position.Y > 0 {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X, Y: arrow.Position.Y - 1},
					Dir:      North,
				})
			}
		case MirrorUpLeft:
			if arrow.Position.Y < maxY {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X, Y: arrow.Position.Y + 1},
					Dir:      South,
				})
			}
		case EmptySpace:
			if arrow.Position.X < maxX {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X + 1, Y: arrow.Position.Y},
					Dir:      East,
				})
			}
		}
	case South:
		switch component {
		case VerticalSplitter:
			if arrow.Position.Y < maxY {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X, Y: arrow.Position.Y + 1},
					Dir:      South,
				})
			}
		case HorizontalSplitter:
			if arrow.Position.X > 0 {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X - 1, Y: arrow.Position.Y},
					Dir:      West,
				})
			}
			if arrow.Position.X < maxX {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X + 1, Y: arrow.Position.Y},
					Dir:      East,
				})
			}
		case MirrorUpRight:
			if arrow.Position.X > 0 {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X - 1, Y: arrow.Position.Y},
					Dir:      West,
				})
			}
		case MirrorUpLeft:
			if arrow.Position.X < maxX {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X + 1, Y: arrow.Position.Y},
					Dir:      East,
				})
			}
		case EmptySpace:
			if arrow.Position.Y < maxY {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X, Y: arrow.Position.Y + 1},
					Dir:      South,
				})
			}
		}
	case West:
		switch component {
		case VerticalSplitter:
			if arrow.Position.Y > 0 {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X, Y: arrow.Position.Y - 1},
					Dir:      North,
				})
			}
			if arrow.Position.Y < maxY {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X, Y: arrow.Position.Y + 1},
					Dir:      South,
				})	
			}
		case HorizontalSplitter:
			if arrow.Position.X > 0 {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X - 1, Y: arrow.Position.Y},
					Dir:      West,
				})
			}
		case MirrorUpRight:
			if arrow.Position.Y < maxY {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X, Y: arrow.Position.Y + 1},
					Dir:      South,
				})
			}
		case MirrorUpLeft:
			if arrow.Position.Y > 0 {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X, Y: arrow.Position.Y - 1},
					Dir:      North,
				})
			}
		case EmptySpace:
			if arrow.Position.X > 0 {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X - 1, Y: arrow.Position.Y},
					Dir:      West,
				})
			}
		}
	case North:
		switch component {
		case VerticalSplitter:
			if arrow.Position.Y > 0 {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X, Y: arrow.Position.Y - 1},
					Dir:      North,
				})
			}
		case HorizontalSplitter:
			if arrow.Position.X > 0 {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X - 1, Y: arrow.Position.Y},
					Dir:      West,
				})
			}
			if arrow.Position.X < maxX {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X + 1, Y: arrow.Position.Y},
					Dir:      East,
				})
			}
		case MirrorUpRight:
			if arrow.Position.X < maxX {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X + 1, Y: arrow.Position.Y},
					Dir:      East,
				})
			}
		case MirrorUpLeft:
			if arrow.Position.X > 0 {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X - 1, Y: arrow.Position.Y},
					Dir:      West,
				})
			}
		case EmptySpace:
			if arrow.Position.Y > 0 {
				newArrows = append(newArrows, Arrow{
					Position: Coord{X: arrow.Position.X, Y: arrow.Position.Y - 1},
					Dir:      North,
				})
			}
		}
	}

	for _, newArrow := range newArrows {
		move(grid, newArrow, visited)
	}

}

func countUniqueCoordinates(grid [][]Component, initialArrow Arrow) int {
	arrow := Arrow{
		Position: initialArrow.Position,
		Dir:      initialArrow.Dir,
	}
	visited := make(map[Arrow]bool)
	uniqueCoords := make(map[Coord]bool)

	move(grid, arrow, visited)

	for arrow := range visited {
		uniqueCoords[arrow.Position] = true
	}

	return len(uniqueCoords)
}

func part_1(grid [][]Component) {
	initialArrow := Arrow{
		Position: Coord{X: 0, Y: 0},
		Dir:      East,
	}
	energizedCount := countUniqueCoordinates(grid, initialArrow)
	fmt.Println("Energized count:", energizedCount)
}

func part_2(grid [][]Component) {
	maxX := len(grid[0]) - 1
	maxY := len(grid) - 1

	maxEnergizedCount := 0

	for i := 0; i < len(grid); i++ {
		westSideArrow := Arrow{
			Position: Coord{X: 0, Y: i},
			Dir:      East,
		}
		westSideEnergizedCount := countUniqueCoordinates(grid, westSideArrow)
		if westSideEnergizedCount > maxEnergizedCount {
			maxEnergizedCount = westSideEnergizedCount
		}

		eastSideArrow := Arrow{
			Position: Coord{X: maxX, Y: i},
			Dir:      West,
		}
		eastSideEnergizedCount := countUniqueCoordinates(grid, eastSideArrow)
		if eastSideEnergizedCount > maxEnergizedCount {
			maxEnergizedCount = eastSideEnergizedCount
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		northSideArrow := Arrow{
			Position: Coord{X: i, Y: 0},
			Dir:      South,
		}
		northSideEnergizedCount := countUniqueCoordinates(grid, northSideArrow)
		if northSideEnergizedCount > maxEnergizedCount {
			maxEnergizedCount = northSideEnergizedCount
		}

		southSideArrow := Arrow{
			Position: Coord{X: i, Y: maxY},
			Dir:      North,
		}
		southSideEnergisedCount := countUniqueCoordinates(grid, southSideArrow)
		if southSideEnergisedCount > maxEnergizedCount {
			maxEnergizedCount = southSideEnergisedCount
		}
	}
	fmt.Println("Max energized count:", maxEnergizedCount)

}

func main() {
	filename := "full.txt"

	grid, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	part_1(grid)
	part_2(grid)
}
