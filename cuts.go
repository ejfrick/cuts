// Package cuts provides useful, generic utilities for working with slices.
package cuts

import (
	"cmp"
	"slices"
)

// Dedupe removes duplicate values from an array of comparable elements.
func Dedupe[T comparable](in []T) []T {
	var l []T
	for _, elem := range in {
		if !slices.Contains(l, elem) {
			l = append(l, elem)
		}
	}

	return l
}

// ChunkBy groups an array of items into batches of the given size.
func ChunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

// FirstWhere searches for the first element in a sorted slice
// that meets the condition function and returns its position.
// If no such element exists, returns -1;
// it also returns a bool saying whether an element matching the
// condition was found in the slice.
// The slice must be sorted in increasing order.
func FirstWhere[S ~[]E, E cmp.Ordered](vals S, where func(val E) bool) (int, bool) {
	ind := len(vals)
	found := false

	n := len(vals)
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if where(vals[h]) && h < ind {
			ind = h
			found = true
			i = h + 1
		} else {
			j = h
		}
	}

	if !found {
		return -1, found
	}

	return ind, found
}

// LastWhere searches for the last element in a sorted slice
// that meets the condition function and returns its position.
// If no such element exists, returns -1;
// it also returns a bool saying whether an element matching the
// condition was found in the slice.
// The slice must be sorted in increasing order.
func LastWhere[S ~[]E, E cmp.Ordered](vals S, where func(val E) bool) (int, bool) {
	ind := 0
	found := false
	n := len(vals)
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if where(vals[h]) && h > ind {
			ind = h
			found = true
			i = h + 1
		} else {
			j = h
		}
	}

	if !found {
		return -1, found
	}

	return ind, found
}

// AnyWhere returns whether any element in the slice
// satisfies the condition function.
func AnyWhere[S ~[]E, E any](vals S, where func(val E) bool) bool {
	found := false

	for _, val := range vals {
		if where(val) {
			found = true
			break
		}
	}

	return found
}

// AllWhere returns whether all elements in the slice
// satisfy the condition function.
func AllWhere[S ~[]E, E any](vals S, where func(val E) bool) bool {
	found := true

	for _, val := range vals {
		if !where(val) {
			found = false
			break
		}
	}

	return found
}

// TODO: finish batcher implementation
//type batcher struct {
//	batchSize int
//	length    int
//}
//
//func (b *batcher) options(opts ...BatchOpt) {
//	for _, opt := range opts {
//		opt(b)
//	}
//}
//
//func (b *batcher) BatchSize() int {
//	if b.batchSize == 0 {
//		if b.length < runtime.NumCPU()*2 {
//			return b.length
//		} else {
//			return b.length / (runtime.NumCPU() * 2)
//		}
//	}
//
//	return b.batchSize
//}
//
//type BatchOpt func(b *batcher)
//
//func BatchSize(size int) BatchOpt {
//	return func(b *batcher) {
//		b.batchSize = size
//	}
//}
