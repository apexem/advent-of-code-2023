package day2

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solution(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	possible := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	scanner := bufio.NewScanner(file)
	compiled, _ := regexp.Compile("(\\d+)\\s*(red|blue|green)")
	ids := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		match := compiled.FindAllString(line, -1)
		gameValid := true
		for _, match := range match {
			split := strings.Fields(match)
			val, _ := strconv.Atoi(split[0])
			color := split[1]
			if val > possible[color] {
				gameValid = false
				break
			}
		}
		if gameValid {
			id := getGameId(line)
			ids = append(ids, id)
		}
	}

	sum := 0
	for _, index := range ids {
		sum += index
	}
	return sum, nil
}

func getGameId(line string) int {
	compiled, err := regexp.Compile("Game (?P<id>\\d+)")
	if err != nil {
		panic(err)
	}
	match := compiled.FindStringSubmatch(line)
	parsed, _ := strconv.Atoi(match[1])
	return parsed
}

type Set struct {
	color      string
	ocurrences int
}
