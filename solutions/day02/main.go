package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILE = "../../inputs/day02.txt"
const TEST = "../../inputs/day02Test.txt"

func main() {
	inputData := readFile(FILE)
	convertedData := convertToInt(inputData)
	totalSafeReports := countSafeReports(convertedData)
	fmt.Printf("Part 1 Answer: %d\n", totalSafeReports)
	totalDampenedSafeReports := dampenReports(convertedData)
	fmt.Printf("Part 2 Answer: %d\n", totalDampenedSafeReports)

}

func readFile(filePath string) []string {
	data := []string{}

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

func convertToInt(data []string) [][]int {
	convertedData := [][]int{}
	for _, line := range data {
		intNumberReport := []int{}
		stringNumbersReport := strings.Split(line, " ")

		for _, number := range stringNumbersReport {
			intValue, _ := strconv.Atoi(number)
			intNumberReport = append(intNumberReport, intValue)
		}
		convertedData = append(convertedData, intNumberReport)
	}
	return convertedData
}

func checkIfSafe(report []int) bool {

	safe := true

	var direction string
	if report[0] < report[1] {
		direction = "up"
	} else if report[0] > report[1] {
		direction = "down"
	} else {
		safe = false
		return safe
	}

	// check if all values are either moving up or down in value, if not set safe to false and continue to next report
	for i := 0; i < len(report)-1; i++ {
		currentDirection := ""
		if report[i] < report[i+1] {
			currentDirection = "up"
		} else if report[i] > report[i+1] {
			currentDirection = "down"
		}
		if direction != currentDirection {
			safe = false
			break
		}
	}

	// check for differences in value, at least 1 and no more than 3
	for i := 0; i < len(report)-1; i++ {
		difference := report[i] - report[i+1]
		if difference < 0 {
			difference = -difference
		}
		if difference > 0 && difference <= 3 {
			continue
		} else {
			safe = false
			break
		}
	}

	return safe
}

func countSafeReports(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		if checkIfSafe(report) {
			safeReports++
		}
	}
	return safeReports
}

// for each report, remove one element and see if it is made safe, if so add one to the count and move on to the next report, if not remove the next element and check again etc
func dampenReports(reports [][]int) int {
	totalDampened := 0

	for _, report := range reports {
		dampened := false

		for i := 0; i < len(report); i++ {
			testSlice := make([]int, 0, len(report)-1)
			testSlice = append(testSlice, report[:i]...)
			testSlice = append(testSlice, report[i+1:]...)

			if checkIfSafe(testSlice) {
				dampened = true
			}
		}

		if dampened {
			totalDampened++
		}

	}
	return totalDampened
}
