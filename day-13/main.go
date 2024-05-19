package main

import (
	"os"
	"bufio"
	"fmt"
)

type Note struct {
	Rows []string
	Columns []string
}

func parseFile(filePath string) ([]Note, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []Note

	rows := make([]string, 0)
	columns := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			lines = append(lines, Note{Rows: rows, Columns: columns})
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
	lines = append(lines, Note{Rows: rows, Columns: columns})

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return lines, nil
}

func compareStrings(str1, str2 string) (bool, bool) {
	if len(str1) != len(str2) {
		return false, false
	}
	smudgeUsed := false
	for i, c := range str1 {
		if string(c) != string(str2[i]) {
			if smudgeUsed {
				return false, false
			} else {
				smudgeUsed = true
			}
		}
	}
	return true, smudgeUsed
}

func main() {
	filename := "full.txt"

	notes, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}
	total := 0
	for _, note := range notes {
		rows := note.Rows
		numRows := len(rows)
		columns := note.Columns
		numColumns := len(columns)
		rowTotal := 0
		colTotal := 0
		smudgeFound := false

		for i := 0; i < numRows - 1; i++ {
			potentialMirror, _ := compareStrings(rows[i], rows[i + 1])
			if potentialMirror {
				legitMirror := true
				span := 0
				if i < numRows / 2 {
					span = i
				} else {
					span = numRows - i - 2
				}
				for j := 0; j <= span; j++ {
					equalStrings, smudgeUsed := compareStrings(rows[i-j], rows[i+j+1])
					if !equalStrings || (smudgeFound && smudgeUsed) {
						legitMirror = false
						smudgeFound = false
						break
					}
					if smudgeUsed {
						smudgeFound = true
					}
				}
				if legitMirror && smudgeFound {
					rowTotal = i+1 
					break
				}
			}
		}
		if rowTotal == 0 {
			for i := 0; i < numColumns - 1; i++ {
				potentialMirror, _ := compareStrings(columns[i], columns[i + 1])
				if potentialMirror {
					legitMirror := true
					span := 0
					if i < numColumns / 2 {
						span = i
					} else {
						span = numColumns - i - 2
					}
					for j := 0; j <= span; j++ {
						equalStrings, smudgeUsed := compareStrings(columns[i-j], columns[i+j+1])
						if !equalStrings || (smudgeFound && smudgeUsed) {
							legitMirror = false
							smudgeFound = false
							break
						}
						if smudgeUsed {
							smudgeFound = true
						}
					}
					if legitMirror && smudgeFound {
						colTotal = i+1 
						break
					}
				}
			}
		}
		total +=colTotal + (rowTotal * 100)

	}
	fmt.Println("Total:", total)
}