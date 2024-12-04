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
	fmt.Println(convertedData[0])
	totalSafeReports := checkIfSafe(convertedData)
	fmt.Println(totalSafeReports)
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

// split each element on " " into their individual string numbers
// convert each string number into integer and put into a slice
// put that slice of ints into a slice containing all the slices of ints i.e [][]int
// run checks to see if it meets the conditions: must be either all increasing or all decreasing, AND each adjacent number can only differ by at least one AND at most 3
// check for all increasing or decreasing

func convertToInt(data []string) [][]int {
	convertedData := [][]int{}
	for _, line := range data {
		intNumberReport := []int{}
		stringNumbersReport := strings.Split(line, " ")
		//convert to ints
		for _, number := range stringNumbersReport {
			intValue, _ := strconv.Atoi(number)
			intNumberReport = append(intNumberReport, intValue)
		}
		convertedData = append(convertedData, intNumberReport)
	}
	return convertedData
}

func checkIfSafe(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		safe := true

		var direction string
		if report[0] < report[1] {
			direction = "up"
		} else if report[0] > report[1] {
			direction = "down"
		} else {
			safe = false
		}

		// check for all either moving up or down in value, if not set safe to false and continue to next report
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

		if safe == true {
			safeReports++
		}

	}
	return safeReports
}
