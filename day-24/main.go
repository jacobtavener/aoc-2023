package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y, Z float64
}

type Velocity struct {
	X, Y, Z float64
}

// ax + by + c = 0
type Equation struct {
	a, b, c float64
}

// ax + by + cz = d
type Equation3D struct {
	a, b, c, d float64
}

type Hailstone struct {
	Position  Coord
	Velocity  Velocity
	Equation  Equation
}

func cleanseString(str string) string {
	return strings.TrimSpace(strings.ReplaceAll(str, "  ", " "))
}

func convertStringListToFloats(stringList []string) ([]float64, error) {
	var floatList []float64

	for _, str := range stringList {
		// Convert each string to an integer
		str = cleanseString(str)
		i, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		floatList = append(floatList, float64(i))
	}

	return floatList, nil
}

func parseFile(filename string) ([]Hailstone, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hailstones := make([]Hailstone, 0)
	for scanner.Scan() {
		rowStr := scanner.Text()
		row := strings.Split(rowStr, " @ ")
		coords, _ := convertStringListToFloats(strings.Split(row[0], ", "))
		velocities, _ := convertStringListToFloats(strings.Split(row[1], ", "))
		px, py, pz := coords[0], coords[1], coords[2]
		vx, vy, vz := velocities[0], velocities[1], velocities[2]

		a := -float64(vy) / float64(vx)
		b := 1.0
		c := -(float64(py) + a*float64(px))

		hailstones = append(hailstones, Hailstone{
			Position: Coord{px, py, pz},
			Velocity: Velocity{vx, vy, vz},
			Equation: Equation{a, b, c},
		})
	}
	return hailstones, nil
}

func findIntersection(equation1, equation2 Equation) (float64, float64) {
	a1, b1, c1 := equation1.a, equation1.b, equation1.c
	a2, b2, c2 := equation2.a, equation2.b, equation2.c

	x := (b1*c2 - b2*c1) / (a1*b2 - a2*b1)
	y := (a2*c1 - a1*c2) / (a1*b2 - a2*b1)
	return x, y
}

func isFutureIntersection(hailstone Hailstone, x, y float64) bool {
	px, py := hailstone.Position.X, hailstone.Position.Y
	vx, vy := hailstone.Velocity.X, hailstone.Velocity.Y

	futureX := true
	futureY := true

	if vx > 0 {
		futureX = x > px
	} else if vx < 0 {
		futureX = x < px
	}

	if vy > 0 {
		futureY = y > py
	} else if vy < 0 {
		futureY = y < py
	}

	return futureX && futureY
}

func part_1(filename string) {
	hailstones, err := parseFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lowerBound := 200000000000000.0
	upperBound := 400000000000000.0
	
	total := 0
	for i := 0; i < len(hailstones); i++ {
		hailstone1 := hailstones[i]
		equation1 := hailstone1.Equation
	
		for j := i + 1; j < len(hailstones); j++ {
			hailstone2 := hailstones[j]
			equation2 := hailstone2.Equation
	
			x, y := findIntersection(equation1, equation2)
			withinTestArea := x >= lowerBound && x <= upperBound && y >= lowerBound && y <= upperBound
			if withinTestArea && isFutureIntersection(hailstone1, x, y) && isFutureIntersection(hailstone2, x, y){
				total++
			}
		}
	}
	fmt.Println("Total:", total)
}

func main() {
	filename := "full.txt"
	part_1(filename)
}
