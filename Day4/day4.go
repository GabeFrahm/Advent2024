package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

type Coords struct {
	row int
	col int
}

func getCrossword() [][]rune {
	f, err := os.Open("../input/day4")
	if err != nil {
		log.Fatal(err)
	}
	scn := bufio.NewScanner(f)

	crossword := [][]rune{}
	for scn.Scan() {
		crossword = append(crossword, []rune(scn.Text()))
	}

	f.Close()

	return crossword
}


/*
	rowDir and colDir: 1 increasing, 0 constant, -1 decreasing
	returns 1 on a match and 0 otherwise
*/
func check(crossword [][]rune, start Coords, rowDir int, colDir int) int {
	r := start.row
	c := start.col

	if
	crossword[r             ][c             ] == 'X' &&
	crossword[r + (1*rowDir)][c + (1*colDir)] == 'M' &&
	crossword[r + (2*rowDir)][c + (2*colDir)] == 'A' &&
	crossword[r + (3*rowDir)][c + (3*colDir)] == 'S' {
		return 1
	}
	return 0
}

// ALIAS FUNCTIONS FOR READABILITY
func checkDiagUpLeft(crossword [][]rune, start Coords) int {
	return check(crossword, start, -1, -1)
}
func checkDiagUpRight(crossword [][]rune, start Coords) int {
	return check(crossword, start, -1, +1)
}
func checkDiagDownLeft(crossword [][]rune, start Coords) int {
	return check(crossword, start, +1, -1)
}
func checkDiagDownRight(crossword [][]rune, start Coords) int {
	return check(crossword, start, +1, +1)
}
func checkUp(crossword [][]rune, start Coords) int {
	return check(crossword, start, -1, 0)
}
func checkDown(crossword [][]rune, start Coords) int {
	return check(crossword, start, +1, 0)
}
func checkLeft(crossword [][]rune, start Coords) int {
	return check(crossword, start, 0, -1)
}
func checkRight(crossword [][]rune, start Coords) int {
	return check(crossword, start, 0, +1)
}


func part1(crossword [][]rune) int {
	matches := 0

	for r, row := range crossword {
		for c, letter := range row {
			if letter != 'X' {continue}

			coords := Coords{r, c}

			// Cardinal directions
			if c >= 3                  {matches += checkLeft(crossword, coords)}
			if c <= len(row) - 4       {matches += checkRight(crossword, coords)}
			if r >= 3 				   {matches += checkUp(crossword, coords)}
			if r <= len(crossword) - 4 {matches += checkDown(crossword, coords)}

			// Diags
			if r >= 3 && c >= 3 {matches += checkDiagUpLeft(crossword, coords)}
			if r >= 3 && c <= len(row) - 4 {matches += checkDiagUpRight(crossword, coords)}
			if r <= len(row) - 4 && c >= 3 {matches += checkDiagDownLeft(crossword, coords)}
			if r <= len(row) - 4 && c <= len(crossword) - 4 {
				matches += checkDiagDownRight(crossword, coords)
			}
		}
	}

	return matches
}

func xMas(crossword [][]rune, start Coords) int {
	// immediately invalid if there isn't an 'A' in the middle
	if crossword[start.row+1][start.col+1] != 'A' {
		return 0
	}

	topLeft := crossword[start.row][start.col]
	topRight := crossword[start.row][start.col + 2]
	bottomLeft := crossword[start.row+2][start.col]
	bottomRight := crossword[start.row+2][start.col+2]

	isDownLeftValid := false
	isDownRightValid := false

	if topLeft == 'M' && bottomRight == 'S' || 
	   topLeft == 'S' && bottomRight == 'M' {
		isDownLeftValid = true
	}
	if topRight == 'M' && bottomLeft == 'S' ||
	   topRight == 'S' && bottomLeft == 'M' {
		isDownRightValid = true
	}

	if isDownLeftValid && isDownRightValid {
		return 1
	}
	return 0
}

func part2(crossword [][]rune) int {
	matches := 0

	for r, row := range crossword {
		for c, letter := range row {
			if letter == 'X' {continue}

			if r <= len(crossword) - 3 && c <= len(row) - 3 {
				matches += xMas(crossword, Coords{r, c})
			}
		}
	}

	return matches
}

func main() {
	crossword := getCrossword()
	fmt.Println("Part 1:", part1(crossword))
	fmt.Println("Part 2:", part2(crossword))
}