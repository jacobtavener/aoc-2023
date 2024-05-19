package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"math"
)


type Result struct {
	Score int
	Copies int
}

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

func findIntersection(a, b []string) []string {
	intersection := make([]string, 0)

	set := make(map[string]bool)
	for _, v := range a {
		set[v] = true
	}

	for _, v := range b {
		if _, ok := set[v]; ok {
			intersection = append(intersection, v)
		}
	}
	return intersection
}

func calculateWinners(game string) []string {
	gameSplit := strings.Split(game[8:], " | ")
	winningNumbersStr, numbersStr := strings.TrimSpace(strings.ReplaceAll(gameSplit[0], "  ", " ")), strings.TrimSpace(strings.ReplaceAll(gameSplit[1], "  ", " "))
	winningNumbers := strings.Split(winningNumbersStr, " ")
	numbers := strings.Split(numbersStr, " ")
	return findIntersection(winningNumbers, numbers)
}

func calculateScore(winners []string) int {
	if len(winners) == 0 {
		return 0
	}
	return int(math.Pow(2, float64(len(winners)) - 1))
}

func calculateGameScore(game string) int {
	winners := calculateWinners(game)
	return calculateScore(winners)
}

func part_1(filename string) {
	games, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}
	score := 0
	for _, game := range games {
		gameScore := calculateGameScore(game)
		score += gameScore
	}
	fmt.Println("Part 1: ", score)
}

func part_2(filename string) {
	games, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}
	score := 0
	scratchCards := make(map[int]int)
	for i, game := range games {
		gameWinners := calculateWinners(game)
		numWinners := len(gameWinners)
		numCopies, exists := scratchCards[i]
		if !exists {
			numCopies = 0
		}
		instances := 1 + numCopies

		for j:=i+1; j<i+1+numWinners; j++ {
			scratchCards[j] += instances
		}
		score += instances
	}
	fmt.Println("Part 2: ", score)
}

func main() {
	part_1("full.txt")
	part_2("full.txt")
}