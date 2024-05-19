package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return lines, nil
}

func deriveCalibrationValue(line string, numberMapping map[string]string) int {
	indexMap := make(map[int]string)
	maxIndex := 0
	minIndex := len(line)
	for k, v := range numberMapping {
		iteratedExample := line
		offset := 0
		index := strings.Index(iteratedExample, k)
		for index != -1 {
			relativeIndex := index + offset
			indexMap[relativeIndex] = v
			if relativeIndex > maxIndex {
				maxIndex = relativeIndex
			}
			if relativeIndex < minIndex {
				minIndex = relativeIndex
			}
			iteratedExample = iteratedExample[index+len(k):]
			offset = len(line) - len(iteratedExample)
			index = strings.Index(iteratedExample, k)
		}
	}
	value, _ := strconv.Atoi(indexMap[minIndex]+indexMap[maxIndex])
	return value
}

func calculateRunningTotal(numberMapping map[string]string) int {
	lines, _ := parseFile("file.txt")
	running_total := 0
	for _, line := range lines {
		running_total += deriveCalibrationValue(line, numberMapping)
	}
	return running_total
}


func part_1() {
	numberMap := make(map[string]string)
	for i := 1; i <= 9; i++ {
		val := fmt.Sprintf("%d", i)
		numberMap[val] = val
	}

	running_total := calculateRunningTotal(numberMap)
	fmt.Println("Part 1: ",running_total)
}

func part_2() {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	numberMapping := make(map[string]string)

	for i, word := range words {
		digit := fmt.Sprintf("%d", i+1)
		numberMapping[word] = digit
		numberMapping[digit] = digit
	}
	running_total := calculateRunningTotal(numberMapping)
	fmt.Println("Part 2: ",running_total)
}



func main() {
	part_1()
	part_2()
}