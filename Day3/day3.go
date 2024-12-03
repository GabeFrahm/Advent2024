package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getInstructions() []string {
	f, err := os.Open("../input/day3")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scn := bufio.NewScanner(f)

	pattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	
	instructionList := []string{}
	for scn.Scan() {
		line := pattern.FindAllString(scn.Text(), -1)
		if line == nil {
			continue
		}
		instructionList = append(instructionList, line...)
	}

	return instructionList
}

func part1() int {
	instructions := getInstructions()

	sum := 0

	for _, instr := range instructions {
		if instr[0:3] != "mul" { continue } // ignore `do()` and `dont()` for part 1
		strSlice := strings.Split(instr[4:len(instr)-1], ",") // extracting arguments from function
		n1, _ := strconv.Atoi(strSlice[0])
		n2, _ := strconv.Atoi(strSlice[1])

		sum += n1 * n2
	}

	return sum
}

func part2() int {
	instructions := getInstructions()

	sum := 0
	isEnabled := true

	for _, instr := range instructions {
		fn, _, _ := strings.Cut(instr, "(")
		switch fn {
		case "do":
			isEnabled = true
		case "don't":
			isEnabled = false
		case "mul":
			if !isEnabled {continue}
			strSlice := strings.Split(instr[4:len(instr)-1], ",") // extracting arguments from function
			n1, _ := strconv.Atoi(strSlice[0])
			n2, _ := strconv.Atoi(strSlice[1])

			sum += n1 * n2
		}
	}

	return sum
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}