package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func isOverlap(e1Min, e1Max, e2Min, e2Max int) bool {
	// Check which range is bigger.
	if (e1Max-e1Min)-(e2Max-e2Min) > 0 {
		// r1 larger.
		if e2Min <= e1Max && e2Max >= e1Min {
			return true
		}
	} else {
		if e1Min <= e2Max && e1Max >= e2Min {
			return true
		}
	}
	return false
}

func getRangeOverlapCount(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		text := scanner.Text()
		elvesSplit := strings.Split(text, ",")
		e1Split := strings.Split(elvesSplit[0], "-")
		e1Min, _ := strconv.Atoi(e1Split[0])
		e1Max, _ := strconv.Atoi(e1Split[1])
		e2Split := strings.Split(elvesSplit[1], "-")
		e2Min, _ := strconv.Atoi(e2Split[0])
		e2Max, _ := strconv.Atoi(e2Split[1])

		if isOverlap(e1Min, e1Max, e2Min, e2Max) {
			count++
		}
	}

	return count, nil
}
