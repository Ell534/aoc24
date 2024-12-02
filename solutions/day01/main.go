package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const file string = "../../inputs/day01.txt"

var leftNumbers []int
var rightNumbers []int

func main() {
	data := readFile(file)
	splitNumbers(data)
	sortNumbers(leftNumbers)
	sortNumbers(rightNumbers)
	result := findDistance(leftNumbers, rightNumbers)
	fmt.Println(result)
}

func readFile(filePath string) []string {
	var data = []string{}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func splitNumbers(data []string) {
	for _, values := range data {
		splitValues := strings.Split(values, "   ")
		leftValue, _ := strconv.Atoi(splitValues[0])
		rightValue, _ := strconv.Atoi(splitValues[1])
		leftNumbers = append(leftNumbers, leftValue)
		rightNumbers = append(rightNumbers, rightValue)
	}
}

func sortNumbers(numbers []int) {
	slices.SortStableFunc(numbers, func(a, b int) int {
		return cmp.Compare(a, b)
	})
}

func findDistance(leftList, rightList []int) int {
	totalDistance := 0

	for i := range rightList {
		currentDistance := rightList[i] - leftList[i]
		if currentDistance < 0 {
			currentDistance = -currentDistance
		}
		totalDistance += currentDistance
	}
	return totalDistance
}
