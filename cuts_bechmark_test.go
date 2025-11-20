package cuts

import (
	"fmt"
	"math/rand"
	"testing"
)

func createIntSlice(maxSize int) []int {
	var slice []int
	for i := 0; i < maxSize*5; i += 5 {
		slice = append(slice, i)
	}

	return slice
}

func createRandomSlice(maxSize int) []int {
	var slice []int
	for i := 0; i < maxSize; i++ {
		slice = append(slice, rand.Intn(10))
	}
	return slice
}

var whereNumTT = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 10000},
	{input: 100000},
}

func BenchmarkDedupe(b *testing.B) {
	for _, tc := range whereNumTT {
		b.Run(fmt.Sprintf("input_size_%d", tc.input), func(b *testing.B) {
			input := createRandomSlice(tc.input)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Dedupe(input)
			}
		})
	}
}

func BenchmarkFirstWhere(b *testing.B) {
	for _, tc := range whereNumTT {
		b.Run(fmt.Sprintf("input_size_%d", tc.input), func(b *testing.B) {
			input := createIntSlice(tc.input)
			target := rand.Intn(tc.input)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = FirstWhere(input, func(val int) bool {
					return val > target
				})
			}
		})
	}
}

func BenchmarkLastWhere(b *testing.B) {
	for _, tc := range whereNumTT {
		b.Run(fmt.Sprintf("input_size_%d", tc.input), func(b *testing.B) {
			input := createIntSlice(tc.input)
			target := rand.Intn(tc.input)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = LastWhere(input, func(val int) bool {
					return val > target
				})
			}
		})
	}
}

func BenchmarkAnyWhere(b *testing.B) {
	for _, tc := range whereNumTT {
		b.Run(fmt.Sprintf("input_size_%d", tc.input), func(b *testing.B) {
			input := createIntSlice(tc.input)
			target := rand.Intn(tc.input)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = AnyWhere(input, func(val int) bool {
					return val > target
				})
			}
		})
	}
}

func BenchmarkAllWhere(b *testing.B) {
	for _, tc := range whereNumTT {
		b.Run(fmt.Sprintf("input_size_%d", tc.input), func(b *testing.B) {
			input := createIntSlice(tc.input)
			target := rand.Intn(tc.input + 1)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = AllWhere(input, func(val int) bool {
					return val > target
				})
			}
		})
	}
}

func BenchmarkChunkBy(b *testing.B) {
	sliceSizes := []int{10, 100, 1000, 10_000}
	chunkSizes := []int{0, 1, 5, 10, 100}
	for _, sliceSize := range sliceSizes {
		for _, chunkSize := range chunkSizes {
			b.Run(fmt.Sprintf("sliceSize=%d,chunkSize=%d", sliceSize, chunkSize), func(b *testing.B) {
				sliceInput := createIntSlice(sliceSize)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = ChunkBy(sliceInput, chunkSize)
				}
			})
		}
	}
}
