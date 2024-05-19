package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	X, Y int
}

func parseFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := make([][]string, 0)
	for scanner.Scan() {
		rowStr := scanner.Text()
		row := make([]string, 0)
		for _, str := range rowStr {
			row = append(row, string(str))
		}
		grid = append(grid, row)
	}
	return grid, nil
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && x < cols && y >= 0 && y < rows
}

func longestPath(grid [][]string, start, end Coord, step int, visited [][]bool, useSlopes bool) int {
	current := start
	rows := len(grid)
	cols := len(grid[0])
	if current == end {
		return step
	}

	currentValue := grid[current.Y][current.X]

	directions := []Coord{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 1}}
	if useSlopes && currentValue != "#" && currentValue != "." {
		switch currentValue {
		case ">":
			directions = []Coord{{X: 1, Y: 0}}
		case "<":
			directions = []Coord{{X: -1, Y: 0}}
		case "^":
			directions = []Coord{{X: 0, Y: -1}}
		case "v":
			directions = []Coord{{X: 0, Y: 1}}
		}
	}

	maxStep := 0 

	for _, dir := range directions {
		next := Coord{current.X + dir.X, current.Y + dir.Y}
		// if next.X == 1 && next.Y == 0 {
		// 	continue
		// }
		if !isValid(next.X, next.Y, rows, cols) {
			continue
		}
		if grid[next.Y][next.X] != "#" {
			if found := visited[next.Y][next.X]; !found {
				visited[next.Y][next.X] = true
				maxStep = max(maxStep, longestPath(grid, next, end, step + 1, visited, useSlopes))
				visited[next.Y][next.X] = false
			}
		}
	}
	return maxStep
}

func part_1(grid [][]string) {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	step := 0
	start := Coord{Y: 0, X: 1}
	end := Coord{Y: rows - 1, X: cols - 2}
	maxSteps:=longestPath(grid, start, end, step, visited, true)
	fmt.Println(maxSteps)
}

func part_2(grid [][]string) {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	step := 0
	start := Coord{Y: 0, X: 1}
	end := Coord{Y: rows - 1, X: cols - 2}
	maxSteps:=longestPath(grid, start, end, step, visited, false)
	fmt.Println(maxSteps)
}


func main() {
	filename := "full.txt"
	grid, err := parseFile(filename)

	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}
	part_1(grid)
	part_2(grid)
}
