# leetconv

Convert LeetCode-style string representations to Go data structures.

## Installation

```bash
go get github.com/Chypakk/leetconv
```

## Usage
### Parse Linked Lists
```go
head, err := leetconv.ParseListNode("[1,2,3]")
if err != nil {
    log.Fatal(err)
}
// head → 1 → 2 → 3 → nil
```
### Parse Slices
```go
slice, err := leetconv.ParseSliceInt("[1, 2, 3]")
if err != nil {
    log.Fatal(err)
}
fmt.Println(slice) // [1 2 3]
```
