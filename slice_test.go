package leetconv

import (
    "reflect"
    "testing"
)

func TestParseSliceInt(t *testing.T){
	tests := []struct {
        name     string
        input    string
        expected []int
        wantErr  bool
    }{
        {
            name:     "empty list",
            input:    "[]",
            expected: []int{},
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
            got, err := ParseSliceInt(currTest.input)

            if (err != nil) != currTest.wantErr {
                t.Errorf("ParseSliceInt(%q) error = %v, wantErr %v", currTest.input, err, currTest.wantErr)
                return
            }

            if currTest.wantErr {
                return
            }

            if !reflect.DeepEqual(got, currTest.expected) {
                t.Errorf("ParseSliceInt(%q) = %v, want %v", currTest.input, got, currTest.expected)
            }
        })
	}
}