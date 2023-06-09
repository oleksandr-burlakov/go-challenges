package main

import (
	"fmt"
	"helpers"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("input-file")
	check(err)
	fmt.Println(string(dat))
	calories := splitCaloryList(string(dat))
	maxCalorie := helpers.Max(calories)
	for i, cal := range calories {
		fmt.Println("Elf number", i, "have", cal, "calories...")
		if maxCalorie == cal {
			fmt.Println("By the way, this elf also have largest amount of calories")
		}
	}
}

func splitCaloryList(data string) []int {
	result := []int{}
	sum := 0
	calorieNumbers := strings.Split(data, "\n")
	for _, cal := range calorieNumbers {
		if cal == "" {
			result = append(result, sum)
			sum = 0
			continue
		}
		calNum, err := strconv.Atoi(cal)
		check(err)
		sum += calNum
	}
	result = append(result, sum)
	return result
}
