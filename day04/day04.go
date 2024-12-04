package day04

import (
	"fmt"
	"os"
	"strings"
)

func RunDay04() {
	fmt.Println("~~~~~~ [ DAY 4 ] ~~~~~~")
	inputFile, err := os.ReadFile("day04/input.txt")
	//inputFile, err := os.ReadFile("day04/example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	matrix := getRuneMatrix(string(inputFile))
	part1(matrix)
	part2(matrix)
}

func getRuneMatrix(str string) [][]rune {
	lines := strings.Split(strings.TrimSpace(str), "\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	return matrix
}

func part1(grid [][]rune) {
	fmt.Println("###### DAY 4 Part 1 ######")

	tot := 0
	target := []rune{'X', 'M', 'A', 'S'}

	checkDirection := func(y, x, dy, dx int) bool {
		for i := 0; i < 4; i++ {
			ny := y + dy*i
			nx := x + dx*i
			if ny < 0 || ny >= len(grid) || nx < 0 || nx >= len(grid[ny]) || grid[ny][nx] != target[i] {
				return false
			}
		}
		return true
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			directions := [][2]int{
				{0, 1},
				{0, -1},
				{1, 0},
				{-1, 0},
				{1, 1},
				{-1, 1},
				{1, -1},
				{-1, -1},
			}

			for _, dir := range directions {
				if checkDirection(y, x, dir[0], dir[1]) {
					tot++
				}
			}

		}
	}
	fmt.Println(tot)
}

// 2545
func getPatterns() [][][]rune {
	patterns := [][][]rune{
		{
			{'M', '.', 'S'},
			{'.', 'A', '.'},
			{'M', '.', 'S'},
		},
		{
			{'M', '.', 'M'},
			{'.', 'A', '.'},
			{'S', '.', 'S'},
		},
		{
			{'S', '.', 'M'},
			{'.', 'A', '.'},
			{'S', '.', 'M'},
		},
		{
			{'S', '.', 'S'},
			{'.', 'A', '.'},
			{'M', '.', 'M'},
		},
	}
	return patterns
}
func part2(grid [][]rune) {
	fmt.Println("###### DAY 4 Part 2 ######")
	//patterns := getPatterns()

	tot := 0
	for y := 0; y < len(grid)-2; y++ {
		for x := 0; x < len(grid[0])-2; x++ {
			if grid[y][x] == 'M' && grid[y][x+2] == 'S' && grid[y+1][x+1] == 'A' && grid[y+2][x] == 'M' && grid[y+2][x+2] == 'S' {
				tot++
			}
			if grid[y][x] == 'M' && grid[y][x+2] == 'M' && grid[y+1][x+1] == 'A' && grid[y+2][x] == 'S' && grid[y+2][x+2] == 'S' {
				tot++
			}

			if grid[y][x] == 'S' && grid[y][x+2] == 'M' && grid[y+1][x+1] == 'A' && grid[y+2][x] == 'S' && grid[y+2][x+2] == 'M' {
				tot++
			}

			if grid[y][x] == 'S' && grid[y][x+2] == 'S' && grid[y+1][x+1] == 'A' && grid[y+2][x] == 'M' && grid[y+2][x+2] == 'M' {
				tot++
			}

		}

	}
	fmt.Println(tot)
}
