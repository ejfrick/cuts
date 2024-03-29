# cuts
[![GoDoc](https://pkg.go.dev/badge/github.com/ejfrick/cuts)](https://pkg.go.dev/github.com/ejfrick/cuts) 
[![Build Status](https://github.com/ejfrick/cuts/actions/workflows/ci.yml/badge.svg)](https://github.com/ejfrick/cuts/actions/workflows/ci.yml)


Provides useful, generic utilities for working with slices.

## Install
`go get github.com/ejfrick/cuts`

## Usage

```go
package main

import (
	"fmt"
	"github.com/ejfrick/cuts"
	"strings"
	"time"
)

func main() {
	// dedupe some values
	words := []string{"hello", "hello", "world"}
	dedupedWords := cuts.Dedupe(words)
	fmt.Println(dedupedWords) // ["hello", "world"]

	// chunk a slice into smaller slices
	bigList := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	chunks := cuts.ChunkBy(bigList, 2) // chunk into lists of two elements
	fmt.Println(chunks)                // [["the", "quick"], ["brown", "fox"], ["jumps", "over"], ["the", "lazy"], ["dog"]]

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// get the first element that is even
	index, found := cuts.FirstWhere(nums, func(val int) bool { return val%2 == 0 })
	fmt.Println(index, found) // 2, true

	// get the last element that is odd
	index, found = cuts.LastWhere(nums, func(val int) bool { return val%2 == 1 })
	fmt.Println(index, found) // 7, true

	// get the last element that is greater than 10
	index, found = cuts.LastWhere(nums, func(val int) bool { return val > 10 })
	fmt.Println(index, found) // -1, false

	// check if any words in the slice contain the letter e
	ok := cuts.AnyWhere(bigList, func(val string) bool { return strings.Contains(val, "e") })
	fmt.Println(ok) // true

	// check if all words in the slice contain the letter e
	// fails fast!
	ok = cuts.AllWhere(bigList, func(val string) bool { return strings.Contains(val, "e") })
	fmt.Println(ok) // false
	
	// get closest value to target value
	acceptableTimes := []time.Duration{time.Hour, time.Hour * 6, time.Hour * 12}
	target := time.Hour * 7
	closest := cuts.SnapTo(acceptableTimes, target)
	fmt.Println(closest.String()) // 6h0m0s

}
```
See the [`cuts`](https://pkg.go.dev/github.com/ejfrick/cuts) package documentation for more usage information.
## License
MIT