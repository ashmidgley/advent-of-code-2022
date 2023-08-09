package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name     string
	size     int
	isFile   bool
	children map[string]*node
	parent   *node
}

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	var current *node
	folders := []*node{}
	sc := bufio.NewScanner(input)
	for sc.Scan() {
		text := sc.Text()
		split := strings.Split(text, " ")
		if len(split) > 2 {
			path := split[2]
			if path == ".." {
				current = current.parent
			} else if path == "/" {
				current = &node{"/", 0, false, make(map[string]*node), nil}
				folders = append(folders, current)
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
			size, _ := strconv.Atoi(split[0])
			fileName := split[1]
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

	toFree := 30000000 - (70000000 - getSize(*folders[0]))
	smallest := getSize(*folders[0])
	for _, folder := range folders {
		size := getSize(*folder)
		if size > toFree && size-toFree < smallest-toFree {
			smallest = size
		}
	}

	fmt.Printf("Smallest: %d\n", smallest)
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
