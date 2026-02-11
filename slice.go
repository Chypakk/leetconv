package leetconv

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ParseSliceInt converts a string like "[1,2,3]" into a slice.
// Returns the []int or an error if the input is malformed.
//
// Examples:
//
//	"[1,2,3]" → [1, 2, 3]
//	"[]"      → []
//	"[5]"     → [5]
func ParseSliceInt(s string) ([]int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, errors.New("input cannot be empty")
	}

	if !strings.HasPrefix(s, "[") || !strings.HasSuffix(s, "]") {
		return nil, fmt.Errorf("invalid format: expected brackets, got %q", s)
	}

	content := s[1 : len(s)-1]

	if strings.TrimSpace(content) == "" {
		return []int{}, nil
	}

	parts := strings.Split(content, ",")
	sl := make([]int, 0, len(parts))

	for i, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			return []int{}, fmt.Errorf("empty element in list at position %d", i)
		}

		val, err := strconv.Atoi(trimmed)
		if err != nil {
			return []int{}, fmt.Errorf("invalid integer at position %d: %w", i, err)
		}

		sl = append(sl, val)
	}

	return sl, nil
}
