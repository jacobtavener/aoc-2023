package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Coord struct {
	X int
	Y int
}

func parseFile(filePath string) ([][]string, map[int]bool, map[int]bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var grid [][]string

	rowCoverage := make(map[int]bool)
	colCoverage := make(map[int]bool)
	rowIndex := 0
	for scanner.Scan() {
		rowStr := scanner.Text()
		row := strings.Split(rowStr, "")
		for i, element := range row {
			if element == "#" {
				rowCoverage[rowIndex] = true
				colCoverage[i] = true
			}
		}
		grid = append(grid, row)
		rowIndex++
	}
	return grid, rowCoverage, colCoverage, nil
}

func findShortestPath(p1 Coord, p2 Coord, rowCoverage map[int]bool, colCoverage map[int]bool, multiplier int) int {
	extraRows := 0
	extraCols := 0
	maxX := int(math.Max(float64(p1.X), float64(p2.X)))
	maxY := int(math.Max(float64(p1.Y), float64(p2.Y)))
	minX := int(math.Min(float64(p1.X), float64(p2.X)))
	minY := int(math.Min(float64(p1.Y), float64(p2.Y)))

	for i := minX; i <= maxX; i++ {
		if !colCoverage[i] {
			extraCols++
		}
	}

	for i := minY; i <= maxY; i++ {
		if !rowCoverage[i] {
			extraRows++
		}
	}
	return int(math.Abs(float64(maxX-minX+(extraCols*multiplier))) + math.Abs(float64(extraRows*multiplier+(maxY-minY))))
}
 
func main() {
	filename := "full.txt"

	grid, rowCoverage, colCoverage, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}

	galaxyMap := make([]Coord, 0)
	for i, row := range grid {
		for j, element := range row {
			if element == "#" {
				galaxyMap = append(galaxyMap, Coord{X: j, Y: i})
			}
		}
	}

	// part 1
	total := 0
	for i := len(galaxyMap)-1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			total += findShortestPath(galaxyMap[i], galaxyMap[j], rowCoverage, colCoverage, 1)
		}
	}
	fmt.Println("Total shortest path:", total)

	// part 2
	total = 0
	for i := len(galaxyMap)-1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			total += findShortestPath(galaxyMap[i], galaxyMap[j], rowCoverage, colCoverage, (1000000-1))
		}
	}
	fmt.Println("Total shortest path:", total)






}