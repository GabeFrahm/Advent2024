package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getLists() ([]int, []int) {
	f, err := os.Open("../input/day1")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	l1 := []int{}
	l2 := []int{}

	scanner := bufio.NewScanner(f)

	// read ints into lists
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "   ")

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		l1 = append(l1, num1)
		l2 = append(l2, num2)
	}

	// sort lists
	slices.Sort(l1)
	slices.Sort(l2)

	return l1, l2
}

func part1(l1 []int, l2 []int) int {
	sum := 0
	for i := range l1 {
		diff := l1[i] - l2[i]
		if diff < 0 {
			diff = -diff
		}

		sum += diff
	}

	return sum
}

func part2(l1 []int, l2 []int) int {
	sum := 0
	
	for _, n1 := range l1 {
		count := 0
		for _, n2 := range l2 {
			if n2 == n1 { count++ }
		}

		sum += n1 * count
	}

	return sum
}

func main() {
	list1, list2 := getLists()

	part1 := part1(list1, list2)
	fmt.Println("Part 1:", part1)

	part2 := part2(list1, list2)
	fmt.Println("Part 2:", part2)
}