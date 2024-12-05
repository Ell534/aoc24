package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const FILE = "../../inputs/day03.txt"
const TEST = "../../inputs/day03Test.txt"

func main() {
	data := readFile(FILE)
	extractedPairs := regexpExtraction(data)
	integers := convertToInt(extractedPairs)
	total := multiplyAndSum(integers)
	fmt.Println(total)

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

func regexpExtraction(input []string) [][]string {
	extractedMatches := [][]string{}
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)
	for _, line := range input {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			extractedMatches = append(extractedMatches, match[1:])
		}
	}
	return extractedMatches
}

func convertToInt(stringPairs [][]string) [][]int {
	intPairs := [][]int{}
	for _, pair := range stringPairs {
		firstVal, _ := strconv.Atoi(pair[0])
		secondVal, _ := strconv.Atoi(pair[1])
		intPair := []int{firstVal, secondVal}
		intPairs = append(intPairs, intPair)
	}
	return intPairs
}

func multiplyAndSum(intPairs [][]int) int {
	total := 0
	for _, pair := range intPairs {
		valueToAdd := pair[0] * pair[1]
		total += valueToAdd
	}
	return total
}
