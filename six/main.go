package main

import (
	"bufio"
	"fmt"
	"os"
)

func getStartOfPacketMarkerIndex(stream string, min int) int {
	items := []rune{}
	index := 0
	for _, v := range stream {
		if len(items) < min {
			items = append(items, v)
		} else {
			items = append(items[1:], v)
			set := map[rune]bool{}
			for _, v := range items {
				if _, ok := set[v]; !ok {
					set[v] = true
				}
			}

			if len(set) == min {
				return index + 1
			}
		}
		index++
	}

	return index
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	stream := scanner.Text()

	result := getStartOfPacketMarkerIndex(stream, 14)
	fmt.Println(result)
}
