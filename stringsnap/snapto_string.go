package stringsnap

import (
	"github.com/junegunn/fzf/src/algo"
	"github.com/junegunn/fzf/src/util"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/schollz/closestmatch"
	"slices"
)

// SnapToStr returns the closest element in a string array
// to the target. It uses the Levenshtein distance implementation
// from github.com/lithammer/fuzzysearch/fuzzy to calculate
// closeness.
func SnapToStr[S ~[]E, E ~string](vals S, target E) E {
	var res E
	index, found := slices.BinarySearch(vals, target)
	switch {
	case found:
		res = target
	case len(vals) < index+1:
		res = vals[len(vals)-1]
	case index == 0:
		res = vals[0]
	default:
		p1 := vals[index]
		p2 := vals[index-1]

		res1 := fuzzy.LevenshteinDistance(string(target), string(p1))
		res2 := fuzzy.LevenshteinDistance(string(target), string(p2))

		if res2 < res1 {
			res = p2
		} else {
			res = p1
		}
	}

	return res
}

// SnapToStrFunc returns the closest element in a string array
// to the target, given a provided fuzzy matcher.
func SnapToStrFunc[S ~[]E, E ~string](vals S, target E, matcher FuzzyMatcher[E]) E {
	var res E
	index, found := slices.BinarySearch(vals, target)
	switch {
	case found:
		res = target
	case len(vals) < index+1:
		res = vals[len(vals)-1]
	case index == 0:
		res = vals[0]
	default:
		res = matcher.Closest(target, vals[index], vals[index-1])
	}

	return res
}

// FuzzyMatcher is an interface for various implementations of fuzzy string matching.
type FuzzyMatcher[T ~string] interface {
	Closest(target, next, prev T) T
}

// FZFV2Matcher is a FuzzyMatcher using fzf's V2 algorithm.
// The V2 algorithm is slower than the fzf's V1 algorithm,
// but provides higher quality matches.
//
// See github.com/junegunn/fzf/src/algo for more information.
func FZFV2Matcher[T ~string](opts ...FZFV2MatcherOpt[T]) FuzzyMatcher[T] {
	m := &fZFV2Matcher[T]{}

	m.options(opts...)

	return m
}

// FZFV2MatcherOpt configures a FZFV2Matcher.
type FZFV2MatcherOpt[T ~string] interface {
	apply(f *fZFV2Matcher[T])
}

// NormalizeV2 configures a FZFV2Matcher to normalize strings.
func NormalizeV2[T ~string]() FZFV2MatcherOpt[T] {
	return v2Normalize[T](true)
}

type v2Normalize[T ~string] bool

func (n v2Normalize[T]) apply(f *fZFV2Matcher[T]) {
	f.normalize = bool(n)
}

// CaseSensitiveV2 configures a FZFV2Matcher to be case-sensitive.
func CaseSensitiveV2[T ~string]() FZFV2MatcherOpt[T] {
	return v2CaseSensitive[T](true)
}

type v2CaseSensitive[T ~string] bool

func (c v2CaseSensitive[T]) apply(f *fZFV2Matcher[T]) {
	f.caseSensitive = bool(c)
}

// ForwardV2 configures a FZFV2Matcher to look forward.
func ForwardV2[T ~string]() FZFV2MatcherOpt[T] {
	return v2Forward[T](true)
}

type v2Forward[T ~string] bool

func (v v2Forward[T]) apply(f *fZFV2Matcher[T]) {
	f.forward = bool(v)
}

// WithSlabV2 configures a FZFV2Matcher to use a custom Slab.
func WithSlabV2[T ~string](slab *util.Slab) FZFV2MatcherOpt[T] {
	return v2Slab[T](*slab)
}

type v2Slab[T ~string] util.Slab

func (s v2Slab[T]) apply(f *fZFV2Matcher[T]) {
	f.slab = (*util.Slab)(&s)
}

type fZFV2Matcher[T ~string] struct {
	caseSensitive bool
	normalize     bool
	forward       bool
	slab          *util.Slab
}

func (f *fZFV2Matcher[T]) options(opts ...FZFV2MatcherOpt[T]) {
	for _, opt := range opts {
		opt.apply(f)
	}
}

func (f *fZFV2Matcher[T]) Closest(target, next, prev T) T {
	input := util.RunesToChars([]rune(target))
	res1, _ := algo.FuzzyMatchV2(f.caseSensitive, f.normalize, f.forward, &input, []rune(next), false, f.slab)
	res2, _ := algo.FuzzyMatchV2(f.caseSensitive, f.normalize, f.forward, &input, []rune(prev), false, f.slab)

	if res2.Score > res1.Score {
		return prev
	}

	return next
}

// FZFV1Matcher is a FuzzyMatcher using fzf's V1 algorithm.
// The V2 algorithm is faster than the fzf's V2 algorithm,
// but provides lower quality matches.
//
// See github.com/junegunn/fzf/src/algo for more information.
func FZFV1Matcher[T ~string](opts ...FZFV1MatcherOpt[T]) FuzzyMatcher[T] {
	m := &fZFV1Matcher[T]{}

	m.options(opts...)

	return m
}

