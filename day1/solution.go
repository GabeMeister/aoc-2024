package day1

import (
	"GabeMeister/aoc-2024/helpers"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	lines := helpers.ReadFile("./day1/small.txt")
	left := []int{}
	right := []int{}

	for _, line := range lines {
		tokens := strings.Split(line, "   ")
		num, _ := strconv.Atoi(tokens[0])
		left = append(left, num)

		num, _ = strconv.Atoi(tokens[1])
		right = append(right, num)
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0

	for i, num := range left {
		leftNum := num
		rightNum := right[i]

		result := helpers.Abs(leftNum - rightNum)

		sum += result
	}

	fmt.Print("\n\n", "*** sum ***", "\n", sum, "\n\n\n")
}

func Part2() {
	lines := helpers.ReadFile("./day1/big.txt")
	left := []int{}
	right := []int{}

	for _, line := range lines {
		tokens := strings.Split(line, "   ")
		num, _ := strconv.Atoi(tokens[0])
		left = append(left, num)

		num, _ = strconv.Atoi(tokens[1])
		right = append(right, num)
	}

	slices.Sort(left)
	slices.Sort(right)

	similarityScore := 0

	// Number -> Frequency
	rightMap := make(map[int]int)

	for _, num := range right {
		rightMap[num]++
	}

	for _, num := range left {
		leftNum := num

		freq, ok := rightMap[leftNum]
		if !ok {
			freq = 0
		}

		score := num * freq
		similarityScore += score
	}

	fmt.Print("\n\n", "*** similarityScore ***", "\n", similarityScore, "\n\n\n")
}
