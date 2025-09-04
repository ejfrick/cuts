package stringsnap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testConst string

const (
	testGoodbye      testConst = "goodbye"
	testHell         testConst = "hell"
	testHello        testConst = "hello"
	testJoyfully     testConst = "joyfully"
	testKindergarten testConst = "kindergarten"
	testHellos       testConst = "hellos"
	testHelloa       testConst = "helloa"
	testHellob       testConst = "hellob"
	testHelloc       testConst = "helloc"
	testLong         testConst = "limitations"
)

func TestSnapToStr(t *testing.T) {
	t.Parallel()
	t.Run("string", func(t *testing.T) {
		t.Parallel()
		tt := []struct {
			name     string
			input    []string
			target   string
			expected string
		}{
			{
				name:     "snapToPrev",
				input:    []string{"hello", "joyfully", "kindergarten"},
				target:   "hellos",
				expected: "hello",
			},
			{
				name:     "snapToNext",
				input:    []string{"hell", "hellos", "joyfully"},
				target:   "helloa",
				expected: "hellos",
			},
			{
				name:     "equidistant",
				input:    []string{"helloa", "helloc", "joyfully"},
				target:   "hellob",
				expected: "helloc",
			},
			{
				name:     "containsTarget",
				input:    []string{"hello", "joyfully", "kindergarten"},
				target:   "joyfully",
				expected: "joyfully",
			},
			{
				name:     "snapToLast",
				input:    []string{"hello", "joyfully", "kindergarten"},
				target:   "limitations",
				expected: "kindergarten",
			},
			{
				name:     "snapToFirst",
				input:    []string{"hello", "joyfully", "kindergarten"},
				target:   "goodbye",
				expected: "hello",
			},
		}
		for _, tc := range tt {
			tc := tc
			t.Run(string(tc.name), func(t *testing.T) {
				t.Parallel()
				actual := SnapToStr(tc.input, tc.target)
				assert.Equal(t, tc.expected, actual)
				assert.IsType(t, tc.expected, actual)
			})
		}
	})
	t.Run("const", func(t *testing.T) {
		t.Parallel()
		tt := []struct {
			name     string
			input    []testConst
			target   testConst
			expected testConst
		}{
			{
				name:     "snapToPrev",
				input:    []testConst{testHello, testJoyfully, testKindergarten},
				target:   testHellos,
				expected: testHello,
			},
			{
				name:     "snapToNext",
				input:    []testConst{testHell, testHellos, testJoyfully},
				target:   testHelloa,
				expected: testHellos,
			},
			{
				name:     "equidistant",
				input:    []testConst{testHelloa, testHelloc, testJoyfully},
				target:   testHellob,
				expected: testHelloc,
			},
			{
				name:     "containsTarget",
				input:    []testConst{testHello, testJoyfully, testKindergarten},
				target:   testJoyfully,
				expected: testJoyfully,
			},
			{
				name:     "snapToLast",
				input:    []testConst{testHello, testJoyfully, testKindergarten},
				target:   testLong,
				expected: testKindergarten,
			},
			{
				name:     "snapToFirst",
				input:    []testConst{testHello, testJoyfully, testKindergarten},
				target:   testGoodbye,
				expected: testHello,
			},
		}

		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				actual := SnapToStr(tc.input, tc.target)
				assert.Equal(t, tc.expected, actual)
				assert.IsType(t, tc.expected, actual)
			})
		}
	})
}

func TestLevenshteinMatcher(t *testing.T) {
	actual := LevenshteinMatcher[string]()
	assert.NotNil(t, actual)
	actual2 := LevenshteinMatcher[testConst]()
	assert.NotNil(t, actual2)
}

func TestLevenshteinMatcher_Closest(t *testing.T) {
	t.Parallel()
	matcher := LevenshteinMatcher[string]()
	tt := []struct {
		name     string
		next     string
		prev     string
		target   string
		expected string
	}{
		{
			name:     "return_prev",
			next:     "joyfully",
			prev:     "hello",
			target:   "hellos",
			expected: "hello",
		},
		{
			name:     "return_next",
			next:     "hellos",
			prev:     "hell",
			target:   "helloa",
			expected: "hellos",
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := matcher.Closest(tc.target, tc.next, tc.prev)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
