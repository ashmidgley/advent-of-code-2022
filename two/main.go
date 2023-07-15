package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
	LOSS     = 0
	DRAW     = 3
	WIN      = 6
	A        = "A"
	B        = "B"
	C        = "C"
)

func getResultScore(result string) int {
	if result == "X" {
		return LOSS
	} else if result == "Y" {
		return DRAW
	}
	return WIN
}

func getRequiredMoveScore(oppMove string, result int) int {
	if result == LOSS {
		if oppMove == A {
			return SCISSORS
		} else if oppMove == B {
			return ROCK
		} else {
			return PAPER
		}
	}

	if result == DRAW {
		if oppMove == A {
			return ROCK
		} else if oppMove == B {
			return PAPER
		} else {
			return SCISSORS
		}
	}

	if oppMove == A {
		return PAPER
	} else if oppMove == B {
		return SCISSORS
	} else {
		return ROCK
	}
}

func calculateTotal(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		moves := strings.Split(text, " ")
		oppMove := moves[0]
		result := getResultScore(moves[1])
		score := getRequiredMoveScore(oppMove, result) + result
		total += score
	}

	return total, nil
}

func main() {
	total, err := calculateTotal("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Total: %d", total)
}
