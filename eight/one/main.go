package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid, err := parseInput(file)
	if err != nil {
		panic(err)
	}

	count, err := calculateVisibleTrees(grid)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Count: %d\n", count)
}

func parseInput(file *os.File) ([][]int, error) {
	grid := [][]int{}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, "")
		heights, err := parseInts(split)
		if err != nil {
			return nil, err
		}
		grid = append(grid, heights)
		i++
	}
	return grid, nil
}

func parseInts(arr []string) ([]int, error) {
	var result []int
	for _, val := range arr {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

func calculateVisibleTrees(grid [][]int) (int, error) {
	count := (len(grid) - 1) * 4
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			curr := grid[i][j]
			if isVisibleHorizontal(grid[i][:j], curr) || isVisibleHorizontal(grid[i][j+1:], curr) || isVisibleTop(grid, i, j) || isVisibleBottom(grid, i, j) {
				count++
			}
		}
	}

	return count, nil
}

func isVisibleHorizontal(arr []int, curr int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] >= curr {
			return false
		}
	}
	return true
}

func isVisibleTop(grid [][]int, y, x int) bool {
	curr := grid[y][x]
	for i := y - 1; i >= 0; i-- {
		if grid[i][x] >= curr {
			return false
		}
	}
	return true
}

func isVisibleBottom(grid [][]int, y, x int) bool {
	curr := grid[y][x]
	for i := y + 1; i < len(grid); i++ {
		if grid[i][x] >= curr {
			return false
		}
	}
	return true
}
