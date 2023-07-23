package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func processStacks(stackLines []string) ([][]string, error) {
	finalLine := stackLines[len(stackLines)-1]
	stackCount, err := strconv.Atoi(strings.TrimSpace(finalLine[len(finalLine)-2:]))
	if err != nil {
		return nil, err
	}

	stacks := make([][]string, stackCount)
	height := len(stackLines) - 2
	for i := 0; i <= height; i++ {
		line := stackLines[i]

		j := 0
		for k := 0; k < len(line)-2; k++ {
			if k == 0 || k%4 == 0 {
				if string(line[k]) == "[" {
					stacks[j] = append(stacks[j], string(line[k+1]))
				}
				j++
			}
		}
	}
	return stacks, err
}

type step struct {
	move int
	from int
	to   int
}

func processSteps(stepLines []string) []step {
	steps := make([]step, len(stepLines))
	for i, line := range stepLines {
		split := strings.Split(line, " ")
		move, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		from--
		to, _ := strconv.Atoi(split[5])
		to--
		steps[i] = step{move, from, to}
	}
	return steps
}

func readFile(fileName string) ([][]string, []step, error) {
	stackLines := []string{}
	stepLines := []string{}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	stacksScanned := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			stacksScanned = true
			continue
		}
		if stacksScanned {
			stepLines = append(stepLines, text)
		} else {
			stackLines = append(stackLines, text)
		}
	}

	stacks, err := processStacks(stackLines)
	steps := processSteps(stepLines)
	return stacks, steps, err
}

func getStackTops(fileName string) (string, error) {
	stacks, steps, err := readFile(fileName)
	if err != nil {
		return "", err
	}

	for _, step := range steps {
		items := make([]string, step.move)
		copy(items, stacks[step.from][:step.move])
		stacks[step.from] = stacks[step.from][step.move:]
		stacks[step.to] = append(items, stacks[step.to]...)
	}

	var builder strings.Builder
	for _, stack := range stacks {
		builder.WriteString(stack[0])
	}

	return builder.String(), nil
}
