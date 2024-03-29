package cuts

import (
	"fmt"
	"math/rand"
	"testing"
)

func createUintSlice(max uint) []uint {
	var slice []uint
	for i := uint(0); i < max; i += uint(5) {
		slice = append(slice, i)
	}

	return slice
}

func createUintptrSlice(max uintptr) []uintptr {
	var slice []uintptr
	for i := uintptr(0); i < max; i += uintptr(5) {
		slice = append(slice, i)
	}

	return slice
}

func createFloat64Slice(max float64) []float64 {
	var slice []float64
	for i := float64(0); i < max; i += float64(5) {
		slice = append(slice, i)
	}

	return slice
}

func BenchmarkSnapTo_nums(b *testing.B) {
	for _, tc := range whereNumTT {
		b.Run(fmt.Sprintf("int/input_size_%d", tc.input), func(b *testing.B) {
			vals := createIntSlice(tc.input)
			target := rand.Intn(tc.input)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = SnapTo(vals, target)
			}
		})
		b.Run(fmt.Sprintf("uint/input_size_%d", tc.input), func(b *testing.B) {
			vals := createUintSlice(uint(tc.input))
			target := uint(rand.Intn(tc.input))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = SnapTo(vals, target)
			}
		})
		b.Run(fmt.Sprintf("uintptr/input_size_%d", tc.input), func(b *testing.B) {
			vals := createUintptrSlice(uintptr(tc.input))
			target := uintptr(rand.Intn(tc.input))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = SnapTo(vals, target)
			}
		})
		b.Run(fmt.Sprintf("float64/input_size_%d", tc.input), func(b *testing.B) {
			vals := createFloat64Slice(float64(tc.input))
			target := float64(rand.Intn(tc.input))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = SnapTo(vals, target)
			}
		})
	}
}
