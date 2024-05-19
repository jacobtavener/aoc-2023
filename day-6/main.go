package main

import (
	"os"
	"bufio"
	"fmt"
	"errors"
	"math"
	"strings"
	"strconv"
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

func cleanseString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func convertLineToFloat64List(value string) []float64 {
	stringList := strings.Split(cleanseString(value), " ")
	var floatList []float64

	for _, str := range stringList {
		i, _ := strconv.ParseFloat(str, 64)
		floatList = append(floatList, i)
	}
	return floatList
}

func concatToFloat64(value string) float64 {
	cleansedValue := cleanseString(value)
	f64Val, _ := strconv.ParseFloat(strings.Join(strings.Split(cleansedValue, " "), ""), 64)
	return f64Val
}

func quadraticFormulaIntegerPossibilities(b float64, c float64) (int, error) {
	var descriminant float64 = float64((b * b) - (4 * c))
	if descriminant < 0 {
		return 0, errors.New("no real solutions")
	}

	x1 := (-b + math.Sqrt(descriminant)) / 2
	x2 := (-b - math.Sqrt(descriminant)) / 2

	possibleLowerBound := math.Min(x1, x2)

	lowerBoundCeil := math.Ceil(possibleLowerBound)
	var lowerBound int
	if possibleLowerBound == lowerBoundCeil {
		lowerBound = int(lowerBoundCeil) + 1
	} else {
		lowerBound = int(lowerBoundCeil)
	}

	possibleUpperBound := math.Max(x1, x2)
	upperBoundFloor := math.Floor(possibleUpperBound)
	var upperBound int
	if possibleUpperBound == upperBoundFloor {
		upperBound = int(math.Floor(possibleUpperBound)) - 1
	} else {
		upperBound = int(math.Floor(possibleUpperBound))
	}
	possibleIntSolutions := upperBound - lowerBound + 1

	return possibleIntSolutions, nil

}



func part_1(filename string) {
	lines, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	times := convertLineToFloat64List(lines[0][5:])
	distances := convertLineToFloat64List(lines[1][10:])

	var total int
	for i, v := range times {
		numPossibleRaceStrategies, _ := quadraticFormulaIntegerPossibilities(-v, distances[i])
		if i == 0 {
			total = numPossibleRaceStrategies
		} else {
			total *= numPossibleRaceStrategies
		}
	}
	fmt.Println("Part 1: ", total)
}



func part_2(filename string) {
	lines, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	time := concatToFloat64(lines[0][5:])
	distance := concatToFloat64(lines[1][10:])

	numPossibleRaceStrategies, err := quadraticFormulaIntegerPossibilities(-time, distance)
	if err != nil {
		fmt.Println("Error calculating quadratic formula:", err)
		return
	}
	fmt.Println("Part 2: ", numPossibleRaceStrategies)

}

func main() {
	part_1("full.txt")
	part_2("full.txt")
}