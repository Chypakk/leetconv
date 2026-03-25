package leetconv

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ParseTreeNode(s string) (*TreeNode, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, errors.New("input cannot be empty")
	}

	if !strings.HasPrefix(s, "[") || !strings.HasSuffix(s, "]") {
		return nil, fmt.Errorf("invalid format: expected brackets, got %q", s)
	}

	content := s[1 : len(s)-1]

	if strings.TrimSpace(content) == "" {
		return &TreeNode{}, nil
	}

	parts := strings.Split(content, ",")
	val, err := strconv.Atoi(parts[0])

	if err != nil {
		return nil, fmt.Errorf("invalid root: got %q", parts[0])
	}

	root := &TreeNode{Val: val}

	queue := []*TreeNode{root}
	i := 1

	for i < len(parts) && len(queue) > 0{
		current := queue[0]
		queue = queue[1:]

		if parts[i] != "null" {
			val, err := strconv.Atoi(parts[i])

			if err != nil {
				return nil, fmt.Errorf("invalid value: got %q", parts[i])
			}

			leftNode := &TreeNode{Val: val}
			current.Left = leftNode
			queue = append(queue, leftNode)
		}
		i++

		if parts[i] != "null" {
			val, err := strconv.Atoi(parts[i])

			if err != nil {
				return nil, fmt.Errorf("invalid value: got %q", parts[i])
			}

			rightNode := &TreeNode{Val: val}
			current.Right = rightNode
			queue = append(queue, rightNode)
		}
		i++


	}

	return root, nil
}
