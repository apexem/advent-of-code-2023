package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	digit1 := FirstDigit(line)
	digit2 := FirstDigit(Reverse(line))
	parsed, err := strconv.Atoi(fmt.Sprint(digit1) + fmt.Sprint(digit2))
	if err != nil {
		panic(err)
	}
	return parsed
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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
			continue
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
	{word: "1", value: 1},
	{word: "2", value: 2},
	{word: "3", value: 3},
	{word: "4", value: 4},
	{word: "5", value: 5},
	{word: "6", value: 6},
	{word: "7", value: 7},
	{word: "8", value: 8},
	{word: "9", value: 9},
	{word: "one", value: 1},
	{word: "two", value: 2},
	{word: "three", value: 3},
	{word: "four", value: 4},
	{word: "five", value: 5},
	{word: "six", value: 6},
	{word: "seven", value: 7},
	{word: "eight", value: 8},
	{word: "nine", value: 9},
	{word: "eno", value: 1},
	{word: "owt", value: 2},
	{word: "eerht", value: 3},
	{word: "ruof", value: 4},
	{word: "evif", value: 5},
	{word: "xis", value: 6},
	{word: "neves", value: 7},
	{word: "thgie", value: 8},
	{word: "enin", value: 9},
}

func FirstDigit(line string) int {
	indices := []struct {
		value int
		index int
	}{}
	for _, val := range wordToDigit {
		index := strings.Index(line, val.word)
		if index == -1 {
			continue
		}
		indices = append(indices, struct {
			value int
			index int
		}{value: val.value, index: index})
	}
	sort.SliceStable(indices, func(i, j int) bool {
		return indices[i].index < indices[j].index
	})
	return indices[0].value
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
