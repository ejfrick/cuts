package cuts

import (
	"fmt"
	"github.com/tjarratt/babble"
	"testing"
)

func createWordList(max int) []string {
	babbler := babble.NewBabbler()
	babbler.Count = 1

	var words []string
	for i := 0; i < max; i++ {
		words = append(words, babbler.Babble())
	}

	return words
}

var snapToStrBTT = []struct {
	input int
}{
	{10},
	{100},
	{1000},
	{10000},
}

func BenchmarkSnapToStr(b *testing.B) {
	for _, tc := range snapToStrBTT {
		b.Run(fmt.Sprintf("input_size_%d", tc.input), func(b *testing.B) {
			babbler := babble.NewBabbler()
			babbler.Count = 1

			target := babbler.Babble()
			words := createWordList(tc.input)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = SnapToStr(words, target)
			}
		})
	}
}

func BenchmarkSnapToStrFunc(b *testing.B) {
	for _, tc := range snapToStrBTT {
		tc := tc
		babbler := babble.NewBabbler()
		babbler.Count = 1

		target := babbler.Babble()
		words := createWordList(tc.input)
		b.ResetTimer()
		b.Run(fmt.Sprintf("matcher=FZFV2Matcher/input_size=%d", tc.input), func(b *testing.B) {
			tt := []struct {
				name          string
				normalize     bool
				caseSensitive bool
				forward       bool
			}{
				{
					name: "no_opts",
				},
				{
					name:      "normalize",
					normalize: true,
				},
				{
					name:          "case_sensitive",
					caseSensitive: true,
				},
				{
					name:    "forward",
					forward: true,
				},
				{
					name:          "normalize_case_sensitve",
					normalize:     true,
					caseSensitive: true,
				},
				{
					name:      "normalize_forward",
					normalize: true,
					forward:   true,
				},
				{
					name:          "case_sensitive_forward",
					caseSensitive: true,
					forward:       true,
				},
				{
					name:          "all_opts",
					normalize:     true,
					caseSensitive: true,
					forward:       true,
				},
			}
			b.ResetTimer()
			for _, optSet := range tt {
				b.Run(fmt.Sprintf("opts=%s", optSet.name), func(b *testing.B) {
					var opts []FZFV2MatcherOpt[string]
					if optSet.normalize {
						opts = append(opts, NormalizeV2[string]())
					}
					if optSet.forward {
						opts = append(opts, ForwardV2[string]())
					}
					if optSet.caseSensitive {
						opts = append(opts, CaseSensitiveV2[string]())
					}
					b.ResetTimer()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							_ = SnapToStrFunc(words, target, FZFV2Matcher[string](opts...))
						}
					})
				})
			}
		})
		b.Run(fmt.Sprintf("matcher=FZFV1Matcher/input_size=%d", tc.input), func(b *testing.B) {
			tt := []struct {
				name          string
				normalize     bool
				caseSensitive bool
				forward       bool
			}{
				{
					name: "no_opts",
				},
				{
					name:      "normalize",
					normalize: true,
				},
				{
					name:          "case_sensitive",
					caseSensitive: true,
				},
				{
					name:    "forward",
					forward: true,
				},
				{
					name:          "normalize_case_sensitve",
					normalize:     true,
					caseSensitive: true,
				},
				{
					name:      "normalize_forward",
					normalize: true,
					forward:   true,
				},
				{
					name:          "case_sensitive_forward",
					caseSensitive: true,
					forward:       true,
				},
				{
					name:          "all_opts",
					normalize:     true,
					caseSensitive: true,
					forward:       true,
				},
			}
			b.ResetTimer()
			for _, optSet := range tt {
				b.Run(fmt.Sprintf("opts=%s", optSet.name), func(b *testing.B) {
					var opts []FZFV1MatcherOpt[string]
					if optSet.normalize {
						opts = append(opts, NormalizeV1[string]())
					}
					if optSet.forward {
						opts = append(opts, ForwardV1[string]())
					}
					if optSet.caseSensitive {
						opts = append(opts, CaseSensitiveV1[string]())
					}
					b.ResetTimer()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							_ = SnapToStrFunc(words, target, FZFV1Matcher[string](opts...))
						}
					})
				})
			}
		})
		b.Run(fmt.Sprintf("matcher=LevenshteinMatcher/opts=no_opts/input_size=%d", tc.input), func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = SnapToStrFunc(words, target, LevenshteinMatcher[string]())
				}
			})
		})
		b.Run(fmt.Sprintf("matcher=BagOfWordsMatcher/input_size=%d", tc.input), func(b *testing.B) {
			tt := []struct {
				name     string
				bagSizes []int
			}{
				{
					name:     "no_opts",
					bagSizes: DefaultBagSizes,
				},
				{
					name:     "slim_bag_sizes",
					bagSizes: []int{2},
				},
				{
					name:     "extended_bag_sizes",
					bagSizes: []int{2, 3, 4, 5, 6},
				},
			}
			b.ResetTimer()
			for _, optSet := range tt {
				b.Run(fmt.Sprintf("opts=%s", optSet.name), func(b *testing.B) {
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							_ = SnapToStrFunc(words, target, BagOfWordsMatcher[string](WithBagSizes[string](optSet.bagSizes)))
						}
					})
				})
			}
		})
	}
}
