package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isIncreasingWithMaxDiff(arr []int, maxdiff int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] || arr[i]-arr[i-1] > maxdiff || arr[i] == arr[i-1] {
			return false
		}
	}
	return true
}

func isDecreasingWithMaxDiff(arr []int, maxdiff int) bool {

	for i := 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] || arr[i-1]-arr[i] > maxdiff || arr[i-1] == arr[i] {
			return false
		}

	}
	return true
}

func RunDay02() {
	inputFile, err := os.ReadFile("day02/example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	rows := strings.Split(string(inputFile), "\n")
	//fmt.Println(inputFile)
	part1(rows)
	//part2()
}
func part1(rows []string) {
	fmt.Println("###### DAY 2 Part 1 ######")

	count := 0
	for _, row := range rows {
		columns := strings.Fields(row)
		intArray := make([]int, len(columns))

		for i, str := range columns {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting %s to integer: %v\n", str, err)
				return
			}
			intArray[i] = num
		}

		if isIncreasingWithMaxDiff(intArray, 3) || isDecreasingWithMaxDiff(intArray, 3) {
			count++
		}

	}

	fmt.Println(count)
}
func part2() {
	fmt.Println("###### DAY 2 Part 2 ######")
}
