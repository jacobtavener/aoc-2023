package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Dish struct {
	Rows    []string
	Columns []string
}

func parseFile(filePath string) ([]Dish, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []Dish

	rows := make([]string, 0)
	columns := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			lines = append(lines, Dish{Rows: rows, Columns: columns})
			rows = make([]string, 0)
			columns = make([]string, 0)
		} else {
			rows = append(rows, line)
			for i, c := range line {
				if len(columns) <= i {
					columns = append(columns, string(c))
				} else {
					columns[i] += string(c)
				}
			}
		}
	}
	lines = append(lines, Dish{Rows: rows, Columns: columns})

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return lines, nil
}

func transposeStringSlice(slice []string) []string {
	newSlice := make([]string, len(slice[0]))
	for i := range slice[0] {
		for j := range slice {
			newSlice[i] += string(slice[j][i])
		}
	}
	return newSlice
}

func rowsToColumns(rows []string) []string {
	columns := make([]string, 0)
	for _, row := range rows {
		for j, c := range row {
			if len(columns) <= j {
				columns = append(columns, string(c))
			} else {
				columns[j] += string(c)
			}
		}
	}
	return columns
}

func findLoad(columns []string) int {
	totalScore := 0
	maxScore := len(columns[0])
	for _, col := range columns {
		for i, c := range col {
			if c == 'O' {
				totalScore += maxScore - i
			}
		}
	}
	return totalScore
}

func reverse(str string) (result string) { 
    for _, v := range str { 
        result = string(v) + result 
    } 
    return
} 

var seenStrings = make(map[string]bool, 0)
var cycleCount = 0
var initialCount = 0

func runCycle(columns []string) ([]string, bool) {
	colLen := len(columns[0])
	newColumns := make([]string, 0)
	for _, col := range columns {
		score := colLen
		newPositions := make(map[int]bool, 0)
		for i, c := range col {
			if c == 'O' {
				newPositions[colLen-score] = true
				score--
			} else if c == '#' {
				score = colLen - i - 1
			} else {
				continue
			}
		}
		newColumn := ""
		for i, c := range col {
			if _, ok := newPositions[i]; ok {
				newColumn += "O"
			} else if c == '#' {
				newColumn += "#"
			} else {
				newColumn += "."
			}
		}
		newColumns = append(newColumns, newColumn)
	}
	rows := transposeStringSlice(newColumns)
	rowLen := len(rows[0])
	newRows := make([]string, 0)
	for _, row := range rows {
		score := rowLen
		newPositions := make(map[int]bool, 0)
		for i, c := range row {
			if c == 'O' {
				newPositions[rowLen-score] = true
				score--
			} else if c == '#' {
				score = rowLen - i - 1
			} else {
				continue
			}
		}
		newRow := ""
		for i, c := range row {
			if _, ok := newPositions[i]; ok {
				newRow += "O"
			} else if c == '#' {
				newRow += "#"
			} else {
				newRow += "."
			}
		}
		newRows = append(newRows, newRow)
	}
	columns = rowsToColumns(newRows)
	newColumns = make([]string, 0)
	for _, col := range columns {
		score := colLen
		newPositions := make(map[int]bool, 0)
		col = reverse(col)
		for i, c := range col {
			if c == 'O' {
				newPositions[colLen-score] = true
				score--
			} else if c == '#' {
				score = colLen - i - 1
			} else {
				continue
			}
		}
		newColumn := ""
		for i, c := range col {
			if _, ok := newPositions[i]; ok {
				newColumn += "O"
			} else if c == '#' {
				newColumn += "#"
			} else {
				newColumn += "."
			}
		}
		newColumns = append(newColumns, reverse(newColumn))
	}
	rows = transposeStringSlice(newColumns)
	newRows = make([]string, 0)
	for _, row := range rows {
		score := rowLen
		newPositions := make(map[int]bool, 0)
		row = reverse(row)
		for i, c := range row {
			if c == 'O' {
				newPositions[rowLen-score] = true
				score--
			} else if c == '#' {
				score = rowLen - i - 1
			} else {
				continue
			}
		}
		newRow := ""
		for i, c := range row {
			if _, ok := newPositions[i]; ok {
				newRow += "O"
			} else if c == '#' {
				newRow += "#"
			} else {
				newRow += "."
			}
		}
		newRows = append(newRows, reverse(newRow))
	}
	columns = rowsToColumns(newRows)
	finalColumns := []string{}
	finalColumns = append(finalColumns, columns...)

	colStr := strings.Join(finalColumns[:], "+")
	breakCycle := false
	if val, ok := seenStrings[colStr]; ok {
		if val {
			// fmt.Println("Cycle ended")
			breakCycle = true
		} else {
			// fmt.Println("Start cycle")
			seenStrings[colStr] = true
			cycleCount++
		}
	} else {
		// fmt.Println("New string")
		seenStrings[colStr] = false
		initialCount++
	}
	return finalColumns, breakCycle
}

func main() {
	filename := "full.txt"

	dishes, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}

	columns := dishes[0].Columns
	for i := 0; i < 1000000000; i++ {
		newColumns, breakCycle := runCycle(columns)
		columns = newColumns
		if breakCycle {
			break
		}

	}

	fmt.Println(initialCount, cycleCount)

	requiredCycles := ((1000000000 - initialCount) % cycleCount) + initialCount

	restartColumns := dishes[0].Columns
	for i := 0; i < requiredCycles; i++ {
		newColumns, _ := runCycle(restartColumns)
		restartColumns = newColumns
	}

	load := findLoad(restartColumns)
	fmt.Println(load)

}
