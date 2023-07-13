package main

import (
	"bufio"
	"fmt"
	"os"
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
	curr := []int{}
	max := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			sum := getSum(curr)
			if sum > max {
				max = sum
			}
			curr = nil
		} else {
			i, err := strconv.Atoi(text)
			if err != nil {
				return 0, err
			}
			curr = append(curr, i)
		}
	}

	sum := getSum(curr)
	if sum > max {
		max = sum
	}

	return max, nil
}

func main() {
	max, elf := getMaxCalories("input.txt")
	fmt.Printf("Max: %d\n", max)
	fmt.Printf("Elf: %d\n", elf)
}
