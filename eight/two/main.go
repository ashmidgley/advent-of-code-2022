package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid, err := parseInput(file)
	if err != nil {
		panic(err)
	}

	score, err := calculateMaxScenicScore(grid)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Score: %d\n", score)
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

func calculateMaxScenicScore(grid [][]int) (int, error) {
	max := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			curr := grid[i][j]
			left := countVisibleLeft(grid[i][:j], curr)
			right := countVisibleRight(grid[i][j+1:], curr)
			top := countVisibleTop(grid, i, j)
			bottom := countVisibleBottom(grid, i, j)
			score := left * right * top * bottom
			if score > max {
				max = score
			}
		}
	}
	return max, nil
}

func countVisibleLeft(arr []int, curr int) int {
	count := 0
	for i := len(arr) - 1; i >= 0; i-- {
		count++
		if arr[i] >= curr {
			break
		}
	}
	return count
}

func countVisibleRight(arr []int, curr int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		count++
		if arr[i] >= curr {
			break
		}
	}
	return count
}

func countVisibleTop(grid [][]int, y, x int) int {
	curr := grid[y][x]
	count := 0
	for i := y - 1; i >= 0; i-- {
		count++
		if grid[i][x] >= curr {
			break
		}
	}
	return count
}

func countVisibleBottom(grid [][]int, y, x int) int {
	curr := grid[y][x]
	count := 0
	for i := y + 1; i < len(grid); i++ {
		count++
		if grid[i][x] >= curr {
			break
		}
	}
	return count
}
