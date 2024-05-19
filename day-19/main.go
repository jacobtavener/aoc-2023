package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
)

type ComparisonOperator interface {
	Compare(int, int) bool
}

type LessThan struct{}

func (lt LessThan) Compare(x, y int) bool {
	return x < y
}

type GreaterThan struct{}

func (gt GreaterThan) Compare(x, y int) bool {
	return x > y
}

type NoOperation struct{}

func (noOp NoOperation) Compare(x, y int) bool {
	return true // No operation always returns true
}

func parseComparisonExpression(expression string) (ComparisonOperator, string, int, error) {
	operators := map[string]ComparisonOperator{
		"<": LessThan{},
		">": GreaterThan{},
	}

	for op, operatorImpl := range operators {
		if strings.Contains(expression, op) {
			parts := strings.Split(expression, op)
			if len(parts) != 2 {
				return nil, "", 0, fmt.Errorf("Invalid expression: %s", expression)
			}

			left := strings.TrimSpace(parts[0])

			right, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				return nil, "", 0, fmt.Errorf("Invalid right operand: %s", parts[1])
			}

			return operatorImpl, left, right, nil
		}
	}

	return nil, "", 0, fmt.Errorf("No valid operator found in expression: %s", expression)
}

type Condition struct {
	Operator ComparisonOperator
	Value    int
}

type Stage struct {
	NextStage string
	Condition Condition
	Part string
}

type Workflow struct {
	Stages []Stage
	FallBack string

}

type Rating struct {
	x, m, a, s int
}

func (r Rating) Sum() int {
	return r.x + r.m + r.a + r.s
}

func parseFile(filename string) (map[string]Workflow, []Rating, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	parsingWorkflows := true
	ratings := make([]Rating, 0)
	workflowMap := make(map[string]Workflow)
	for scanner.Scan() {
		lineStr := scanner.Text()
		if lineStr == "" {
			parsingWorkflows = false
		} else if parsingWorkflows {
			lineSplit := strings.Split(lineStr, "{")
			workflowName := lineSplit[0]
			workflowStr := lineSplit[1]
			workflowStagesSplit := strings.Split(workflowStr[:len(workflowStr)-1], ",")
			numStages := len(workflowStagesSplit)
			workflowStages := make([]Stage, 0)
			for _, stage := range workflowStagesSplit[:numStages-1] {
				stageSplit := strings.Split(stage, ":")
				condition, nextStage := stageSplit[0], stageSplit[1]
				operator, part, value, _ := parseComparisonExpression(condition)
				stage := Stage{NextStage: nextStage, Condition: Condition{Operator: operator, Value: value}, Part: part}
				workflowStages = append(workflowStages, stage)
			}
			fallBack := workflowStagesSplit[numStages-1]
			workflowMap[workflowName] = Workflow{Stages: workflowStages, FallBack: fallBack}
		} else {
			ratingStr := lineStr[1 : len(lineStr)-1]
			ratingStrSplit := strings.Split(ratingStr, ",")
			var x, m, a, s int
			for _, rating := range ratingStrSplit {
				ratingSplit := strings.Split(rating, "=")
				val, _ := strconv.Atoi(ratingSplit[1])
				switch ratingSplit[0] {
				case "x":
					x = val
				case "m":
					m = val
				case "a":
					a = val
				case "s":
					s = val
				}

			}
			rating := Rating{x: x, m: m, a: a, s: s}
			ratings = append(ratings, rating)
		}
	}
	return workflowMap, ratings, nil
}

func part_1(filename string) {
	workflowMap, ratings, err := parseFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	total := 0
	for _, rating := range ratings {

		key := "in"
		workflow := workflowMap[key]
		endKeyFound := false
		endKey := ""
		allKeys := []string{key}
		for !endKeyFound {
			useFallback := true
			for _, stage := range workflow.Stages {
				val := 0
				switch stage.Part {
					case "x":
						val = rating.x
					case "m":
						val = rating.m
					case "a":
						val = rating.a
					case "s":
						val = rating.s
				}
				passedStage := stage.Condition.Operator.Compare(val, stage.Condition.Value)
				if passedStage {
					key = stage.NextStage
					useFallback = false
					allKeys = append(allKeys, key)
					if key == "A" || key == "R" {
						endKey = key
						endKeyFound = true
						break
					}
					workflow = workflowMap[key]
					break
				}
			}
			if useFallback {
				key = workflow.FallBack
			}
			allKeys = append(allKeys, key)
			if key == "A" || key == "R" {
				endKey = key
				endKeyFound = true
				break
			} else {
				workflow = workflowMap[key]
			}
		}
		if endKey == "A" {
			total += rating.Sum()
		}

	}
	fmt.Println("Total: ", total)

}

type RatingRange struct {
	Min, Max int
}

type RatingRanges struct {
	x, m, a, s RatingRange
}

