package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var wordToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func convertToDigit(word string) int {
	if val, ok := wordToDigit[word]; ok {
		return val
	}

	return int(word[0] - '0')
}

func matchDigits(line string) []string {
	re := regexp.MustCompile(`^(one|two|three|four|five|six|seven|eight|nine|\d)`)
	var matches []string
	for i := range line {
		subLine := line[i:]
		if match := re.FindString(subLine); match != "" {
			matches = append(matches, match)
		}
	}
	return matches
}

func extractFirstAndLastDigitsFromLine(line string) (int, int, error) {
	matches := matchDigits(line)
	first := convertToDigit(matches[0])
	last := convertToDigit(matches[len(matches)-1])

	return first, last, nil
}

func extractCalibrationValueFromLine(line string) (calibrationValue int) {
	first, last, err := extractFirstAndLastDigitsFromLine(line)
	check(err)

	calibrationValue = int(first)*10 + int(last)

	return
}

func calculateSum(scanner *bufio.Scanner) (sum int) {
	for scanner.Scan() {
		line := scanner.Text()
		sum += extractCalibrationValueFromLine(line)
	}

	return
}

func main() {
	file, err := os.Open("./input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	check(scanner.Err())

	sum := calculateSum(scanner)

	fmt.Println(sum)
}
