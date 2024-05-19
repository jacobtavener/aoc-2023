package main

import (
	"os"
	"bufio"
	"fmt"
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

	var codes []string

	for scanner.Scan() {
		line := scanner.Text()
		codeList := strings.Split(line, ",")
		codes = append(codes, codeList...)
	}
	return codes, nil
}

func calculateHash(str string) int {
	hash := 0
	for _, c := range str {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return hash
}

func part_1(filename string) {
	codes, err := parseFile(filename)
	if err != nil {	
		fmt.Println("Error parsing file:", err)
		return
	}
	total := 0
	for _, code := range codes {
		total += calculateHash(code)
	}
	fmt.Println("Total: ", total)
}

func main() {
	filename := "full.txt"
	codes, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}
	focalMap := make(map[string]int)
	boxMap := make(map[int][]string)
	for i:=0; i<256; i++ {
		boxMap[i] = make([]string, 0)
	}
	for _, code := range codes {
		if code[len(code)-1] == '-' {
			operation := code[:len(code)-1]
			boxNumber := calculateHash(operation)
			if _, exists := focalMap[operation]; exists {
				delete(focalMap, operation)
				box := boxMap[boxNumber]
				for i, c := range box {
					if c == operation {
						box = append(box[:i], box[i+1:]...)

					}
				}
				boxMap[boxNumber] = box
				
			}

		} else {
			codeSplit := strings.Split(code, "=")
			operation, focalStr := codeSplit[0], codeSplit[1]
			boxNumber := calculateHash(operation)
			focal, _ := strconv.Atoi(focalStr)
			if _, exists := focalMap[operation]; !exists {
				boxMap[boxNumber] = append(boxMap[boxNumber], operation)
			}
			focalMap[operation] = focal
		}
	}
	total := 0
	for i, box := range boxMap {
		if len(box) == 0 {
			continue
		} else {
			for j, operation := range box {
				total += (i+1) * (j+1) * focalMap[operation]
			}
		}
	}
	fmt.Println("Total: ", total)
	
}