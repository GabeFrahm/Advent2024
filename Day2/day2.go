package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
	returns true if the report passes.
*/
func reportMarch(lvls []int) bool {
	if lvls[1] > lvls[0] {
		// INCREASING
		for i := range lvls {
			if i != len(lvls) - 1 {
				next := lvls[i+1]
				this := lvls[i]

				if !(next > this && (next - this) <= 3) {
					return false
				}
			} else {
				return true
			}
		}
	} else {
		// DECREASING
		for i := range lvls {
			if i != len(lvls) - 1 {
				next := lvls[i+1]
				this := lvls[i]

				if !(next < this && (this - next) <= 3) {
					return false
				}
			} else {
				return true
			}
		}
	}

	return false // will never reach this
}

func part1() int {
	f, err := os.Open("../input/day2")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scn := bufio.NewScanner(f)

	safe := 0

	for scn.Scan() {
		stringLvls := strings.Split(scn.Text(), " ");
		lvls := []int{};
		for _, e := range stringLvls {
			lvl, _ := strconv.Atoi(e)
			lvls = append(lvls, lvl)
		}

		if reportMarch(lvls) {
			safe++
		}
	}

	return safe
}

func part2() int {
	f, err := os.Open("../input/day2")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scn := bufio.NewScanner(f)

	safe := 0

	for scn.Scan() {
		stringLvls := strings.Split(scn.Text(), " ");
		lvls := []int{};
		for _, e := range stringLvls {
			lvl, _ := strconv.Atoi(e)
			lvls = append(lvls, lvl)
		}

		if reportMarch(lvls) {
			safe++
			continue
		}

		// fail testing
		// I tried a more sophisticated approach, that involved reporting
		// the specific index where the report failed, but it required extra
		// work outside of the function to determine whether the failed index
		// or the next index was the actual problem, and even then my method
		// couldn't find all the cases and I was on the clock...
		// may revisit this for a better approach that isn't brute-force
		for i := range lvls {
			lvlsn := make([]int, len(lvls))
			copy(lvlsn, lvls)
			if reportMarch(append(lvlsn[:i], lvlsn[i+1:]...)) {
				safe++
				break
			}
		}
	}

	return safe
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}