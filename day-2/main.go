package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var colourMap = map[string]int{
	"red": 12,
	"green": 13,
	"blue": 14,
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

func part_1(filename string) {
	games, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	possibleGames := 0
	for _, game := range games {
		gameSplit := strings.Split(game, ": ")

		gameIdFullStr, roundStr := gameSplit[0], gameSplit[1]
		
		rounds := strings.Split(roundStr, "; ")

		validRounds := 0
		for _, round := range rounds {
			balls := strings.Split(round, ", ")
			validBalls := 0
			for _, ball := range balls {
				ballSplit := strings.Split(ball, " ")
				countStr, colour := ballSplit[0], ballSplit[1]
				if count, _ := strconv.Atoi(countStr); count <= colourMap[colour] {
					validBalls++
				}
			}
			if len(balls) == validBalls {
				validRounds++
			}
		}
		
		if len(rounds) == validRounds {
			gameId, _ := strconv.Atoi(gameIdFullStr[5:])
			possibleGames += gameId
		}

	}

	fmt.Println("Part 1:", possibleGames)
}

func part_2(filename string) {
	games, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	total := 0
	for _, game := range games {
		maxColourMap := map[string]int{
			"red": 0,
			"green": 0,
			"blue": 0,
		}
		gameSplit := strings.Split(game, ": ")
		roundStr := gameSplit[1]
		rounds := strings.Split(roundStr, "; ")
		for _, round := range rounds {
			balls := strings.Split(round, ", ")
			for _, ball := range balls {
				ballSplit := strings.Split(ball, " ")
				countStr, colour := ballSplit[0], ballSplit[1]
				count, _ := strconv.Atoi(countStr)
				if count > maxColourMap[colour] {
					maxColourMap[colour] = count
				}

			}
		}
		total += maxColourMap["red"] * maxColourMap["green"] * maxColourMap["blue"]
	}

	fmt.Println("Part 2:", total)
}


func main() {
	filename := "full.txt"
	part_1(filename)
	part_2(filename)
}