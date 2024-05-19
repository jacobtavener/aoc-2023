package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numbersList := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		numStr := strings.Fields(line)


		var numbers []int
		for _, str := range numStr {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil, err
			}
			numbers = append(numbers, num)
		}
		numbersList = append(numbersList, numbers)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return numbersList, nil
}

func findDiffSequence(line []int) int {
	diffSeq := make([]int, 0)
	endSeq := true
	for i := 0; i < len(line)-1; i++ {
		diff := line[i+1] - line[i]
		diffSeq = append(diffSeq, diff)
		if endSeq && diff != 0 {
			endSeq = false
		}
	}
	extra := 0
	if !endSeq {
		extra = findDiffSequence(diffSeq)
	}
	return diffSeq[len(diffSeq)-1] + extra
}

func findDiffSequenceHistory(line []int) int {
	diffSeq := make([]int, 0)
	endSeq := true
	for i := 0; i < len(line)-1; i++ {
		diff := line[i+1] - line[i]
		diffSeq = append(diffSeq, diff)
		if endSeq && diff != 0 {
			endSeq = false
		}
	}
	extra := 0
	if !endSeq {
		extra = findDiffSequenceHistory(diffSeq)
	}
	return diffSeq[0] - extra
}

func findNextVal(line []int) int {
	extra := findDiffSequence(line)
	return line[len(line)-1] + extra
}

func findNextValHistory(line []int) int {
	extra := findDiffSequenceHistory(line)
	return line[0] - extra
}

func part_1(filename string) {
	lines, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
	}

	total := 0
	for _, line := range lines {
		total += findNextVal(line)
	}
	fmt.Println("Total:", total)
}

func part_2(filename string) {
	lines, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
	}

	total := 0
	for _, line := range lines {
		total += findNextValHistory(line)
	}
	fmt.Println("Total:", total)
}

func main() {
	filename := "full.txt"
	part_1(filename)
	part_2(filename)


}