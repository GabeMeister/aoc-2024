package day6

import (
	"GabeMeister/aoc-2024/helpers"
	"fmt"
	"strings"
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Location [2]int

func getRotatedDirection(direction Direction) Direction {
	switch direction {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	}

	panic("Should not have gotten here")
}

func Part1() {
	lines := helpers.ReadFile("./day6/big.txt")
	size := len(lines)

	grid := make([][]string, len(lines))
	for row, line := range lines {
		grid[row] = strings.Split(line, "")
	}

	currLoc := Location{0, 0}

	// Find the guard
	for row := range grid {
		for col := range grid[row] {
			c := grid[row][col]
			if c == "^" {
				currLoc = Location{row, col}
				break
			}
		}
	}

	direction := UP
	nextLoc := Location{currLoc[0], currLoc[1]}

	for {
		if direction == UP {
			nextLoc = Location{currLoc[0] - 1, currLoc[1]}
		} else if direction == RIGHT {
			nextLoc = Location{currLoc[0], currLoc[1] + 1}
		} else if direction == DOWN {
			nextLoc = Location{currLoc[0] + 1, currLoc[1]}
		} else if direction == LEFT {
			nextLoc = Location{currLoc[0], currLoc[1] - 1}
		}

		// The guard has exited the grid, so we can exit the loop
		if nextLoc[0] < 0 || nextLoc[0] >= size || nextLoc[1] < 0 || nextLoc[1] >= size {
			grid[currLoc[0]][currLoc[1]] = "X"
			break
		}

		// We've ran into an obstacle and gotta rotate
		if grid[nextLoc[0]][nextLoc[1]] == "#" {
			direction = getRotatedDirection(direction)
			continue
		}

		grid[currLoc[0]][currLoc[1]] = "X"
		currLoc = Location{nextLoc[0], nextLoc[1]}
	}

	total := 0
	for _, row := range grid {
		for _, c := range row {
			if c == "X" {
				total += 1
			}
		}
	}

	fmt.Print("\n\n", "*** total ***", "\n", total, "\n\n\n")

}

func Part2() {
	lines := helpers.ReadFile("./day6/big.txt")
	size := len(lines)

	grid := make([][]string, len(lines))
	for row, line := range lines {
		grid[row] = strings.Split(line, "")
	}

	startLoc := Location{0, 0}

	// Find the guard
	for row := range grid {
		for col := range grid[row] {
			c := grid[row][col]
			if c == "^" {
				startLoc = Location{row, col}
				break
			}
		}
	}

	// Array of all available locations to put an obstacle
	openLocations := []Location{}
	for row := range grid {
		for col := range grid[row] {
			loc := grid[row][col]

			if loc == "." {
				openLocations = append(openLocations, Location{row, col})
			}
		}
	}

	total := 0

	for _, openLoc := range openLocations {
		direction := UP
		currLoc := Location{startLoc[0], startLoc[1]}
		nextLoc := Location{currLoc[0], currLoc[1]}

		// Locations the guard has been to already. The key is a string of:
		// `ROW_COL_DIRECTION`. For example, `4_5_2`. Direction is essentially an
		// enum number. The value is just a bool because it doesn't matter, we
		// essentially just want a set.
		pastLocations := make(map[string]bool)

		// Put an obstacle there
		grid[openLoc[0]][openLoc[1]] = "#"

		loopFound := false

		// run the full loop to see if the guard escapes
		for {
			if direction == UP {
				nextLoc = Location{currLoc[0] - 1, currLoc[1]}
			} else if direction == RIGHT {
				nextLoc = Location{currLoc[0], currLoc[1] + 1}
			} else if direction == DOWN {
				nextLoc = Location{currLoc[0] + 1, currLoc[1]}
			} else if direction == LEFT {
				nextLoc = Location{currLoc[0], currLoc[1] - 1}
			}

			// The guard has exited the grid, so we can exit the loop
			if nextLoc[0] < 0 || nextLoc[0] >= size || nextLoc[1] < 0 || nextLoc[1] >= size {
				break
			}

			// We've ran into an obstacle and gotta rotate
			if grid[nextLoc[0]][nextLoc[1]] == "#" {
				direction = getRotatedDirection(direction)
				continue
			}

			key := fmt.Sprintf("%d_%d_%d", currLoc[0], currLoc[1], direction)
			if pastLocations[key] {
				// Break out of the loop because we've been to the same location with
				// the same direction before, which means we're in a loop
				loopFound = true
				break
			}

			pastLocations[key] = true

			currLoc = Location{nextLoc[0], nextLoc[1]}
		}

		if loopFound {
			total += 1
		}

		// take away the obstacle
		grid[openLoc[0]][openLoc[1]] = "."
	}

	fmt.Print("\n\n", "*** total ***", "\n", total, "\n\n\n")
}
