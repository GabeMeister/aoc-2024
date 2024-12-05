package day2

import (
	"GabeMeister/aoc-2024/helpers"
	"fmt"
	"strconv"
	"strings"
)

func isIncreasing(nums []int) bool {
	return nums[0] < nums[1]
}

func isSafelyIncreasing(nums []int, skipIdx int) bool {
	tempNums := []int{}

	for i, num := range nums {
		if i == skipIdx {
			continue
		}

		tempNums = append(tempNums, num)
	}

	current := tempNums[0]
	for _, num := range tempNums[1:] {
		if num >= current+1 && num <= current+3 {
			current = num
		} else {
			return false
		}
	}

	return true
}

func isSafelyDecreasing(nums []int, skipIdx int) bool {
	tempNums := []int{}

	for i, num := range nums {
		if i == skipIdx {
			continue
		}

		tempNums = append(tempNums, num)
	}

	current := tempNums[0]
	for _, num := range tempNums[1:] {
		if num >= current-3 && num <= current-1 {
			current = num
		} else {
			return false
		}
	}

	return true
}

func Part1() {
	total := 0

	lines := helpers.ReadFile("./day2/big.txt")
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		nums := []int{}
		for _, token := range tokens {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}

		if isIncreasing(nums) {
			result := isSafelyIncreasing(nums, -1)
			if result {
				total += 1
			}
		} else {
			result := isSafelyDecreasing(nums, -1)
			if result {
				total += 1
			}
		}
	}

	fmt.Print("\n\n", "*** total safe reports ***", "\n", total, "\n\n\n")
}

func isSafelyIncreasingWithDampener(nums []int) bool {
	// If it's already good, then no need to continue
	if isSafelyIncreasing(nums, -1) {
		return true
	}

	// Try "skipping" each number and see if it's safe
	for i := range nums {
		if isSafelyIncreasing(nums, i) {
			return true
		}
	}

	return false
}

func isSafelyDecreasingWithDampener(nums []int) bool {
	// If it's already good, then no need to continue
	if isSafelyDecreasing(nums, -1) {
		return true
	}

	// Try "skipping" each number and see if it's safe
	for i := range nums {
		if isSafelyDecreasing(nums, i) {
			return true
		}
	}

	return false
}

func Part2() {
	total := 0

	lines := helpers.ReadFile("./day2/big.txt")
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		nums := []int{}
		for _, token := range tokens {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}

		if isSafelyIncreasingWithDampener(nums) {
			total += 1
		} else if isSafelyDecreasingWithDampener(nums) {
			total += 1
		}
	}

	fmt.Print("\n\n", "*** total safe reports ***", "\n", total, "\n\n\n")

}
