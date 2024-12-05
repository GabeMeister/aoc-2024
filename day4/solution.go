package day4

import (
	"GabeMeister/aoc-2024/helpers"
	"fmt"
	"strings"
)

func getHorizontal(grid [][]string) int {
	total := 0

	// Spelled going from left to right, so include all columns except the final 3
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0])-3; col++ {
			x := grid[row][col]
			m := grid[row][col+1]
			a := grid[row][col+2]
			s := grid[row][col+3]

			if x == "X" && m == "M" && a == "A" && s == "S" {
				total += 1
			}
		}
	}

	// Spelled going from right to left, so include all columns except the first 3
	for row := 0; row < len(grid); row++ {
		for col := 3; col < len(grid[0]); col++ {
			x := grid[row][col]
			m := grid[row][col-1]
			a := grid[row][col-2]
			s := grid[row][col-3]

			if x == "X" && m == "M" && a == "A" && s == "S" {
				total += 1
			}
		}
	}

	return total
}

func getVertical(grid [][]string) int {
	total := 0

	// Spelled "downwards", so leave off the bottom 3 rows
	for row := 0; row < len(grid)-3; row++ {
		for col := 0; col < len(grid[0]); col++ {
			x := grid[row][col]
			m := grid[row+1][col]
			a := grid[row+2][col]
			s := grid[row+3][col]

			if x == "X" && m == "M" && a == "A" && s == "S" {
				total += 1
			}
		}
	}

	// Spelled "upwards", so leave off the top 3 rows
	for row := 3; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			x := grid[row][col]
			m := grid[row-1][col]
			a := grid[row-2][col]
			s := grid[row-3][col]

			if x == "X" && m == "M" && a == "A" && s == "S" {
				total += 1
			}
		}
	}

	return total
}

func getDiagonal(grid [][]string) int {
	total := 0

	// Spelled "down and right", so leave off the bottom 3 rows, and right 3 columns
	for row := 0; row < len(grid)-3; row++ {
		for col := 0; col < len(grid[0])-3; col++ {
			x := grid[row][col]
			m := grid[row+1][col+1]
			a := grid[row+2][col+2]
			s := grid[row+3][col+3]

			if x == "X" && m == "M" && a == "A" && s == "S" {
				total += 1
			}
		}
	}

	// Spelled "up and right", so leave off the top 3 rows, and right 3 columns
	for row := 3; row < len(grid); row++ {
		for col := 0; col < len(grid[0])-3; col++ {
			x := grid[row][col]
			m := grid[row-1][col+1]
			a := grid[row-2][col+2]
			s := grid[row-3][col+3]

			if x == "X" && m == "M" && a == "A" && s == "S" {
				total += 1
			}
		}
	}

	// Spelled "down and left", so leave off the bottom 3 rows, and left 3 columns
	for row := 0; row < len(grid)-3; row++ {
		for col := 3; col < len(grid[0]); col++ {
			x := grid[row][col]
			m := grid[row+1][col-1]
			a := grid[row+2][col-2]
			s := grid[row+3][col-3]

			if x == "X" && m == "M" && a == "A" && s == "S" {
				total += 1
			}
		}
	}

	// Spelled "up and left", so leave off the top 3 rows, and left 3 columns
	for row := 3; row < len(grid); row++ {
		for col := 3; col < len(grid[0]); col++ {
			x := grid[row][col]
			m := grid[row-1][col-1]
			a := grid[row-2][col-2]
			s := grid[row-3][col-3]

			if x == "X" && m == "M" && a == "A" && s == "S" {
				total += 1
			}
		}
	}

	return total
}

