package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coor struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return // exit
	}
	defer file.Close()

	textMap := loadTextMap(file)

	findXmas(textMap)
}

func loadTextMap(file *os.File) [][]string {
	scanner := bufio.NewScanner(file)
	var textMap [][]string
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		textMap = append(textMap, row)
	}
	return textMap
}

func findXmas(textMap [][]string) {
	total := 0
	for y := 0; y < len(textMap); y++ {
		for x := 0; x < len(textMap[y]); x++ {
			if textMap[y][x] == "X" {
				total += findRestOfXmax(textMap, x, y, [3]string{"M", "A", "S"})
			}
		}
	}
	fmt.Println(total)
}

func findRestOfXmax(textMap [][]string, initX int, initY int, letters [3]string) int {
	coors := getNextLetterPositions(textMap, initX, initY, letters[0])
	total := 0
	for _, nextLetterCoor := range coors {
		dx := nextLetterCoor.x - initX
		dy := nextLetterCoor.y - initY
		nextStepDirection := Coor{dx, dy}
		nextX := nextLetterCoor.x + nextStepDirection.x
		nextY := nextLetterCoor.y + nextStepDirection.y
		isThere := isLetterInPosition(textMap, nextX, nextY, letters[1])
		if isThere {
			nextOfTheNextX := nextLetterCoor.x + (nextStepDirection.x * 2)
			nextOfTheNextY := nextLetterCoor.y + (nextStepDirection.y * 2)
			andThere := isLetterInPosition(textMap, nextOfTheNextX, nextOfTheNextY, letters[2])
			if andThere {
				total += 1
			}
		}
	}
	return total
}

func getNextLetterPositions(textMap [][]string, initX int, initY int, nextLetter string) []Coor {
	var coors []Coor
	for y := initY - 1; y <= initY+1; y++ {
		if y >= 0 && y < len(textMap) {
			for x := initX - 1; x <= initX+1; x++ {
				if x >= 0 && x < len(textMap[y]) {
					if textMap[y][x] == nextLetter {
						coors = append(coors, Coor{x, y})
					}
				}
			}
		}
	}
	return coors
}

func isLetterInPosition(textMap [][]string, x int, y int, letter string) bool {
	if x >= 0 && x < len(textMap[0]) && y >= 0 && y < len(textMap) {
		if textMap[y][x] == letter {
			return true
		}
	}
	return false
}
