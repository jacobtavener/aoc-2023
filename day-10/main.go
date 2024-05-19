package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"math"
	"slices"
)

type Coord struct {
	X int
	Y int
}

type Move struct {
	Delta Coord
	Dir   Direction
}

type Pipe struct {
	Position Coord
	Dir 	Direction
}


type Direction string
const (
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
	North Direction = "N"
)

type PipeType string
const (
	Vertical   PipeType = "|"
	Horizontal PipeType = "-"
	NE         PipeType = "L"
	NW         PipeType = "J"
	SW         PipeType = "7"
	SE         PipeType = "F"
	Start 	   PipeType = "S"	
	Ground	   PipeType = "."	
)

func stringsToPipeTypes(strList []string) []PipeType {
	pipeTypes := make([]PipeType, len(strList))
	for i, str := range strList {
		pipeTypes[i] = PipeType(str)
	}
	return pipeTypes
}

func isValidMove(pipeType PipeType, direction Direction) bool {
	validMoves := map[Direction][]PipeType{
		North: {Vertical, SW, SE},
		East:  {Horizontal, SW, NW},
		South: {Vertical, NE, NW},
		West:  {Horizontal, NE, SE},
	}
	return slices.Contains(validMoves[direction], pipeType)
}

func parseFile(filePath string) ([][]PipeType, []int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines [][]PipeType
	var startPosition []int
	lineNum := 0
	for scanner.Scan() {
		lineStr := scanner.Text()
		if startPosition == nil {
			sPosition := strings.Index(lineStr, string(Start))
			if sPosition != -1 {
				startPosition = []int{lineNum, sPosition}
			}
		}
		line := strings.Split(lineStr, "")
		lines = append(lines, stringsToPipeTypes(line))
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil, err
	}

	return lines, startPosition, nil
}

func findNextPipe(currentPipe Pipe, pipeType PipeType, pipePath *[]Coord) Pipe {
	position, direction := currentPipe.Position, currentPipe.Dir
	y, x := position.Y, position.X
	newX, newY, newDirection := x, y, direction

	switch direction {
	case North:
		switch pipeType {
		case Vertical:
			newY = y-1
			newDirection = North
		case SW:
			newX = x-1
			newDirection = West
		case SE:
			newX = x+1
			newDirection = East
		}
	case East:
		switch pipeType {
		case Horizontal:
			newX = x+1
			newDirection = East
		case SW:
			newY = y+1
			newDirection = South
		case NW:
			newY = y-1
			newDirection = North
		}
	case South:
		switch pipeType {
		case Vertical:
			newY = y+1
			newDirection = South
		case NE:
			newX = x+1
			newDirection = East
		case NW:
			newX = x-1
			newDirection = West
		}
	case West:
		switch pipeType {
		case Horizontal:
			newX = x-1
			newDirection = West
		case NE:
			newY = y-1
			newDirection = North
		case SE:
			newY = y+1
			newDirection = South
		}
	}
	newCoord := Coord{Y:newY, X:newX}
	*pipePath = append(*pipePath, newCoord)
	return Pipe{Position: newCoord, Dir: newDirection}
}

func furthestPointInPipePath(x int) int {
	if x%2 == 0 {
		return x / 2
	} else {
		return (x / 2) + 1
	}
}

func findInitialPipe(grid [][]PipeType, startPosition []int) Pipe {
	adjacentPositions := []Move{
		{Delta: Coord{Y:0, X:1}, Dir: East},
		{Delta: Coord{Y:1, X:0}, Dir: South},
		{Delta: Coord{Y:0, X:-1}, Dir: West},
		{Delta: Coord{Y:-1, X:0}, Dir: North},
	}
	var nextPipe Pipe
	for _, position := range adjacentPositions {
		delta := position.Delta
		direction := position.Dir

		y := startPosition[0] + delta.Y
		x := startPosition[1] + delta.X
		

		if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[0]) {
			continue
		}
		if isValidMove(grid[y][x], direction) {
			nextPipe = Pipe{Position: Coord{Y:y, X:x}, Dir: direction}
			break
		}
	}
	return nextPipe
}


func findPipePath(grid [][]PipeType, currentPipe Pipe, currentPipeType PipeType) []Coord {
	pipePath := make([]Coord, 0)
	for {
		if currentPipeType == Start {
			break
		}
		nextPipe := findNextPipe(currentPipe, currentPipeType, &pipePath)
		currentPipe = nextPipe
		currentPosition := currentPipe.Position
		currentPipeType = grid[currentPosition.Y][currentPosition.X]
	}
	return pipePath
}

func shoelaceFormula(pipePath []Coord) int {
	n := len(pipePath)
	if n < 3 {
		return 0.0
	}
	area := 0
	for i := 0; i < n-1; i++ {
		area += (pipePath[i].X*pipePath[i+1].Y) - (pipePath[i+1].X*pipePath[i].Y)
	}
	area += (pipePath[n-1].X*pipePath[0].Y) - (pipePath[0].X*pipePath[n-1].Y)
	area = int(math.Abs(float64(area)) / 2.0)

	return area
}


func main() {
	filename := "full.txt"

	grid, startPosition, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	initialPipe := findInitialPipe(grid, startPosition)
	curPosition := initialPipe.Position
	initialPipeType := grid[curPosition.Y][curPosition.X]
	pipePath := findPipePath(grid, initialPipe, initialPipeType)

	// Part 1
	pipeLength := len(pipePath)
	fmt.Println("Furthest point in cycle:", furthestPointInPipePath(pipeLength))

	// Part 2
	totalArea := shoelaceFormula(pipePath)
	pipeArea := pipeLength / 2
	fmt.Println("Nest Area: ", totalArea - pipeArea)

		
}