package leetconv

import (
    "reflect"
    "testing"
)

func toSlice(head *ListNode) []int {
    var result []int
    for node := head; node != nil; node = node.Next {
        result = append(result, node.Val)
    }
    return result
}

func TestParseListNode(t *testing.T){
	tests := []struct {
        name     string
        input    string
        expected []int
        wantErr  bool
    }{
        {
            name:     "empty list",
            input:    "[]",
            expected: nil,
            wantErr:  false,
        },
        {
            name:     "single element",
            input:    "[5]",
            expected: []int{5},
            wantErr:  false,
        },
        {
            name:     "multiple elements",
            input:    "[1,2,4]",
            expected: []int{1, 2, 4},
            wantErr:  false,
        },
        {
            name:     "with spaces",
            input:    "[1, 2, 3]",
            expected: []int{1, 2, 3},
            wantErr:  false,
        },
        {
            name:     "negative numbers",
            input:    "[-1,0,1]",
            expected: []int{-1, 0, 1},
            wantErr:  false,
        },
        {
            name:     "missing brackets",
            input:    "1,2,3",
            wantErr:  true,
        },
        {
            name:     "invalid number",
            input:    "[1,x,3]",
            wantErr:  true,
        },
        {
            name:     "missing number",
            input:    "[1, ,3]",
            wantErr:  true,
        },
        {
            name:     "empty string",
            input:    "",
            wantErr:  true,
        },
    }

	for _, currTest := range tests{
		t.Run(currTest.name, func(t *testing.T) {
            head, err := ParseListNode(currTest.input)

            if (err != nil) != currTest.wantErr {
                t.Errorf("ParseListNode(%q) error = %v, wantErr %v", currTest.input, err, currTest.wantErr)
                return
            }

            if currTest.wantErr {
                return
            }

            got := toSlice(head)
            if !reflect.DeepEqual(got, currTest.expected) {
                t.Errorf("ParseListNode(%q) = %v, want %v", currTest.input, got, currTest.expected)
            }
        })
	}
}