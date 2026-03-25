package leetconv

import (
	"fmt"
	"testing"
)

func TestParseTreeNode(t *testing.T ){
	tests := []struct {
        name     string
        input    string
        expected *TreeNode
        wantErr  bool
    }{
		{
			name:     "empty list",
			input:    "[]",
			expected: &TreeNode{},
			wantErr:  false,
		},
		{
			name:     "single node",
			input:    "[1]",
			expected: &TreeNode{Val: 1},
			wantErr:  false,
		},
		{
			name:     "simple tree",
			input:    "[1,2,3]",
			expected: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			wantErr:  false,
		},
		{
			name:     "tree with null left",
			input:    "[1,null,2]",
			expected: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2}},
			wantErr:  false,
		},
		{
			name:     "tree with null right",
			input:    "[1,2,null]",
			expected: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: nil},
			wantErr:  false,
		},
		{
			name:     "invalid format - no brackets",
			input:    "1,2,3",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "invalid value - not a number",
			input:    "[1,abc,3]",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "empty string",
			input:    "",
			expected: nil,
			wantErr:  true,
		},
	}

	for _, currTest := range tests{
		t.Run(currTest.name, func(t *testing.T) {
            got, err := ParseTreeNode(currTest.input)

            if (err != nil) != currTest.wantErr {
                t.Errorf("ParseTreeNode(%q) error = %v, wantErr %v", currTest.input, err, currTest.wantErr)
                return
            }

            if currTest.wantErr {
                return
            }

            if !compareTrees(got, currTest.expected) {
                t.Errorf("ParseTreeNode(%q) structures don't match", currTest.input)
				t.Logf("got: %v", treeToString(got))
				t.Logf("want: %v", treeToString(currTest.expected))
            }
        })
	}
}

func compareTrees(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return compareTrees(a.Left, b.Left) && compareTrees(a.Right, b.Right)
}

func treeToString(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	return fmt.Sprintf("%d(%s,%s)", root.Val, treeToString(root.Left), treeToString(root.Right))
}