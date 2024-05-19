/*
This solution is convoluted and frankly a bit mental.
I decided to convert the cards to a base 15 number system to make it easier to compare hands
but this wasn't necessary and I could have just compared the hands directly.

It gets the job done but it's not the most elegant solution.
*/


package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"sort"
	"math"
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

var cardMap = map[string]string{
	"T": "10",
	"J": "1",
	"Q": "12",
	"K": "13",
	"A": "14",
}

func convertStringListToNumRep(stringList []string) (int, error) {
	binInt := 0
	length := len(stringList)
	for i:=0; i<length; i++ {
		cardStr:= stringList[i]
		if val, ok := cardMap[cardStr]; ok {
			cardStr = val
		}
		card, err := strconv.Atoi(cardStr)
		if err != nil {
			return 0, err
		}
		// use base 15 to convert the cards to a single integer
		binInt += card * int(math.Pow(15, float64(length-i)))
	}
	return binInt, nil
}

func classifyHand(hand []string) string {
	initialHandBreakdown := make(map[string]int)
	handBreakdown := make(map[string]int)
	translatedHand := make([]string, len(hand))
    copy(translatedHand, hand[:])
	for _, card := range hand {
		// if _, ok := initialHandBreakdown[card]; ok {
		// 	initialHandBreakdown[card]++
		// } else {
		// 	initialHandBreakdown[card] = 1
		// }
		initialHandBreakdown[card]++

	}
	maxVal := 0
	var maxLabel string
	for label, val := range initialHandBreakdown {
		if val > maxVal && label != "J" {
			maxVal = val
			maxLabel = label
		}
	}
	if maxLabel != "" {
		for i, card := range hand {
			if card == "J" {
				translatedHand[i] = maxLabel
			}
		}
	}

	for _, card := range translatedHand {
		// if _, ok := handBreakdown[card]; ok {
		// 	handBreakdown[card]++
		// } else {
		// 	handBreakdown[card] = 1
		// }
		handBreakdown[card]++
	}

	if len(handBreakdown) == 1 {
		return "five of a kind"
	}
	if len(handBreakdown) == 2 {
		maxVal := 0
		for _, val := range handBreakdown {
			if val > maxVal {
				maxVal = val
			}
		}
		if maxVal == 4 {
			return "four of a kind"
		} else {
			return "full house"
		}
	}
	if len(handBreakdown) == 3 {
		maxVal := 0
		for _, val := range handBreakdown {
			if val > maxVal {
				maxVal = val
			}
		}
		if maxVal == 3 {
			return "three of a kind"
		} else {
			return "two pair"
		}
	}
	if len(handBreakdown) == 4 {
		return "one pair"
	}
	return "high card"

}

func main() {
	filename := "full.txt"
	lines, err := parseFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	classifiedHands := map[string][]int{
		"five of a kind": make([]int, 0),
		"four of a kind": make([]int, 0),
		"full house": make([]int, 0),
		"three of a kind": make([]int, 0),
		"two pair": make([]int, 0),
		"one pair": make([]int, 0),
		"high card": make([]int, 0),
	}
	handsOrdering := []string{
		"five of a kind",
		"four of a kind",
		"full house",
		"three of a kind",
		"two pair",
		"one pair",
		"high card",
	}
	handBidMap := make(map[int]int)
	handBinStrMap := make(map[int][]string)
	for _, line := range lines {
		lineStr := cleanseString(line)
		lineSplit := strings.Split(lineStr, " ")
		handstr, bidStr := lineSplit[0], lineSplit[1]
		hand := strings.Split(handstr, "")
		handClassification := classifyHand(hand)
		digitHand, err := convertStringListToNumRep(hand)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		handBinStrMap[digitHand] = hand
		handBidMap[digitHand] = bid
		classifiedHands[handClassification] = append(classifiedHands[handClassification], digitHand)
	}
	multiplier := len(lines)
	totalScore := 0
	for _, c := range handsOrdering {
		if len(classifiedHands[c]) > 0 {
			hands := classifiedHands[c]
			sort.Sort(sort.Reverse(sort.IntSlice(hands)))
			for _, hand := range hands {
				totalScore += handBidMap[hand] * multiplier
				multiplier--
			}
		}	
	}
	fmt.Println(totalScore)
}