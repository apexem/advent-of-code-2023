package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		calibrationValue := processLine(line)
		total += calibrationValue
	}
	return total, nil
}

func processLine(line string) int {
	numbers := lineToNumbers(line)
	digit1, digit2 := numbers[0], numbers[len(numbers)-1]
	parsed, err := strconv.Atoi(fmt.Sprint(digit1) + fmt.Sprint(digit2))
	if err != nil {
		panic(err)
	}
	return parsed
}

func lineToNumbers(line string) []int {
	i, j := 0, 1
	result := []int{}
	for {
		if i == len(line) {
			break
		}
		if j > len(line) {
			break
		}
		val := line[i:j]
		num, err := strconv.Atoi(val)
		if err == nil {
			result = append(result, num)
			i += 1
			j += 1
			continue
		}
		contains, exact, number := doesMatch(val)
		if contains && exact {
			i = j - 1
			result = append(result, number)
			continue
		}
		if contains && !exact {
			j += 1
			continue
		}
		if !contains && i == j-1 {
			i = j
			j++
		}
		if !contains {
			i = j - 1
			continue
		}
	}
	return result
}

var wordToDigit = []struct {
	word  string
	value int
}{
	{word: "one", value: 1},
	{word: "two", value: 2},
	{word: "three", value: 3},
	{word: "four", value: 4},
	{word: "five", value: 5},
	{word: "six", value: 6},
	{word: "seven", value: 7},
	{word: "eight", value: 8},
	{word: "nine", value: 9},
}

func doesMatch(word string) (bool, bool, int) {
	for _, val := range wordToDigit {
		if val.word == word {
			return true, true, val.value
		}
		if strings.HasPrefix(val.word, word) {
			return true, false, -1
		}
	}
	return false, false, -1
}
