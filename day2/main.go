package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return // exit
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safe := 0
	for scanner.Scan() {
		line := scanner.Text()
		safe += computeRow(line)
	}
	fmt.Println(safe)
}

func computeRow(line string) int {
	elements := strings.Split(line, " ")
	safeLine := true
	firstPos, _ := strconv.Atoi(elements[0])
	secondPos, _ := strconv.Atoi(elements[1])
	direction := firstPos-secondPos > 0

	for i := 1; i < len(elements); i++ {
		prev, _ := strconv.Atoi(elements[i-1])
		curr, _ := strconv.Atoi(elements[i])

		diff := math.Abs(float64(prev - curr))
		if (diff == 0 || diff > 3) || ((prev-curr > 0) != direction) {
			safeLine = false
			break
		}
	}

	if safeLine {
		return 1
	} else {
		return 0
	}
}
