package main

import (
	"bufio"
	"fmt"
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
	rules := computeOrderingRules(scanner)
	pagesList := computeOrderingPages(scanner)

	var correctPages [][]int
	for _, pages := range pagesList {
		if isPageCorrect(pages, rules) {
			correctPages = append(correctPages, pages)
		}
	}
	fmt.Println(computeFinalSum(correctPages))
}

func isPageCorrect(pages []int, rules map[int][]int) bool {
	for i := len(pages) - 1; i >= 0; i-- {
		page := pages[i]
		valuesAfter, exists := rules[page]
		if exists == false {
			continue
		}
		subarray := pages[:i]
		for _, value := range valuesAfter {
			if containsValue(subarray, value) {
				return false
			}
		}
	}
	return true
}

func computeFinalSum(correctPages [][]int) int {
	total := 0
	for _, page := range correctPages {
		middle := len(page) / 2
		total += page[middle]
	}
	return total
}

func computeOrderingPages(scanner *bufio.Scanner) [][]int {
	var pages [][]int
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, ",")
		pages = append(pages, convertStringsToInts(elements))
	}
	return pages
}

func computeOrderingRules(scanner *bufio.Scanner) map[int][]int {
	rules := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			return rules
		}
		elements := strings.Split(line, "|")
		firstPos, _ := strconv.Atoi(elements[0])
		secondPos, _ := strconv.Atoi(elements[1])
		rules[firstPos] = append(rules[firstPos], secondPos)
	}
	return rules
}

func containsValue(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func convertStringsToInts(strings []string) []int {
	ints := make([]int, len(strings))

	for i, s := range strings {
		num, _ := strconv.Atoi(s)
		ints[i] = num
	}

	return ints
}
