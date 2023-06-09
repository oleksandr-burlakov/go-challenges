package main

import (
	"fmt"
	"helpers"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("puzzle_input")
	helpers.Check(err)
	result := calculatePoints(string(data))
	fmt.Println("My score will be", result)
}

func calculatePoints(data string) (points int) {
	rows := strings.Split(data, "\n")
	for _, row := range rows {
		points += determineRowPoints(row)
	}
	return points
}

func determineRowPoints(row string) (points int) {
	parts := strings.Split(row, " ")
	opponent := parts[0]
	my := parts[1]
	points += determineMatchStatePoints(opponent, my)
	points += determineActionPoints(my)
	return points
}

func determineMatchStatePoints(opponent string, my string) int {
	if (opponent == "A" && my == "X") ||
		(opponent == "B" && my == "Y") ||
		(opponent == "C" && my == "Z") {
		return 3
	} else if (opponent == "A" && my == "Y") ||
		(opponent == "B" && my == "Z") ||
		(opponent == "C" && my == "X") {
		return 6
	} else {
		return 0
	}
}

func determineActionPoints(action string) (points int) {
	switch action {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		return 0
	}
}
