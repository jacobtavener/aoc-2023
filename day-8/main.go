package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func extractValues(s string) []string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "(")
	s = strings.TrimSuffix(s, ")")
	values := strings.Split(s, ",")
	for i, v := range values {
		values[i] = strings.TrimSpace(v)
	}
	return values
}


func convertStringListToInts(stringList []string) ([]int, error) {
	var intList []int

	for _, str := range stringList {
		// Convert each string to an integer
		i, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intList = append(intList, i)
	}

	return intList, nil
}

func parseFile1(filePath string) ([]int, map[string][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, err
	}
	defer file.Close()
	var sequence []int
	mapping := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if sequence == nil {
			sequenceRaw := strings.ReplaceAll(strings.ReplaceAll(line, "L", "0"), "R", "1")
			sequenceStr := strings.Split(sequenceRaw, "")
			sequence, err = convertStringListToInts(sequenceStr)
			if err != nil {
				fmt.Println("Error converting string list to ints:", err)
				return nil, nil, err
			}
			continue
		}
		if line == "" {	
			continue
		}
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}
		key := strings.TrimSpace(parts[0])
		values := extractValues(parts[1])
		mapping[key] = values
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil, err
	}

	return sequence, mapping, nil
}

func parseFile2(filePath string) ([]int, []string, map[string][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, nil, err
	}
	defer file.Close()
	var sequence []int
	startingNodes := make([]string, 0)
	mapping := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if sequence == nil {
			sequenceRaw := strings.ReplaceAll(strings.ReplaceAll(line, "L", "0"), "R", "1")
			sequenceStr := strings.Split(sequenceRaw, "")
			sequence, err = convertStringListToInts(sequenceStr)
			if err != nil {
				fmt.Println("Error converting string list to ints:", err)
				return nil, nil, nil, err
			}
			continue
		}
		if line == "" {	
			continue
		}
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}
		key := strings.TrimSpace(parts[0])
		values := extractValues(parts[1])
		mapping[key] = values

		if strings.HasSuffix(key, "A") {
			startingNodes = append(startingNodes, key)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil, nil, err
	}

	return sequence, startingNodes, mapping, nil
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(values ...int) int {
	result := 1
	for _, value := range values {
		result = (result * value) / gcd(result, value)
	}
	return result
}


func part1(filename string) {
	sequence, mapping, err := parseFile1(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	startingDestination := "AAA"
	finalDestination := "ZZZ"
	sequencePosition := 0
	sequenceLength := len(sequence)
	currentPosition := startingDestination
	numSteps := 0
	for {
		if currentPosition == finalDestination {
			break
		}
		currentPosition = mapping[currentPosition][sequence[sequencePosition]]
		numSteps++
		sequencePosition = (sequencePosition + 1) % sequenceLength
	}
	fmt.Println("Part 1: ", numSteps)

}


func part2(filename string) {
	sequence, startingNodes, mapping, err := parseFile2(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	loopList := make([]int, 0)
	sequenceLength := len(sequence)
	for _, startingNode := range startingNodes {
		currentPosition := startingNode
		numSteps := 0
		sequencePosition := 0
		reachedDestination := false
		loopDetection := make(map[string]bool)
		loopDetected := false
		loopIterations := 0
		for {
			moveRepr := currentPosition + strconv.Itoa(sequencePosition)
			_, foundLoop := loopDetection[moveRepr]
			if foundLoop {
				if loopDetected {
					break
				}
				loopDetected = true
				loopDetection = make(map[string]bool)
			}
			if strings.HasSuffix(currentPosition, "Z") {
				reachedDestination = true
			}
			loopDetection[moveRepr] = true
			currentPosition = mapping[currentPosition][sequence[sequencePosition]]
			if !reachedDestination {
				numSteps++
			}
			if loopDetected {
				loopIterations++
			}
			sequencePosition = (sequencePosition + 1) % sequenceLength
		}
		loopList = append(loopList, numSteps)
	}
	fmt.Println("Part 2:", lcm(loopList...))
}

func main() {
	filename := "full.txt"
	part1(filename)
	part2(filename)
}