package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func getSum(curr []int) int {
	sum := 0
	for _, val := range curr {
		sum += val
	}
	return sum
}

func getMaxCalories(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sums := []int{}
	curr := []int{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			sums = append(sums, getSum(curr))
			curr = nil
		} else {
			i, err := strconv.Atoi(text)
			if err != nil {
				return 0, err
			}
			curr = append(curr, i)
		}
	}
	sums = append(sums, getSum(curr))

	sort.Ints(sums)

	l := len(sums)
	return sums[l-3] + sums[l-2] + sums[l-1], nil
}
