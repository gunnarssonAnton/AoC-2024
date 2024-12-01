package day01

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func RunDay01() {
	//inputFile, err := os.ReadFile("day01/example.txt")
	inputFile, err := os.ReadFile("day01/input.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	rows := strings.Split(string(inputFile), "\n")

	var leftColumn []int
	var rightColumn []int

	for _, row := range rows {
		columns := strings.Fields(row)

		if len(columns) == 2 {
			leftValue, err := strconv.Atoi(columns[0])
			if err != nil {
				fmt.Println("Error converting left column value:", err)
				continue
			}
			rightValue, err := strconv.Atoi(columns[1])
			if err != nil {
				fmt.Println("Error converting right column value:", err)
				continue
			}
			leftColumn = append(leftColumn, leftValue)
			rightColumn = append(rightColumn, rightValue)
		}
	}
	sort.Ints(leftColumn)
	sort.Ints(rightColumn)
	//part1(rightColumn, leftColumn)
	part2(rightColumn, leftColumn)

}

func countOccur(slice []int, target int) int {
	count := 0
	for _, num := range slice {
		if num == target {
			count++
		}
	}

	return count
}

func part1(rightColumn []int, leftColumn []int) {
	fmt.Println("###### DAY 1 Part 1 ######")

	if len(rightColumn) != len(leftColumn) {
		fmt.Println("ERROR, not the same lenght")
		return
	}
	var sum int = 0
	for i := 0; i < len(leftColumn); i++ {
		leftvalue := leftColumn[i]
		rightValue := rightColumn[i]

		if leftvalue > rightValue {
			sum += (leftvalue - rightValue)
		} else {
			sum += (rightValue - leftvalue)
		}

	}
	fmt.Println(sum)
}

func part2(rightColumn []int, leftColumn []int) {
	fmt.Println("###### DAY 1 Part 2 ######")

	if len(rightColumn) != len(leftColumn) {
		fmt.Println("ERROR, not the same lenght")
		return
	}
	var sum int = 0
	for i := 0; i < len(leftColumn); i++ {
		leftvalue := leftColumn[i]
		occur := countOccur(rightColumn, leftvalue)
		sum += (leftvalue * occur)
	}

	fmt.Println(sum)
}
