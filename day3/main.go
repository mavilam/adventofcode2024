package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	file, err := os.Open("input.txt")
	if err != nil {
		return // exit
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += computeRow(line, re)
	}
	fmt.Println(sum)
}

func computeRow(line string, re *regexp.Regexp) int {
	allMatches := re.FindAllString(line, -1)
	sum := 0
	for i := 0; i < len(allMatches); i++ {
		parts := strings.Split(allMatches[i], ",")
		firstPos, _ := strconv.Atoi(strings.Replace(parts[0], "mul(", "", -1))
		secondPos, _ := strconv.Atoi(strings.Replace(parts[1], ")", "", -1))
		sum += firstPos * secondPos
	}
	return sum
}
