package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func RunDay03() {
	inputFile, err := os.ReadFile("day03/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	//part1(string(inputFile))
	part2(string(inputFile))

}

func part1(file string) {
	fmt.Println("###### DAY 3 Part 1 ######")
	pattern := `mul\(\d+,\d+\)`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(file, -1)
	var sum = 0
	for _, str := range matches {
		numPattern := `\d+(\.\d+)?`
		r := regexp.MustCompile(numPattern)
		m := r.FindAllString(str, -1)
		//fmt.Println(m)
		num1, err1 := strconv.Atoi(m[0])
		num2, err2 := strconv.Atoi(m[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting numbers:", m[0], m[1])
			continue
		}
		sum += (num1 * num2)
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func part2(input string) {

	fmt.Println("###### DAY 3 Part 2 ######")
	pattrn := `mul\(\d+,\d+\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattrn)
	matches := re.FindAllString(input, -1)

	g := true
	sum := 0
	for _, str := range matches {
		if str == "do()" {
			g = true
		} else if str == "don't()" {
			g = false
		} else {
			if g {
				numPattern := `\d+(\.\d+)?`
				r := regexp.MustCompile(numPattern)
				m := r.FindAllString(str, -1)
				//fmt.Println(m)
				num1, err1 := strconv.Atoi(m[0])
				num2, err2 := strconv.Atoi(m[1])
				if err1 != nil || err2 != nil {
					fmt.Println("Error converting numbers:", m[0], m[1])
					continue
				}
				sum += (num1 * num2)
			}
		}

	}
	fmt.Printf("Part 2: %d,\n", sum)

}
