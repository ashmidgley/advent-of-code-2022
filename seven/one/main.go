package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	parent   *node
	children map[string]*node
	name     string
	size     int
	isFile   bool
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var current *node
	folders := []*node{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		text := sc.Text()
		split := strings.Split(text, " ")
		if len(split) > 2 {
			path := split[2]
			if path == ".." {
				current = current.parent
			} else if path == "/" {
				current = &node{
					name:     "/",
					size:     0,
					isFile:   false,
					children: make(map[string]*node),
					parent:   nil,
				}
			} else {
				current = current.children[path]
			}
		} else if split[0] == "dir" {
			folderName := split[1]
			child := &node{
				name:     folderName,
				size:     0,
				isFile:   false,
				children: make(map[string]*node),
				parent:   current,
			}
			current.children[folderName] = child
			folders = append(folders, current.children[folderName])
		} else if split[0] != "$" {
			fileName := split[1]
			size, _ := strconv.Atoi(split[0])
			child := &node{
				name:     fileName,
				size:     size,
				isFile:   true,
				children: nil,
				parent:   current,
			}
			current.children[fileName] = child
		}
	}

	sum := 0
	for _, folder := range folders {
		size := getSize(*folder)
		if size <= 100000 {
			sum += size
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}

func getSize(root node) (size int) {
	if root.isFile {
		return root.size
	}

	for _, child := range root.children {
		size += getSize(*child)
	}
	return
}
