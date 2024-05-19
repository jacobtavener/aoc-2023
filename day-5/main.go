package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"sync"
	"math"
	// "time"
)

type Map struct {
	Source int
	Destination int
	Range int
}

type SeedRange struct {
	Start int
	Range int
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

func cleanseString(str string) string {
	return strings.TrimSpace(strings.ReplaceAll(str, "  ", " "))
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

func setUp1(almanacStr []string, seedsPtr *[]int, almanacMapPtr *map[string][]Map, sourceMapPtr *map[string]string) {
	mapping := ""
	for i, v := range almanacStr {
		if i == 0 {
			seedValues := cleanseString(v[6:])
			*seedsPtr, _ = convertStringListToInts(strings.Split(seedValues, " "))
		} else if v == "" {
			continue
		} else if strings.Contains(v, "map") {
			mapSplit := strings.Split(v, " ")
			mapping = mapSplit[0]
			(*almanacMapPtr)[mapping] = make([]Map, 0)
			mappingSplit := strings.Split(mapping, "-to-")
			(*sourceMapPtr)[mappingSplit[0]] = mappingSplit[1]
		} else {
			maps := strings.Split(cleanseString(v), " ")
			destination, _ := strconv.Atoi(maps[0])
			source, _ := strconv.Atoi(maps[1])
			rng, _ := strconv.Atoi(maps[2])

			(*almanacMapPtr)[mapping] = append((*almanacMapPtr)[mapping], Map{Source: source, Destination: destination, Range: rng})
		}
	}	
}

func setUp2(almanacStr []string, seedsPtr *[]SeedRange, almanacMapPtr *map[string][]Map, sourceMapPtr *map[string]string) {
	mapping := ""
	for i, v := range almanacStr {
		if i == 0 {
			seedValues := cleanseString(v[6:])
			seedsList, _ := convertStringListToInts(strings.Split(seedValues, " "))
			numPairs := len(seedsList) / 2
			for i := 0; i < numPairs; i++ {
				*seedsPtr = append(*seedsPtr, SeedRange{Start: seedsList[2*i], Range: seedsList[2*i+1]})
			}
		} else if v == "" {
			continue
		} else if strings.Contains(v, "map") {
			mapSplit := strings.Split(v, " ")
			mapping = mapSplit[0]
			(*almanacMapPtr)[mapping] = make([]Map, 0)
			mappingSplit := strings.Split(mapping, "-to-")
			(*sourceMapPtr)[mappingSplit[0]] = mappingSplit[1]
		} else {
			maps := strings.Split(cleanseString(v), " ")
			destination, _ := strconv.Atoi(maps[0])
			source, _ := strconv.Atoi(maps[1])
			rng, _ := strconv.Atoi(maps[2])

			(*almanacMapPtr)[mapping] = append((*almanacMapPtr)[mapping], Map{Source: source, Destination: destination, Range: rng})
		}
	}
}	


func part_1(filename string) {
	almanacStr, err := parseFile(filename)
	
	if err != nil {
		fmt.Println("Error parsing file:", err)
	}
	
	seeds:= make([]int, 0)
	almanacMap := make(map[string][]Map)
	sourceMap := make(map[string]string)

	
	setUp1(almanacStr, &seeds, &almanacMap, &sourceMap)

	minLocation := math.MaxInt
	for _, seed := range seeds {
		source := "seed"
		curValue := seed
		for {
			destination, exists := sourceMap[source]
			if !exists {
				break
			}
			for _, v := range almanacMap[source+"-to-"+destination] {
				if v.Source <= curValue && curValue < v.Source+v.Range {
					curValue += (v.Destination - v.Source)
					break
				}
			}
			source = destination
		}
		if curValue < minLocation {
			minLocation = curValue
		}
	}
	fmt.Println("Part 1: ", minLocation)
}

func processSeedRange(seedRange SeedRange, almanacMap map[string][]Map, sourceMap map[string]string, minLocationChannel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	minLocation := math.MaxInt
	for seed := seedRange.Start; seed < seedRange.Start + seedRange.Range; seed++ {
		source := "seed"
		curValue := seed
		for {
			destination, exists := sourceMap[source]
			if !exists {
				break
			}
			for _, v := range almanacMap[source+"-to-"+destination] {
				if v.Source <= curValue && curValue < v.Source+v.Range {
					curValue += (v.Destination - v.Source)
					break
				}
			}
			source = destination
		}
		if curValue < minLocation {
			minLocation = curValue
		}
	}
	minLocationChannel <- minLocation
}

var wg sync.WaitGroup

func part_2(filename string) {
	almanacStr, err := parseFile(filename)
	
	if err != nil {
		fmt.Println("Error parsing file:", err)
	}
	
	seeds:= make([]SeedRange, 0)
	almanacMap := make(map[string][]Map)
	sourceMap := make(map[string]string)

	setUp2(almanacStr, &seeds, &almanacMap, &sourceMap)

	numProcesses := len(seeds)
	minLocationChannel := make(chan int, numProcesses)
	wg.Add(numProcesses)

	for _, seed := range seeds {
		go processSeedRange(seed, almanacMap, sourceMap, minLocationChannel, &wg)
		
	}

	go func() {
		wg.Wait()
		close(minLocationChannel)
	}()

	minLocation := math.MaxInt
	for minLocationVal := range minLocationChannel {
		if minLocationVal < minLocation {
			minLocation = minLocationVal
		}
	}
	fmt.Println("Part 2: ", minLocation)
}

func main() {
	part_1("full.txt")
	// Takes a few mins but works eventually...
	part_2("full.txt")
}