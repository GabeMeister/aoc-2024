package day3

import (
	"GabeMeister/aoc-2024/helpers"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	total := 0
	lines := helpers.ReadFile("./day3/big.txt")

	// Compile the regex pattern
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	for _, line := range lines {
		matches := regex.FindAllString(line, -1)

		for _, match := range matches {
			numPair := strings.ReplaceAll(match, "mul(", "")
			numPair = strings.ReplaceAll(numPair, ")", "")
			tokens := strings.Split(numPair, ",")

			nums := []int{}
			for _, token := range tokens {
				num, _ := strconv.Atoi(token)
				nums = append(nums, num)
			}

			total += nums[0] * nums[1]
		}
	}

	fmt.Print("\n\n", "*** total ***", "\n", total, "\n\n\n")
}

type Command struct {
	Start       int
	End         int
	CommandType string
	Text        string
}

func Part2() {
	lines := helpers.ReadFile("./day3/big.txt")
	line := ""
	for _, l := range lines {
		line += l
	}

	// Compile the mulRegex pattern
	mulRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	commands := []Command{}
	total := 0

	mulIndices := mulRegex.FindAllStringIndex(line, -1)

	for _, mulIdx := range mulIndices {
		start := mulIdx[0]
		end := mulIdx[1]

		commands = append(commands, Command{
			Start:       start,
			End:         end,
			CommandType: "mul",
			Text:        line[start:end],
		})
	}

	doIndices := doRegex.FindAllStringIndex(line, -1)
	for _, doIdx := range doIndices {
		start := doIdx[0]
		end := doIdx[1]

		commands = append(commands, Command{
			Start:       start,
			End:         end,
			CommandType: "do",
			Text:        line[start:end],
		})
	}

	dontIndices := dontRegex.FindAllStringIndex(line, -1)
	for _, dontIdx := range dontIndices {
		start := dontIdx[0]
		end := dontIdx[1]

		commands = append(commands, Command{
			Start:       start,
			End:         end,
			CommandType: "dont",
			Text:        line[start:end],
		})
	}

	slices.SortFunc(commands, func(a Command, b Command) int {
		if a.Start < b.Start {
			return -1
		} else if a.Start > b.Start {
			return 1
		} else {
			return 0
		}
	})

	enabled := true

	for _, command := range commands {
		switch command.CommandType {
		case "mul":
			if enabled {
				numPair := strings.ReplaceAll(command.Text, "mul(", "")
				numPair = strings.ReplaceAll(numPair, ")", "")
				tokens := strings.Split(numPair, ",")

				nums := []int{}
				for _, token := range tokens {
					num, _ := strconv.Atoi(token)
					nums = append(nums, num)
				}

				total += nums[0] * nums[1]
			}
		case "do":
			enabled = true
		case "dont":
			enabled = false
		}
	}

	fmt.Print("\n\n", "*** total pt 2 ***", "\n", total, "\n\n\n")
}
