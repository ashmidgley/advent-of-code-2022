package main

import (
	"bufio"
	"fmt"
	"os"
)

func getPriority(r rune) rune {
	if 'a' <= r && r <= 'z' {
		return r - 96
	}
	return r - 38
}

func getPrioritySum(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matches := []rune{}
	index := 0
	group := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		group = append(group, text)
		if index == 2 {
			r1 := group[0]
			r2 := group[1]
			r3 := group[2]

			s1 := map[rune]bool{}
			for _, val := range r1 {
				if _, ok := s1[val]; !ok {
					s1[val] = true
				}
			}

			s2 := map[rune]bool{}
			for _, val := range r2 {
				if _, ok := s1[val]; ok {
					s2[val] = true
				}
			}

			for _, val := range r3 {
				if _, ok := s2[val]; ok {
					matches = append(matches, val)
					break
				}
			}

			group = nil
			index = 0
		} else {
			index++
		}
	}

	total := 0
	for _, val := range matches {
		total += int(getPriority(val))
	}

	return total, nil
}

func main() {
	sum, err := getPrioritySum("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Priority Sum: %d\n", sum)
}
