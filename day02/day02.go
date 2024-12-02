package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafeLine(arr []int, maxDiff int) bool {
	if isIncreasingWithMaxDiff(arr, maxDiff) || isDecreasingWithMaxDiff(arr, maxDiff) {
		return true
	} else {
		return false
	}
}

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
	inputFile, err := os.ReadFile("day02/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	rows := strings.Split(string(inputFile), "\n")
	//fmt.Println(inputFile)
	part1(rows)
	part2(rows)
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

		if isSafeLine(intArray, 3) {
			count++
		}

	}
	fmt.Printf("PART 1: %d\n", count)

}
func part2(rows []string) {
	fmt.Println("###### DAY 2 Part 2 ######")
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

		if isSafeLine(intArray, 3) {
			count++
		} else {
			for index := 1; index < len(intArray); index++ {

				intArray = append(intArray[:index], intArray[index+1:]...)

				if isSafeLine(intArray, 3) {
					count++
					break
				}
			}

		}

	}
	fmt.Printf("PART 2: %d", count)
}
