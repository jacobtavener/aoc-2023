package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func parseFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var grid [][]int
	for scanner.Scan() {
		lineStr := scanner.Text()
		var line []int
		for _, str := range lineStr {
			i, err := strconv.Atoi(string(str))
			if err != nil {
				return nil, err
			}
			line = append(line, i)
		}
		grid = append(grid, line)
	}
	return grid, nil

}

type Vertex struct {
	X int
	Y int
}

type State struct {
	Position Vertex
	Direction Vertex
	DirectionCount int
}

type Item struct {
	State    *State
	Distance int
	Index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func initializePriorityQueue() PriorityQueue {
	priorityQueue := make(PriorityQueue, 0)

	heap.Push(&priorityQueue, &Item{
		State: &State{
			Position: Vertex{},
			Direction: Vertex{X: 1},
			DirectionCount: 0,
		},
		Distance: 0,
	})
	heap.Push(&priorityQueue, &Item{
		State: &State{
			Position: Vertex{},
			Direction: Vertex{Y: 1},
			DirectionCount: 0,
		},
		Distance: 0,
	})
	return priorityQueue
}

var PossibleDirections = []Vertex{
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 0, Y: -1},
}

func findPossibleNextMoves(grid [][]int, state *State, maxMoves, minMoves int) []State {
	position := state.Position
	x := position.X
	y := position.Y
	result := make([]State, 0)

	for _, dir := range PossibleDirections {
		newX := x + dir.X
		newY := y + dir.Y
		newDir := dir
		
		var dirCount int
		if newDir.X == state.Direction.X && newDir.Y == state.Direction.Y {
			dirCount = state.DirectionCount + 1
		} else {
			dirCount = 1	
		}

		outsideOfGrid := !(newX >= 0 && newY >=0 && newX < len(grid[0]) && newY < len(grid))
		tooFar := dirCount > maxMoves
		oppositeDirection := newDir.X * -1 == state.Direction.X && newDir.Y * -1 == state.Direction.Y
		wrongDirection := newDir.X != state.Direction.X && newDir.Y != state.Direction.Y && state.DirectionCount < minMoves

		if !(outsideOfGrid || tooFar || oppositeDirection || wrongDirection) {
			result = append(result, State{
				Position: Vertex{X: newX, Y: newY},
				Direction: newDir,
				DirectionCount: dirCount,
			})
		}
	}
	return result
}

func dijkstra(grid [][]int, priorityQueue PriorityQueue, maxMoves, minMoves int) map[State]int {
	minDistanceMap := make(map[State]int)
	
	for priorityQueue.Len() > 0 {
		item := heap.Pop(&priorityQueue).(*Item)
		state := item.State
		distance := item.Distance
		if _, exists := minDistanceMap[*state]; !exists {
			minDistanceMap[*state] = distance
	
			possibleNextMoves := findPossibleNextMoves(grid, state, maxMoves, minMoves)
			for _, possibleNextMove := range possibleNextMoves {
				nextState := possibleNextMove
				if _, exists := minDistanceMap[nextState]; !exists {
					distanceFromNode := distance + grid[nextState.Position.Y][nextState.Position.X]

					heap.Push(&priorityQueue, &Item{
						State: &nextState,
						Distance: distanceFromNode,
					})
				}
			}
		}

	}
	return minDistanceMap
	
}

func part_1(filename string) {
	grid, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}

	PriorityQueue := initializePriorityQueue()
	minDistanceMap := dijkstra(grid, PriorityQueue, 3, 0)
	minDistance := 0
	for state, distance := range minDistanceMap {
		if state.Position.X == len(grid[0]) - 1 && state.Position.Y == len(grid) - 1 {
			if minDistance == 0 || distance < minDistance {
				minDistance = distance
			}
		}
	}
	fmt.Println("Part 1:", minDistance)
}

func part_2(filename string) {
	grid, err := parseFile(filename)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}

	PriorityQueue := initializePriorityQueue()
	minDistanceMap := dijkstra(grid, PriorityQueue, 10, 4)
	minDistance := 0
	for state, distance := range minDistanceMap {
		if state.Position.X == len(grid[0]) - 1 && state.Position.Y == len(grid) - 1 {
			if minDistance == 0 || distance < minDistance {
				minDistance = distance
			}
		}
	}
	fmt.Println("Part 2:", minDistance)
}

func main() {
	filename := "full.txt"
	part_1(filename)
	part_2(filename)
}	