type FZFV1MatcherOpt[T ~string] interface {
	apply(f *fZFV1Matcher[T])
}

// NormalizeV1 configures a FZFV1Matcher to normalize strings.
func NormalizeV1[T ~string]() FZFV1MatcherOpt[T] {
	return v1Normalize[T](true)
}

type v1Normalize[T ~string] bool

func (n v1Normalize[T]) apply(f *fZFV1Matcher[T]) {
	f.normalize = bool(n)
}

// CaseSensitiveV1 configures a FZFV1Matcher to be case-sensitive.
func CaseSensitiveV1[T ~string]() FZFV1MatcherOpt[T] {
	return v1CaseSensitive[T](true)
}

type v1CaseSensitive[T ~string] bool

func (c v1CaseSensitive[T]) apply(f *fZFV1Matcher[T]) {
	f.caseSensitive = bool(c)
}

// ForwardV1 configures a FZFV1Matcher to look forward.
func ForwardV1[T ~string]() FZFV1MatcherOpt[T] {
	return v1Forward[T](true)
}

type v1Forward[T ~string] bool

func (v v1Forward[T]) apply(f *fZFV1Matcher[T]) {
	f.forward = bool(v)
}

// WithSlabV1 configures a FZFV1Matcher to use a custom Slab.
func WithSlabV1[T ~string](slab *util.Slab) FZFV1MatcherOpt[T] {
	return v1Slab[T](*slab)
}

type v1Slab[T ~string] util.Slab

func (s v1Slab[T]) apply(f *fZFV1Matcher[T]) {
	f.slab = (*util.Slab)(&s)
}

type fZFV1Matcher[T ~string] struct {
	caseSensitive bool
	normalize     bool
	forward       bool
	slab          *util.Slab
}

func (f *fZFV1Matcher[T]) options(opts ...FZFV1MatcherOpt[T]) {
	for _, opt := range opts {
		opt.apply(f)
	}
}

func (f *fZFV1Matcher[T]) Closest(target, next, prev T) T {
	input := util.RunesToChars([]rune(target))

	res1, _ := algo.FuzzyMatchV1(f.caseSensitive, f.normalize, f.forward, &input, []rune(next), false, f.slab)
	res2, _ := algo.FuzzyMatchV1(f.caseSensitive, f.normalize, f.forward, &input, []rune(prev), false, f.slab)

	if res2.Score > res1.Score {
		return prev
	}

	return next
}

// LevenshteinMatcher is a FuzzyMatcher that uses
// Levenshtein distance to calculate closeness,
// as implemented by github.com/lithammer/fuzzysearch/fuzzy.
func LevenshteinMatcher[T ~string]() FuzzyMatcher[T] {
	return &levenshteinMatcher[T]{}
}

type levenshteinMatcher[T ~string] struct{}

func (r *levenshteinMatcher[T]) Closest(target, next, prev T) T {
	res1 := fuzzy.LevenshteinDistance(string(target), string(next))
	res2 := fuzzy.LevenshteinDistance(string(target), string(prev))

	if res2 < res1 {
		return prev
	}

	return next
}

// BagOfWordsMatcher is a FuzzyMatcher that uses a bag-of-words approach
// to calculate closeness, as implemented by
// github.com/schollz/closestmatch.
func BagOfWordsMatcher[T ~string](opts ...BagOfWordsMatcherOpt[T]) FuzzyMatcher[T] {
	m := &bagOfWordsMatcher[T]{bagSizes: DefaultBagSizes}

	m.options(opts...)

	return m
}

// BagOfWordsMatcherOpt configures a BagOfWordsMatcher
type BagOfWordsMatcherOpt[T ~string] interface {
	apply(b *bagOfWordsMatcher[T])
}

// WithBagSizes configures the bag sizes of BagOfWordsMatcher
func WithBagSizes[T ~string](b []int) BagOfWordsMatcherOpt[T] {
	return bagSizes[T](b)
}

// DefaultBagSizes are the default bag sizes uses by
// BagOfWordsMatcher, if none are specified.
var DefaultBagSizes = []int{2, 3, 4}

type bagSizes[T ~string] []int

func (w bagSizes[T]) apply(b *bagOfWordsMatcher[T]) {
	b.bagSizes = w
}

type bagOfWordsMatcher[T ~string] struct {
	bagSizes []int
}

func (b *bagOfWordsMatcher[T]) options(opts ...BagOfWordsMatcherOpt[T]) {
	for _, opt := range opts {
		opt.apply(b)
	}
}

func (b *bagOfWordsMatcher[T]) Closest(target, next, prev T) T {
	var res T
	words := []string{string(next), string(prev)}
	cm := closestmatch.New(words, b.bagSizes)

	var closest any
	closest = cm.Closest(string(target))

	res = closest.(T)

	return res
}
