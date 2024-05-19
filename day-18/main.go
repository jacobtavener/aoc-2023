package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)


type Direction string

const (
	Up Direction = "U"
	Down Direction = "D"
	Left Direction = "L"
	Right Direction = "R"
)

type Dig struct {
	Direction Direction
	Distance int
	Color  string
}

type Coord struct {
	X int
	Y int
}

func parseFile(filename string) ([]Dig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var digPlan []Dig
	for scanner.Scan() {
		lineStr := scanner.Text()
		var dig Dig
		lineSplit := strings.Split(lineStr, " ")
		dig.Direction = Direction(lineSplit[0])
		dig.Distance, err = strconv.Atoi(lineSplit[1])
		if err != nil {
			return nil, err
		}
		dig.Color = lineSplit[2][1: len(lineSplit[2]) - 1]
		digPlan = append(digPlan, dig)
	}
	return digPlan, nil

}

func parseFile2(filename string) ([]Dig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var digPlan []Dig
	for scanner.Scan() {
		lineStr := scanner.Text()
		var dig Dig
		lineSplit := strings.Split(lineStr, " ")
		if err != nil {
			return nil, err
		}
		dig.Color = lineSplit[2][1: len(lineSplit[2]) - 1]
		distanceHex := dig.Color[1:6]
		distance , _ := strconv.ParseInt(distanceHex, 16, 64)
		dig.Distance = int(distance)
		directionHex := dig.Color[6:]
		switch directionHex {
			case "0":
				dig.Direction = Right
			case "1":
				dig.Direction = Down
			case "2":
				dig.Direction = Left
			case "3":
				dig.Direction = Up
			}
		digPlan = append(digPlan, dig)


	}
	return digPlan, nil

}


func shoelaceFormula(pipePath []Coord) int {
	n := len(pipePath)
	if n < 3 {
		return 0.0
	}
	area := 0
	for i := 0; i < n-1; i++ {
		area += (pipePath[i].X*pipePath[i+1].Y) - (pipePath[i+1].X*pipePath[i].Y)
	}
	area += (pipePath[n-1].X*pipePath[0].Y) - (pipePath[0].X*pipePath[n-1].Y)
	area = int(math.Abs(float64(area)) / 2.0)

	return area
}

func picksFormula(perimeter, internalArea int) int {
	return internalArea + (perimeter/2) + 1
}

func part_1(filename string) {
	digPlan, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}
	position := Coord{0, 0}
	path := []Coord{}
	for _, _dig := range digPlan {
		dig := _dig
		newX := position.X
		newY := position.Y
		for i := 0; i < dig.Distance; i++ {
			switch dig.Direction {
			case Up:
				newY--
			case Down:
				newY++
			case Left:
				newX--
			case Right:
				newX++
			}
			newCoord := Coord{X: newX, Y: newY}
			path = append(path, newCoord)
		}
		position = Coord{X: newX, Y: newY}
	}
	
	count := shoelaceFormula(path)
	fmt.Println("Total Squares: ", picksFormula(len(path), count))
}

func part_2(filename string) {
	digPlan, err := parseFile2(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}
	position := Coord{0, 0}
	path := []Coord{}
	for _, _dig := range digPlan {
		dig := _dig
		newX := position.X
		newY := position.Y
		for i := 0; i < dig.Distance; i++ {
			switch dig.Direction {
			case Up:
				newY--
			case Down:
				newY++
			case Left:
				newX--
			case Right:
				newX++
			}
			newCoord := Coord{X: newX, Y: newY}
			path = append(path, newCoord)
		}
		position = Coord{X: newX, Y: newY}
	}
	
	count := shoelaceFormula(path)
	fmt.Println("Total Squares: ", picksFormula(len(path), count))


}

func main() {
	filename := "example.txt"
	part_1(filename)
	part_2(filename)
}