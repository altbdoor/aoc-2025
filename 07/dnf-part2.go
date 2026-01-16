package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func getLeaves(node *TreeNode) []*TreeNode {
	if node == nil {
		return nil
	}

	if node.Left == nil && node.Right == nil {
		return []*TreeNode{node}
	}

	leftLeaves := getLeaves(node.Left)
	rightLeaves := getLeaves(node.Right)

	leaves := append(leftLeaves, rightLeaves...)
	return leaves
}

func dnfPart2(file *os.File) {
	scanner := bufio.NewScanner(file)
	root := TreeNode{Value: -1}
	lineLen := -1
	var manualLeaves []*TreeNode

	for scanner.Scan() {
		line := scanner.Text()

		// cache line length
		if lineLen == -1 {
			lineLen = len(line)
		}

		// prepare root
		if root.Value == -1 {
			currentSourceIdx := strings.IndexRune(line, 'S')
			root.Value = currentSourceIdx
			manualLeaves = append(manualLeaves, &root)
			continue
		}

		// skip blank lines
		blankLine := strings.Repeat(".", lineLen)
		if line == blankLine {
			continue
		}

		// iter leaves, and build tree
		removeLeaves := make([]int, 0)
		newManualLeaves := make([]*TreeNode, 0)

		for leafIdx, leaf := range manualLeaves {
			strIdx := leaf.Value

			if line[strIdx] == '^' {
				// if split, build the child nodes
				leaf.Left = &TreeNode{
					Value: strIdx - 1,
				}
				leaf.Right = &TreeNode{
					Value: strIdx + 1,
				}

				// mark current node as no longer leaf
				removeLeaves = append(removeLeaves, leafIdx)

				// mark child nodes as new leaves
				newManualLeaves = append(newManualLeaves, leaf.Left)
				newManualLeaves = append(newManualLeaves, leaf.Right)
			}
		}

		// reverse-based approach to trimming leaves
		slices.Sort(removeLeaves)
		slices.Reverse(removeLeaves)

		for _, idx := range removeLeaves {
			manualLeaves = append(manualLeaves[:idx], manualLeaves[idx+1:]...)
		}

		// add back the new child nodes
		manualLeaves = append(manualLeaves, newManualLeaves...)
		fmt.Println("leaves count", len(manualLeaves))
	}

	finalLeaves := getLeaves(&root)
	fmt.Println("count:", len(finalLeaves))
}
