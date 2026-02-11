package leetconv

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ListNode represents a node in a singly-linked list.
// Matches LeetCode's definition exactly.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ParseListNode converts a string like "[1,2,3]" into a linked list.
// Returns the head node or an error if the input is malformed.
//
// Examples:
//
//	"[1,2,3]" → ListNode(1) → ListNode(2) → ListNode(3) → nil
//	"[]"      → nil
//	"[5]"     → ListNode(5) → nil
func ParseListNode(s string) (*ListNode, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, errors.New("input cannot be empty")
	}

	if !strings.HasPrefix(s, "[") || !strings.HasSuffix(s, "]") {
		return nil, fmt.Errorf("invalid format: expected brackets, got %q", s)
	}

	content := s[1 : len(s)-1]

	if strings.TrimSpace(content) == "" {
		return nil, nil
	}

	parts := strings.Split(content, ",")
	var head, curr *ListNode

	for i, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			return nil, fmt.Errorf("empty element in list at position %d", i)
		}

		val, err := strconv.Atoi(trimmed)
		if err != nil {
			return nil, fmt.Errorf("invalid integer at position %d: %w", i, err)
		}

		node := &ListNode{Val: val}
		if head == nil {
			head = node
			curr = node
		} else {
			curr.Next = node
			curr = node
		}
	}

	return head, nil
}