func checkOperation(operator ComparisonOperator, ratingRange RatingRange, value int) (RatingRange, RatingRange) {
	// returns RatingRange that is in range, and RatingRange that is out of range
	currentMin := ratingRange.Min
	currentMax := ratingRange.Max
	switch operator.(type) {
		case LessThan:
			if value < currentMin {
				return RatingRange{}, ratingRange
			}
			if value > currentMax {
				return ratingRange, RatingRange{}
			} else {
				return RatingRange{Min: currentMin, Max: value-1}, RatingRange{Min: value, Max: currentMax}
			}
		case GreaterThan:
			if value > currentMax {
				return RatingRange{}, ratingRange
				} else if value < currentMin {
				return ratingRange, RatingRange{}
			} else {
				return RatingRange{Min: value+1, Max: currentMax}, RatingRange{Min: currentMin, Max: value}
			}
		case NoOperation:
			return ratingRange, RatingRange{}	
	}
	return ratingRange, RatingRange{}
}

func processWorkflow(workflow Workflow, workflowMap map[string]Workflow, _ratingRanges RatingRanges, acceptedRanges []RatingRanges) []RatingRanges {
	ratingRanges := _ratingRanges
	for _, stage := range workflow.Stages {
		var ratingRange RatingRange
		switch stage.Part {
			case "x":
				ratingRange = ratingRanges.x
			case "m":
				ratingRange = ratingRanges.m
			case "a":
				ratingRange = ratingRanges.a
			case "s":
				ratingRange = ratingRanges.s
		}

		correctRatingRange, incorrectRatingRange := checkOperation(stage.Condition.Operator, ratingRange, stage.Condition.Value)
		if correctRatingRange.Max - correctRatingRange.Min != 0 {
			// there are correct values
			key := stage.NextStage
			var newRatingRanges RatingRanges
			switch stage.Part {
				case "x":
					newRatingRanges = RatingRanges{x: correctRatingRange, m: ratingRanges.m, a: ratingRanges.a, s: ratingRanges.s}
				case "m":
					newRatingRanges = RatingRanges{x: ratingRanges.x, m: correctRatingRange, a: ratingRanges.a, s: ratingRanges.s}
				case "a":
					newRatingRanges = RatingRanges{x: ratingRanges.x, m: ratingRanges.m, a: correctRatingRange, s: ratingRanges.s}
				case "s":
					newRatingRanges = RatingRanges{x: ratingRanges.x, m: ratingRanges.m, a: ratingRanges.a, s: correctRatingRange}

			}
			if key == "A" || key == "R" {
				if key == "A" {
					fmt.Println("Correct Rating Range: ", newRatingRanges)
					acceptedRanges = append(acceptedRanges, newRatingRanges)
				}
			} else {
				newWorkflow := workflowMap[key]
				acceptedRanges = processWorkflow(newWorkflow, workflowMap, newRatingRanges, acceptedRanges)
			}
		}
		if incorrectRatingRange.Max - incorrectRatingRange.Min != 0 {
			// there are incorrect values continue in stages loop - must redefine ranges.
			switch stage.Part {
			case "x":
				ratingRanges.x = incorrectRatingRange
			case "m":
				ratingRanges.m = incorrectRatingRange
			case "a":
				ratingRanges.a = incorrectRatingRange
			case "s":
				ratingRanges.s = incorrectRatingRange
			}
			fmt.Println("Incorrect Rating Range: ", incorrectRatingRange, "I am now running")
		}
	}
	key := workflow.FallBack
	if key == "A" || key == "R" {
		if key == "A" {
			fmt.Println("Correct Rating Range: ", ratingRanges)
			acceptedRanges = append(acceptedRanges, ratingRanges)
		}
	} else {
		newWorkflow := workflowMap[key]
		newRatingRanges := RatingRanges{x: ratingRanges.x, m: ratingRanges.m, a: ratingRanges.a, s: ratingRanges.s}
		acceptedRanges = processWorkflow(newWorkflow, workflowMap, newRatingRanges, acceptedRanges)
	}
	return acceptedRanges
}

func part_2(filename string) {
	workflowMap, _, err := parseFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(workflowMap)
	ratingRanges := RatingRanges{
		x: RatingRange{Min: 1, Max: 4000},
		m: RatingRange{Min: 1, Max: 4000},
		a: RatingRange{Min: 1, Max: 4000},
		s: RatingRange{Min: 1, Max: 4000},
	}

	key := "in"
	workflow := workflowMap[key]

	acceptedRatingRanges := make([]RatingRanges, 0)

	acceptedRanges := processWorkflow(workflow, workflowMap, ratingRanges, acceptedRatingRanges)

	totalTotal := 0
	for _, acceptedRange := range acceptedRanges {
		total := (acceptedRange.a.Max - acceptedRange.a.Min + 1) * (acceptedRange.m.Max - acceptedRange.m.Min + 1) * (acceptedRange.x.Max - acceptedRange.x.Min + 1) * (acceptedRange.s.Max - acceptedRange.s.Min + 1)
		totalTotal += total


	}
	fmt.Println("Total Total: ", totalTotal)
}

func main() {
	filename := "full.txt"
	// part_1(filename)
	part_2(filename)
}