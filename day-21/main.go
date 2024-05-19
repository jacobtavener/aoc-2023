package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

type Coord = [2]int



func parseFile(filename string) ([][]string, Coord, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, Coord{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	startPosition := Coord{0,0}
	garden := make([][]string, 0)

	yIndex := 0
	for scanner.Scan() {
		lineStr := scanner.Text()
		xCoords := strings.Split(lineStr, "")
		gardenRow := make([]string, len(xCoords))
		for xIndex, xCoord := range xCoords {
			if xCoord == "S" {
				startPosition = Coord{yIndex, xIndex}
			}
			gardenRow[xIndex] = xCoord
		}
		garden = append(garden, gardenRow)
		yIndex++
	}
	return garden, startPosition, nil

}

func findMoves(startPosition Coord, garden [][]string, numMoves int) int {
	moveCount := 0
	positions := []Coord{startPosition}
	possibleMovements := [4][2]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
	totalPossibleMoves := 0
	for moveCount < numMoves {
		nextPositionSet := make(map[Coord]bool)
		for _, position := range positions {
			for _, movement := range possibleMovements {
				newY := position[0] + movement[0]
				newX := position[1] + movement[1]
				// Roll over logic
				yConsidered := newY % len(garden)
				if yConsidered < 0 {
					yConsidered = len(garden) + yConsidered
				}
				xConsidered := newX % len(garden[0])
				if xConsidered < 0 {
					xConsidered = len(garden[0]) + xConsidered
				}
	
				if garden[yConsidered][xConsidered] != "#" {
					nextPositionSet[Coord{newY, newX}] = true
				}
			}
		}
		positions = make([]Coord, len(nextPositionSet))
		index := 0
		totalPossibleMoves = len(nextPositionSet)
		for coord := range nextPositionSet {
			positions[index] = coord
			index++
		}
		moveCount++
	
	}
	return totalPossibleMoves
}

func part_1(filename string) {
	garden, startPosition, err := parseFile(filename)
	if err != nil {
		panic(err)
	}
	numMoves := 64
	totalPossibleMoves := findMoves(startPosition, garden, numMoves)
	fmt.Println(totalPossibleMoves)
}

func findQuadraticCoefficients(u0, u1, u2 int) (int, int, int) {
	a := (u2+u0-2*u1)/2
    b := u1-u0-a
    c := u0
	return a, b, c
}

func part_2(filename string) {
	garden, startPosition, err := parseFile(filename)
	if err != nil {
		panic(err)
	}
	totalSteps := 26501365
	height := len(garden)
	mod := totalSteps % height
	u0 := findMoves(startPosition, garden, mod)
	u1 := findMoves(startPosition, garden, mod + height)
	u2 := findMoves(startPosition, garden, mod + 2*height)

	a, b, c := findQuadraticCoefficients(u0, u1, u2)
	f := func (x int) int {
		return a * x * x + b * x + c
	}
	
	fmt.Println(f((totalSteps / height)))

}

func main() {
	filename := "full.txt"
	// part_1(filename)
	part_2(filename)
}