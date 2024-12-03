package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseTree(s string) *BinNode {

	if !strings.Contains(s, "(") {
		value, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Incorrect string")
			return nil
		}

		return NewBinNode(value)
	}

	rootStart := strings.Index(s, "(") + 1
	rootValue, err := strconv.Atoi(strings.TrimSpace(s[rootStart : rootStart+1]))
	if err != nil {
		fmt.Println("Incorrect string")
		return nil
	}

	stack := 0
	leftStart := rootStart + 1
	var leftSubtree, rightSubtree string

	for i := rootStart; i < len(s); i++ {
		if s[i] == '(' {
			stack++
		} else if s[i] == ')' {
			stack--
			if stack == 0 {
				leftSubtree = strings.TrimSpace(s[leftStart+1 : i+1])
				if i+2 < len(s) {
					rightSubtree = strings.TrimSpace(s[i+2 : len(s)-1])
				}
				break
			}
		}
	}

	root := &BinNode{value: rootValue}
	if leftSubtree != "" {
		root.left = parseTree(leftSubtree)
	}
	if rightSubtree != "" {
		root.right = parseTree(rightSubtree)
	}

	return root
}
