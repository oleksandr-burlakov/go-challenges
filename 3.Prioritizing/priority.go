package main

import (
	"fmt"
	"helpers"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input")
	helpers.Check(err)
	result := analyzePackagesList(string(data))
	fmt.Println("Result is", result)
}

func analyzePackagesList(data string) (ret int) {
	rows := strings.Split(data, "\n")
	for _, row := range rows {
		ret += analyzeRow(row)
	}
	return ret
}

func analyzeRow(row string) (ret int) {
	letter := findCommonPackage(row)
	fmt.Println("Letter number:", letter)
	ret = calculateLetterValue(letter)
	return ret
}

func calculateLetterValue(letter int) (ret int) {
	if letter < 0 {
		ret = 0
	} else if letter > 60 {
		ret = letter - 60
	} else {
		ret = letter - 40
	}
	return ret
}

func findCommonPackage(row string) (ret int) {
	letterMap := make(map[rune]int)
	for _, char := range row {
		if val, ok := letterMap[char]; ok {
			letterMap[char] = val + 1
		} else {
			letterMap[char] = 0
		}
	}
	for k, v := range letterMap {
		if v > 1 {
			return int(k)
		}
	}
	return -1
}
