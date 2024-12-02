package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return // exit
	}
	defer file.Close()

	totalDiff := 0
	column1, column2 := getColumns(file)
	sort.Ints(column1)
	sort.Ints(column2)

	for i := 0; i < len(column1); i++ {
		totalDiff += int(math.Abs(float64(column1[i] - column2[i])))
	}
	fmt.Println(totalDiff)
}

func getColumns(file *os.File) ([]int, []int) {
	var column1 []int
	var column2 []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, "   ")
		value1, err1 := strconv.Atoi(columns[0])
		value2, err2 := strconv.Atoi(columns[1])
		if err1 == nil && err2 == nil {
			column1 = append(column1, value1)
			column2 = append(column2, value2)
		}
	}

	return column1, column2
}
