package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

type Record struct {
	Condition []string
	Group []int
	Total int
	Unknown []int
	Known []int
}

func convertStringListToInts(stringList []string) ([]int, int, error) {
	var intList []int
	total := 0
	for _, str := range stringList {
		// Convert each string to an integer
		i, err := strconv.Atoi(str)
		if err != nil {
			return nil, 0, err
		}
		total += i
		intList = append(intList, i)
	}

	return intList, total, nil
}

func findOccurrences(condition []string, character string) []int {
	indices := []int{}
	for i, c := range condition {
		if c == character {
			indices = append(indices, i)
		}
	}
	return indices
}


func parseFile(filePath string, multiplier int) ([]Record, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var records []Record

	for scanner.Scan() {
		lineStr := scanner.Text()
		lineSplitBase := strings.Split(lineStr, " ")
		repeatedConditions := make([]string, multiplier)
		repeatedGroups := make([]string, multiplier)
		for i := 0; i < multiplier; i++ {
			repeatedConditions[i] = lineSplitBase[0]
			repeatedGroups[i] = lineSplitBase[1]
		}
		lineSplit := []string{strings.Join(repeatedConditions, "?"), strings.Join(repeatedGroups, ",")}
		condition := strings.Split(lineSplit[0], "")
		unknown := findOccurrences(condition, "?")
		known := findOccurrences(condition, "#")
		groupStr := strings.Split(lineSplit[1], ",")
		group, total, err := convertStringListToInts(groupStr)
		if err != nil {
			fmt.Println("Error converting string list to ints:", err)
			return nil, err
		}
		record := Record{
			Condition: condition,
			Group: group,
			Total: total,
			Unknown: unknown,
			Known: known,
		}
		records = append(records, record)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return records, nil
}

func isValidRecord(record Record) bool {
	condition := record.Condition
	group := record.Group
	curGroupIndex := 0
	maxGroupIndex := len(group) - 1
	curDamagedCount := 0
	isValid := true
	fullyFormed := true
	for _, c := range condition {
		if c == "?" {
			fullyFormed = false
			break
		}
		if c == "#" {
			curDamagedCount++
			if curGroupIndex > maxGroupIndex {
				isValid = false
				break
			}
			if curDamagedCount > group[curGroupIndex] {
				isValid = false
				break
			}
		}
		if c == "." {
			if curDamagedCount > 0 {
				curDamagedCount = 0
				curGroupIndex++
			}
		}
	}
	if fullyFormed && curGroupIndex < len(group) - 1 {
		isValid = false
	}
	return isValid
}

func copyRecord(record Record) Record {
    // Create a new copy of the Record
    conditionCopy := make([]string, len(record.Condition))
    copy(conditionCopy, record.Condition)

    groupCopy := make([]int, len(record.Group))
    copy(groupCopy, record.Group)

    unknownCopy := make([]int, len(record.Unknown))
    copy(unknownCopy, record.Unknown)

    knownCopy := make([]int, len(record.Known))
    copy(knownCopy, record.Known)

    return Record{
        Condition: conditionCopy,
        Group:     groupCopy,
        Total:     record.Total,
        Unknown:   unknownCopy,
        Known:     knownCopy,
    }
}


func generatePermutations(record Record, validPermutationCount *int, operationalCount, damagedCount, index int) {
	if operationalCount == 0 && damagedCount == 0 {
		if isValidRecord(record) {
			*validPermutationCount++
		}
		return
	}
	if operationalCount > 0 {
		recordCopy := copyRecord(record)
		recordCopy.Condition[recordCopy.Unknown[index]] = "."
		if isValidRecord(recordCopy) {
			generatePermutations(recordCopy, validPermutationCount, operationalCount - 1, damagedCount, index + 1)
		}
	}
	if damagedCount > 0 {
		recordCopy := copyRecord(record)
		recordCopy.Condition[recordCopy.Unknown[index]] = "#"
		if isValidRecord(recordCopy) {
			generatePermutations(recordCopy, validPermutationCount, operationalCount, damagedCount - 1, index + 1)
		}
	}
}



func findValidPermutationCount(record Record) int {
	damagedCount := record.Total - len(record.Known)
	operationalCount := len(record.Unknown) - damagedCount
	validPermutationCount := 0
	generatePermutations(record, &validPermutationCount, operationalCount, damagedCount, 0)
	return validPermutationCount
}

func main() {
	filename := "full.txt"
	records, err := parseFile(filename, 5)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}
	permuationCount := 0
	for _, record := range records {
		rcd := record
		validPermuationCount := findValidPermutationCount(rcd)

		fmt.Println("Valid permutation count:", validPermuationCount)
		permuationCount += validPermuationCount
	}
	fmt.Println("Permutation count:", permuationCount)
}