func Part1() {
	lines := helpers.ReadFile("./day4/big.txt")
	rowCount := len(lines)

	grid := make([][]string, rowCount)

	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	horizontal := getHorizontal(grid)
	fmt.Print("\n\n", "*** horizontal ***", "\n", horizontal, "\n\n\n")

	vertical := getVertical(grid)
	fmt.Print("\n\n", "*** vertical ***", "\n", vertical, "\n\n\n")

	diagonal := getDiagonal(grid)
	fmt.Print("\n\n", "*** diagonal ***", "\n", diagonal, "\n\n\n")

	total := horizontal + vertical + diagonal
	fmt.Print("\n\n", "*** total ***", "\n", total, "\n\n\n")
}

// M.S
// .A.
// M.S
func getFormation1(grid [][]string) int {
	total := 0

	// Because of the diagnoal nature of the spelling, leave off the last two rows
	// and cols
	for row := 0; row < len(grid)-2; row++ {
		for col := 0; col < len(grid[0])-2; col++ {
			m1 := grid[row][col]
			s1 := grid[row][col+2]
			a1 := grid[row+1][col+1]
			m2 := grid[row+2][col]
			s2 := grid[row+2][col+2]

			if m1 == "M" && s1 == "S" && a1 == "A" && m2 == "M" && s2 == "S" {
				total += 1
			}
		}
	}

	return total
}

// S.M
// .A.
// S.M
func getFormation2(grid [][]string) int {
	total := 0

	// Because of the diagnoal nature of the spelling, leave off the last two rows
	// and cols
	for row := 0; row < len(grid)-2; row++ {
		for col := 0; col < len(grid[0])-2; col++ {
			s1 := grid[row][col]
			m1 := grid[row][col+2]
			a1 := grid[row+1][col+1]
			s2 := grid[row+2][col]
			m2 := grid[row+2][col+2]

			if m1 == "M" && s1 == "S" && a1 == "A" && m2 == "M" && s2 == "S" {
				total += 1
			}
		}
	}

	return total

}

// M.M
// .A.
// S.S
func getFormation3(grid [][]string) int {
	total := 0

	// Because of the diagnoal nature of the spelling, leave off the last two rows
	// and cols
	for row := 0; row < len(grid)-2; row++ {
		for col := 0; col < len(grid[0])-2; col++ {
			m1 := grid[row][col]
			m2 := grid[row][col+2]
			a1 := grid[row+1][col+1]
			s1 := grid[row+2][col]
			s2 := grid[row+2][col+2]

			if m1 == "M" && s1 == "S" && a1 == "A" && m2 == "M" && s2 == "S" {
				total += 1
			}
		}
	}

	return total

}

// S.S
// .A.
// M.M
func getFormation4(grid [][]string) int {
	total := 0

	// Because of the diagnoal nature of the spelling, leave off the last two rows
	// and cols
	for row := 0; row < len(grid)-2; row++ {
		for col := 0; col < len(grid[0])-2; col++ {
			s1 := grid[row][col]
			s2 := grid[row][col+2]
			a1 := grid[row+1][col+1]
			m1 := grid[row+2][col]
			m2 := grid[row+2][col+2]

			if m1 == "M" && s1 == "S" && a1 == "A" && m2 == "M" && s2 == "S" {
				total += 1
			}
		}
	}

	return total

}

func Part2() {
	lines := helpers.ReadFile("./day4/big.txt")
	rowCount := len(lines)

	grid := make([][]string, rowCount)

	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	f1 := getFormation1(grid)
	fmt.Print("\n\n", "*** f1 ***", "\n", f1, "\n\n\n")
	f2 := getFormation2(grid)
	fmt.Print("\n\n", "*** f2 ***", "\n", f2, "\n\n\n")
	f3 := getFormation3(grid)
	fmt.Print("\n\n", "*** f3 ***", "\n", f3, "\n\n\n")
	f4 := getFormation4(grid)
	fmt.Print("\n\n", "*** f4 ***", "\n", f4, "\n\n\n")

	total := f1 + f2 + f3 + f4
	fmt.Print("\n\n", "*** total ***", "\n", total, "\n\n\n")
}
